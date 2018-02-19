package models

import "time"

// OrganizationLoginLog Model
type OrganizationLoginLog struct {
	ID             int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	LoginTime      time.Time `xorm:"login_time NOTNULL"`
	LoginLocation  string    `xorm:"login_location NOTNULL"`
	LoginDevice    string    `xorm:"login_device NOTNULL"`
	OrganizationID int       `xorm:"organization_id NOTNULL INDEX(organization_id_idx)"`
}

type OrganizationLoginLogDataAccessObject struct{}

var OrganizationLoginLogDAO *OrganizationLoginLogDataAccessObject

func NewOrganizationLoginLog(login_time time.Time, login_location string,
							 login_device string, organization_id int) {
	return &OrganizationLoginLog{LoginTime:login_time, LoginLocation:login_location,
								 LoginDevice:login_device, OrganizationID:organization_id}
}

func (*OrganizationLoginLogDataAccessObject) TableName() string {
	return "organization_login_logs"
}
