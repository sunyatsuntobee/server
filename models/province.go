package models

// Province Model
type Province struct {
	ID   int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	Name string `xorm:"name VARCHAR(20) NOTNULL" json:"name"`
}
