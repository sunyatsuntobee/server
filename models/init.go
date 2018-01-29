package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

const (
	dbDSN string = "/tobee"
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
