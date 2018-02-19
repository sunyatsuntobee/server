package models

import "time"

// OrganizationLoginLog Model
type OrganizationLoginLog struct {
	ID             int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	LoginTime      time.Time `xorm:"login_time DATETIME NOTNULL"`
	LoginLocation  string    `xorm:"login_location VARCHAR(20) NOTNULL"`
	LoginDevice    string    `xorm:"login_device VARCHAR(20) NOTNULL"`
	OrganizationID int       `xorm:"organization_id NOTNULL INDEX(organization_id_idx)"`
}

type OrganizationLoginLogDataAccessObject struct{}

var OrganizationLoginLogDAO *OrganizationLoginLogDataAccessObject

func (*OrganizationLoginLogDataAccessObject) TableName() string {
	return "organization_login_logs"
}
