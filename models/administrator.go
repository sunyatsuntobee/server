package models

// Administrator Model
type Administrator struct {
	ID       int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	Name     string `xorm:"name VARCHAR(20) NOTNULL"`
	Password string `xorm:"password VARCHAR(50) NOTNULL"`
	Level    int    `xorm:"level INT NOTNULL"`
}

// NewAdministrator creates a new administrator
func NewAdministrator(name string, password string, level int) *Administrator {
	return &Administrator{
		Name:     name,
		Password: password,
		Level:    level,
	}
}
