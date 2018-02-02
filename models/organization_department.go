package models

// OrganizationDepartment Model
type OrganizationDepartment struct {
	ID             int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	Name           string `xorm:"name VARCHAR(45) NOTNULL"`
	OrganizationID int    `xorm:"organization_id INT NOTNULL INDEX(organization_id_idx)"`
}

type OrganizationDepartmentDataAccessObject struct{}

var OrganizationDepartmentDAO *OrganizationDepartmentDataAccessObject

func (*OrganizationDepartmentDataAccessObject) TableName() string {
	return "organization_departments"
}
