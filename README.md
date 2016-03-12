# hopen-orm
orm开发框架

package orm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/who246/hopen-orm/db"
	"testing"
)

type MyModel struct {
	Id         int    `field:"id"`
	CreateTime string `field:"create_time"`
	ModifyTime string `field:"modify_time"`
	Name       string `field:"name"`
	Nums       int    `field:"nums"`
}

func Test_List(t *testing.T) {
	db.RegisterDef("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", "root", "123", "127.0.0.1", 3306, "test"))
	m := &MyModel{}
	o := New(m)
	sql := "SELECT * FROM test.goblog_type where id = ?"
	ms, err := o.List(sql, 2)
	if err != nil {
		fmt.Printf("err", err)
		return
	}
	t.Log(len(ms))
	for _, m := range ms {
		u := m.(*MyModel)
		fmt.Println(u.CreateTime)
		fmt.Println(u.Id)
	}
}

func Test_one(t *testing.T) {
	db.RegisterDef("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", "root", "123", "127.0.0.1", 3306, "test"))
	m := &MyModel{}
	o := New(m)
	sql := "SELECT * FROM test.goblog_type where id = ?"
	obj, err := o.One(sql, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(obj)
}

func Test_save(t *testing.T) {
	db.RegisterDef("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", "root", "123", "127.0.0.1", 3306, "test"))
	m := &MyModel{}
	o := New(m)
	sql := "INSERT INTO test.goblog_type (id,create_time,modify_time,name,nums) VALUES (?,?,?,?,?)"
	obj := &MyModel{Id: 5, CreateTime: "2015-10-17 17:12:01", ModifyTime: "2015-10-17 17:12:01", Name: "testttt", Nums: 2}
	o.Save(sql, obj)
}
