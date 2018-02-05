package models

import "github.com/sunyatsuntobee/server/logger"

// OrganizationDepartment Model
type OrganizationDepartment struct {
	ID             int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	Name           string `xorm:"name VARCHAR(45) NOTNULL" json:"name"`
	OrganizationID int    `xorm:"organization_id INT NOTNULL INDEX(organization_id_idx)" json:"organization_id"`
}

type OrganizationDepartmentDataAccessObject struct{}

var OrganizationDepartmentDAO *OrganizationDepartmentDataAccessObject

func (*OrganizationDepartmentDataAccessObject) TableName() string {
	return "organization_departments"
}

func (*OrganizationDepartmentDataAccessObject) DeleteByOID(oid int) {
	var buf OrganizationDepartment
	_, err := orm.Table(OrganizationDepartmentDAO.TableName()).
		Where("organization_id=?", oid).
		Unscoped().Delete(&buf)
	logger.LogIfError(err)
}

func (*OrganizationDepartmentDataAccessObject) InsertOne(
	department *OrganizationDepartment) {
	_, err := orm.Table(OrganizationDepartmentDAO.TableName()).
		InsertOne(department)
	logger.LogIfError(err)
}
