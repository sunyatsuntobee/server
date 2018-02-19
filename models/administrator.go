package models

// Administrator Model
type Administrator struct {
	ID       int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	Name     string `xorm:"name VARCHAR(45) NOTNULL"`
	Password string `xorm:"password VARCHAR(45) NOTNULL"`
	Level    int    `xorm:"level INT NOTNULL"`
}

func NewAdministrator(name string, password string, level int) {
	return Administrator{Name:name, Password:password, Level:level}
}