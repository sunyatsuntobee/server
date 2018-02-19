package models

import "github.com/sunyatsuntobee/server/logger"

// OrganizationDepartment Model
type OrganizationDepartment struct {
	ID             int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	Name           string `xorm:"name VARCHAR(45) NOTNULL" json:"name"`
	OrganizationID int    `xorm:"organization_id INT NOTNULL INDEX(organization_id_idx)" json:"organization_id"`
}

// OrganizationDepartmentDataAccessObject provides database access for Model
// OrganizationDepartment
type OrganizationDepartmentDataAccessObject struct{}

// OrganizationDepartmentDAO instance of OrganizationDepartmentDataAccessObject
var OrganizationDepartmentDAO *OrganizationDepartmentDataAccessObject

// NewOrganizationDepartment creates a new organization department
func NewOrganizationDepartment(name string,
	organizationID int) *OrganizationDepartment {
	return &OrganizationDepartment{
		Name:           name,
		OrganizationID: organizationID,
	}
}

// TableName returns table name
func (*OrganizationDepartmentDataAccessObject) TableName() string {
	return "organization_departments"
}

// DeleteByOID deletes all organization departments according to an
// organization ID
func (*OrganizationDepartmentDataAccessObject) DeleteByOID(oid int) {
	var buf OrganizationDepartment
	_, err := orm.Table(OrganizationDepartmentDAO.TableName()).
		Where("organization_id=?", oid).
		Unscoped().Delete(&buf)
	logger.LogIfError(err)
}

// InsertOne inserts a new Organization Department
func (*OrganizationDepartmentDataAccessObject) InsertOne(
	department *OrganizationDepartment) {
	_, err := orm.Table(OrganizationDepartmentDAO.TableName()).
		InsertOne(department)
	logger.LogIfError(err)
}
