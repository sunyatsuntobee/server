package models

import "time"

type Administrator_login_logs struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	LoginTime time.Time `xorm:"login_time DATETIME NOTNULL"`
	LoginLocation string `xorm:"login_location VARCHAR(45) NOTNULL"`
	LoginDevice string `xorm:"login_device VARCHAR(45) NOTNULL"`
	AdministratorID int `xorm:"administrator_id INT NOTNULL INDEX(administrator_id_idx)"`
}