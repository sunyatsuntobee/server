package models

import "github.com/sunyatsuntobee/server/logger"

// OrganizationContactRelationship Model
type OrganizationContactRelatonship struct {
	ID             int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	OrganizationID int `xorm:"organization_id INT NOTNULL INDEX(organization_id_idx)" json:"organization_id"`
	ContactID      int `xorm:"contact_id INT NOTNULL INDEX(contact_id_idx)" json:"contact_id"`
}

type OrganizationContactRelationshipDataAccessObject struct{}

var OrganizationContactRelationshipDAO *OrganizationContactRelationshipDataAccessObject

func NewOrganizationContactRealationship(organization_id int, contact_id int) {
	return &OrganizationContactRelatonship{OrganizationID:organization_id, ContactID:contact_id}
}

func (*OrganizationContactRelationshipDataAccessObject) TableName() string {
	return "organization_contact_relationships"
}

func (*OrganizationContactRelationshipDataAccessObject) DeleteByOID(oid int) {
	var buf OrganizationContactRelatonship
	_, err := orm.Table(OrganizationContactRelationshipDAO.TableName()).
		Where("organization_id=?", oid).Unscoped().Delete(&buf)
	logger.LogIfError(err)
}

func (*OrganizationContactRelationshipDataAccessObject) InsertOne(
	r *OrganizationContactRelatonship) {
	_, err := orm.Table(OrganizationContactRelationshipDAO.TableName()).
		InsertOne(r)
	logger.LogIfError(err)
}
