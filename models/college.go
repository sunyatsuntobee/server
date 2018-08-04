package models

import "github.com/sunyatsuntobee/server/logger"

//college model
type College struct {
	ID              int     `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	Name            string  `xorm:"name VARCHAR(50) NOTNULL" json:"name"`
	LogoURL         string  `xorm:"logo_url VARCHAR(50)" json:"logo_url"`
	ImageURL        string  `xorm:"image_url VARCHAR(50)" json:"image_url"`
}

//CollegeDataAccessObject provides database access for Model College
type CollegeDataAccessObject struct{}

// CollegeDAO instance of CollegeDataAccessObject
var CollegeDAO *CollegeDataAccessObject

// TableName returns table name
func (*CollegeDataAccessObject) TableName() string {
	return "colleges"
}

// NewCollege creates a new college
func NewCollege(name string) *College {
	return &College{
		Name:  name,
	}
}

// FindAll finds all colleges
func (*CollegeDataAccessObject) FindAll() []College {
	colleges := make([]College, 0)
	err := orm.Table(CollegeDAO.TableName()).Find(&colleges)
	logger.LogIfError(err)
	return colleges
}


// InsertOne inserts a College
func (*CollegeDataAccessObject) InsertOne(college *College) {
	_, err := orm.Table(CollegeDAO.TableName()).InsertOne(college)
	logger.LogIfError(err)
}

// UpdateOne updates a College
func (*CollegeDataAccessObject) UpdateOne(college *College) {
	_, err := orm.Table(CollegeDAO.TableName()).ID(college.ID).Update(college)
	logger.LogIfError(err)
}

// FindByID finds an college according to an ID
func (*CollegeDataAccessObject) FindByID(id int) (College, bool) {
	var college College
	has, err := orm.Table(CollegeDAO.TableName()).ID(id).Get(&college)
	logger.LogIfError(err)
	return college, has
}

//FindByName finds an college according to its name
func (*CollegeDataAccessObject) FindByName(name string) (College, bool) {
	var college College
	has, err := orm.Table(CollegeDAO.TableName()).Where("name=?", name).Get(&college)
	logger.LogIfError(err)
	return college, has

}

//DeleteByID delete an college according to an ID
func (*CollegeDataAccessObject) DeleteByID(id int) {
	var college College
	_, err := orm.Table(CollegeDAO.TableName()).ID(id).Delete(&college)
	logger.LogIfError(err)
}