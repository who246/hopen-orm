package orm

import (
	"database/sql"
	"github.com/who246/hopen-orm/db"
	"reflect"
)

func NewAndRegister(obj interface{}, registerName string) *Orm {
	return &Orm{registerName: registerName, t: obj}
}

func New(obj interface{}) *Orm {
	return NewAndRegister(obj, db.DefRegisterName)
}

type Orm struct {
	registerName string
	t            interface{}
}

func (m *Orm) GetConn() *sql.DB {
	if m.registerName == "" {
		return db.GetDefConn()
	}
	return db.GetConn(m.registerName)
}
func (m *Orm) Save(sql string, obj interface{}) (int64, error) {
	valus := reflect.ValueOf(obj).Elem()
	param := []interface{}{}
	for i := 0; i < valus.NumField(); i++ {
		typ := valus.Type().Field(i)
		val := valus.Field(i)
		if len(typ.Tag.Get("field")) > 0 {
			param = append(param, val.Interface())
		}
	}

	res, err := m.GetConn().Exec(sql, param...)
	if err != nil {
		fmt.Println(err)
	}
	return res.LastInsertId()
}

func (m *Orm) Update(sql string, obj interface{}) (int64, error) {
	valus := reflect.ValueOf(obj).Elem()
	param := []interface{}{}
	for i := 0; i < valus.NumField(); i++ {
		typ := valus.Type().Field(i)
		val := valus.Field(i)
		if len(typ.Tag.Get("field")) > 0 {
			param = append(param, val.Interface())
		}
	}
	res, err := m.GetConn().Exec(sql, param...)
	if err != nil {
		fmt.Println(err)
	}
	return res.RowsAffected()
}

func (m *Orm) List(query string, args ...interface{}) ([]interface{}, error) {
	rows, err := m.GetConn().Query(query, args...)
	if err != nil {
		return nil, err
	}
	return bulidList(rows, m.t)
}

func (m *Orm) One(query string, args ...interface{}) (interface{}, error) {
	row := m.GetConn().QueryRow(query, args...)
	return bulidOne(row, m.t)
}
