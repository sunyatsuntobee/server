package models

import "github.com/sunyatsuntobee/server/logger"

// OrganizationsContactors Model
type OrganizationsContactors struct {
	ID             int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	OrganizationID int `xorm:"organization_id INT NOTNULL INDEX(organization_id_idx)" json:"organization_id"`
	ContactorID    int `xorm:"contactor_id INT NOTNULL INDEX(contact_id_idx)" json:"contact_id"`
}

type OrganizationsContactorsDataAccessObject struct{}

var OrganizationsContactorsDAO *OrganizationsContactorsDataAccessObject

func (*OrganizationsContactorsDataAccessObject) TableName() string {
	return "organizations_contactors"
}

func NewOrganizationsContactors(organizationId int,
	contactId int) *OrganizationsContactors {
	return &OrganizationsContactors{
		OrganizationID: organizationId,
		ContactorID:    contactId,
	}
}

func (*OrganizationsContactorsDataAccessObject) DeleteByOID(oid int) {
	var buf OrganizationsContactors
	_, err := orm.Table(OrganizationsContactorsDAO.TableName()).
		Where("organization_id=?", oid).Unscoped().Delete(&buf)
	logger.LogIfError(err)
}

func (*OrganizationsContactorsDataAccessObject) InsertOne(
	r *OrganizationsContactors) {
	_, err := orm.Table(OrganizationsContactorsDAO.TableName()).
		InsertOne(r)
	logger.LogIfError(err)
}
