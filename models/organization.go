package models

import "github.com/sunyatsuntobee/server/logger"

// Organization Model
type Organization struct {
	ID          int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	Name        string `xorm:"name VARCHAR(45) NOTNULL" json:"name"`
	Phone       string `xorm:"phone VARCHAR(45) NOTNULL UNIQUE" json:"phone"`
	Password    string `xorm:"password VARCHAR(45) NOTNULL" json:"password"`
	Collage     string `xorm:"collage VARCHAR(45) NOTNULL" json:"collage"`
	LogoURL     string `xorm:"logo_url VARCHAR(45)" json:"logo_url"`
	Description string `xorm:"description VARCHAR(45)" json:"description"`
}

type OrganizationDataAccessObject struct{}

var OrganizationDAO *OrganizationDataAccessObject

func (*OrganizationDataAccessObject) TableName() string {
	return "organizations"
}

func (*OrganizationDataAccessObject) FindAll() []Organization {
	organizations := make([]Organization, 0)
	err := orm.Table(OrganizationDAO.TableName()).Find(&organizations)
	logger.LogIfError(err)
	return organizations
}

func (*OrganizationDataAccessObject) FindFullByID(id int) []OrganizationFull {
	organizations := make([]OrganizationFull, 0)
	err := orm.Table(OrganizationDAO.TableName()).
		Join("LEFT OUTER", "organization_contact_relationships",
			"organizations.id=organization_contact_relationships.organization_id").
		Join("LEFT OUTER", UserDAO.TableName(),
			"organization_contact_relationships.contact_id=users.id").
		Join("LEFT OUTER", OrganizationDepartmentDAO.TableName(),
			"organizations.id=organization_departments.organization_id").
		Join("LEFT OUTER", OrganizationLoginLogDAO.TableName(),
			"organizations.id=organization_login_logs.organization_id").
		Join("LEFT OUTER", ActivityDAO.TableName(),
			"organizations.id=activities.organization_id").
		Join("LEFT OUTER", ActivityStageDAO.TableName(),
			"activities.id=activity_stages.activity_id").
		Where("organizations.id=?", id).Find(&organizations)
	logger.LogIfError(err)
	return organizations

}

func (*OrganizationDataAccessObject) FindByID(id int) (Organization, bool) {
	var o Organization
	has, err := orm.Table(OrganizationDAO.TableName()).ID(id).Get(&o)
	logger.LogIfError(err)
	return o, has
}

func (*OrganizationDataAccessObject) UpdateOne(o *Organization) {
	_, err := orm.Table(OrganizationDAO.TableName()).ID(o.ID).Update(o)
	logger.LogIfError(err)
}
