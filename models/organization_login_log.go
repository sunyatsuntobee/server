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

// OrganizationLoginLogDataAccessObject provides database access for Model
// OrganizationLoginLog
type OrganizationLoginLogDataAccessObject struct{}

// OrganizationLoginLogDAO instance of OrganizationLoginLogDataAccessObject
var OrganizationLoginLogDAO *OrganizationLoginLogDataAccessObject

// NewOrganizationLoginLog creates a new NewOrganizationLoginLog
func NewOrganizationLoginLog(loginTime time.Time, loginLocation string,
	loginDevice string, organizationID int) *OrganizationLoginLog {
	return &OrganizationLoginLog{
		LoginTime:      loginTime,
		LoginLocation:  loginLocation,
		LoginDevice:    loginDevice,
		OrganizationID: organizationID,
	}
}

// TableName returns table name
func (*OrganizationLoginLogDataAccessObject) TableName() string {
	return "organization_login_logs"
}
