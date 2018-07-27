package models

import (
	_ "github.com/go-sql-driver/mysql" // Mysql Driver
	"github.com/go-xorm/xorm"
)

const (
	// dbDSN string = "tobeeNew:Passw0r_@tcp(120.79.53.185)/tobee"
	dbDSN           string = "tobee:Passw0r_@/tobee"
	mysqlTimeFormat string = "2006-01-02 15:04:05"
)

var (
	orm *xorm.Engine
)

func init() {
	var err error
	orm, err = xorm.NewEngine("mysql", dbDSN)
	if err != nil {
		panic(err)
	}
}
