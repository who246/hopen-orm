package orm

import (
	"database/sql"
	"reflect"
)

func bulidList(rows *sql.Rows, o interface{}) ([]interface{}, error) {
	var models []interface{}
	val := []interface{}{}
	for rows.Next() {
		m := reflect.New(reflect.TypeOf(o).Elem()).Elem()
		val = val[len(val):]
		for i := 0; i < m.NumField(); i++ {
			if name := m.Type().Field(i).Tag.Get("field"); len(name) > 0 {
				val = append(val, m.Field(i).Addr().Interface())
			}
		}
		rows.Scan(val...)
		models = append(models, m.Addr().Interface())
	}

	return models, nil
}

func bulidOne(row *sql.Row, o interface{}) (interface{}, error) {
	val := []interface{}{}
	m := reflect.New(reflect.TypeOf(o).Elem()).Elem()
	val = val[len(val):]
	for i := 0; i < m.NumField(); i++ {
		if name := m.Type().Field(i).Tag.Get("field"); len(name) > 0 {
			val = append(val, m.Field(i).Addr().Interface())
		}
	}
	err := row.Scan(val...)
	if err != nil {
		return nil, err
	}
	return m.Addr().Interface(), nil
}
