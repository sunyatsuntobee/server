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

// NewAdministratorLoginLog creates a new AdministratorLoginLog
func NewAdministratorLoginLog(loginTime time.Time, loginLocation string,
	loginDevice string, administratorID int) *AdministratorLoginLog {
	return &AdministratorLoginLog{
		LoginTime:       loginTime,
		LoginLocation:   loginLocation,
		LoginDevice:     loginDevice,
		AdministratorID: administratorID,
	}
}
