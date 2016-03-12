package db

import (
	"database/sql"
	"sync"
)

type config struct {
	registerName string
	driverName   string
	db           *sql.DB
}

var (
	connects        map[string]*config
	registersNameMu sync.Mutex
	isRegister      bool
)

const (
	DefRegisterName = "default"
)

func init() {
	isRegister = false
	connects = make(map[string]*config)
}
func RegisterDef(driverName, url string) {
	Register(DefRegisterName, driverName, url)
}
func Register(registerName, driverName, url string) {
	db, err := sql.Open(driverName, url)
	if err != nil {
		panic(err)
	}
	c := &config{registerName: registerName, driverName: driverName, db: db}
	connects[registerName] = c
	isRegister = true
}

func GetConn(registerName string) *sql.DB {
	if !isRegister {
		panic("[error]please register driver")
	}
	registersNameMu.Lock()
	config := connects[registerName]
	registersNameMu.Unlock()
	return config.db
}
func GetDefConn() *sql.DB {
	return GetConn(DefRegisterName)
}
