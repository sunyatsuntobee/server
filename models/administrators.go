package models

type Administrators struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	Name string `xorm:"name VARCHAR(45) NOTNULL"`
	Password string `xorm:"password VARCHAR(45) NOTNULL"`
	Level int `xorm:"level INT NOTNULL"`
}