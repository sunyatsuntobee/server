package models

import (
	_ "github.com/go-sql-driver/mysql" // Mysql Driver
	"github.com/go-xorm/xorm"
)

const (
	// dbDSN string = "tobee:123321@tcp(120.79.53.185:3306)/tobee"
	// dbDSN  string = "tobee:Passw0r_@/tobee"
	dbDSN string = "tobee:@/tobee"

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
