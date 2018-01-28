package models

import "time"

// User Model
type User struct {
	ID          int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	Username    string    `xorm:"username VARCHAR(45) NOTNULL"`
	Phone       string    `xorm:"phone PK VARCHAR(45) NOTNULL"`
	Password    string    `xorm:"password VARCHAR(45) NOTNULL"`
	Location    string    `xorm:"location VARCHAR(45) NOTNULL"`
	CreateTime  time.Time `xorm:"create_time TIMESTAMP NOTNULL CREATED"`
	VIP         bool      `xorm:"vip INT NOTNULL"`
	Camera      string    `xorm:"camera VARCHAR(45)"`
	Description string    `xorm:"description VARCHAR(45)"`
	Occupation  string    `xorm:"occupation VARCHAR(45)"`
	Collage     string    `xorm:"collage VARCHAR(45)"`
}
