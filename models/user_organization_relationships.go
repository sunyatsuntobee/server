package models

type User_organization_relationships struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	UserID int `xorm:"user_id NOTNULL INDEX(user_id_idx)"`
	OrganizationID int `xorm:"organization_id INT NOTNULL INDEX(organization_id_idx)"`
}