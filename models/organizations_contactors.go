package models

import "github.com/sunyatsuntobee/server/logger"

// OrganizationsContactors Model
type OrganizationsContactors struct {
	ID             int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	OrganizationID int `xorm:"organization_id INT NOTNULL INDEX(organization_id_idx)" json:"organization_id"`
	ContactorID    int `xorm:"contactor_id INT NOTNULL INDEX(contact_id_idx)" json:"contact_id"`
}

// OrganizationsContactorsDataAccessObject provides database access for Model
// OrganizationsContactors
type OrganizationsContactorsDataAccessObject struct{}

// OrganizationsContactorsDAO instance of
// OrganizationsContactorsDataAccessObject
var OrganizationsContactorsDAO *OrganizationsContactorsDataAccessObject

// TableName returns table name
func (*OrganizationsContactorsDataAccessObject) TableName() string {
	return "organizations_contactors"
}

// NewOrganizationsContactors creates a new relationship
func NewOrganizationsContactors(organizationID int,
	contactID int) *OrganizationsContactors {
	return &OrganizationsContactors{
		OrganizationID: organizationID,
		ContactorID:    contactID,
	}
}

// DeleteByOID deletes all contactors according to an organization ID
func (*OrganizationsContactorsDataAccessObject) DeleteByOID(oid int) {
	var buf OrganizationsContactors
	_, err := orm.Table(OrganizationsContactorsDAO.TableName()).
		Where("organization_id=?", oid).Unscoped().Delete(&buf)
	logger.LogIfError(err)
}

// InsertOne inserts a new relationship
func (*OrganizationsContactorsDataAccessObject) InsertOne(
	r *OrganizationsContactors) {
	_, err := orm.Table(OrganizationsContactorsDAO.TableName()).
		InsertOne(r)
	logger.LogIfError(err)
}
