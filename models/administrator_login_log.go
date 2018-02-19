package models

import "time"

// AdministratorLoginLog Model
type AdministratorLoginLog struct {
	ID              int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	LoginTime       time.Time `xorm:"login_time DATETIME NOTNULL"`
	LoginLocation   string    `xorm:"login_location VARCHAR(20) NOTNULL"`
	LoginDevice     string    `xorm:"login_device VARCHAR(20) NOTNULL"`
	AdministratorID int       `xorm:"administrator_id INT NOTNULL INDEX(administrator_id_idx)"`
}

func NewAdministratorLoginLog(login_time time.Time, login_location string,
	login_device string, administrator_id int) *AdministratorLoginLog {
	return &AdministratorLoginLog{
		LoginTime:       login_time,
		LoginLocation:   login_location,
		LoginDevice:     login_device,
		AdministratorID: administrator_id,
	}
}
