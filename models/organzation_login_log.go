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
