package models

// Organization Model
type Organization struct {
	ID          int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	Name        string `xorm:"name VARCHAR(45) NOTNULL"`
	Phone       string `xorm:"phone VARCHAR(45) NOTNULL UNIQUE"`
	Password    string `xorm:"password VARCHAR(45) NOTNULL"`
	Collage     string `xorm:"collage VARCHAR(45) NOTNULL"`
	LogoURL     string `xorm:"logo_url VARCHAR(45)"`
	Description string `xorm:"description VARCHAR(45)"`
}
