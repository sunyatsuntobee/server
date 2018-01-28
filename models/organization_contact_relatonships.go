package models

type Organization_contact_relatonships struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	OrganizationID int `xorm:"organization_id INT NOTNULL INDEX(organization_id_idx)"`
	ContactID int `xorm:"contact_id INT NOTNULL INDEX(contact_id_idx)"`
}