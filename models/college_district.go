package models

import "github.com/sunyatsuntobee/server/logger"

//college_district model
type CollegeDistrict struct {
	ID           int     `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	Name         string  `xorm:"name VARCHAR(50) NOTNULL" json:"name"`
	CollegeID    int     `xorm:"college_id INT NOTNULL INDEX(college_id_idx)" json:"college_id"`
}

//CollegeDistrictDataAccessObject provides database access for Model CollegeDistrict
type CollegeDistrictDataAccessObject struct{}

// CollegeDistrictDAO instance of CollegeDistrictDataAccessObject
var CollegeDistrictDAO *CollegeDistrictDataAccessObject

// TableName returns table name
func (*CollegeDistrictDataAccessObject) TableName() string {
	return "college_districts"
}

// NewCollege creates a new collegeDistrict
func NewCollegeDistrict(name string, college_id int) *CollegeDistrict{
	return &CollegeDistrict{
		Name:  name,
        CollegeID: college_id,
	}
}

// FindAll finds all colleges
func (*CollegeDistrictDataAccessObject) FindAll() []CollegeDistrict {
	collegeDistricts := make([]CollegeDistrict, 0)
	err := orm.Table(CollegeDistrictDAO.TableName()).Find(&collegeDistricts)
	logger.LogIfError(err)
	return colleges
}


// InsertOne inserts a CollegeDistrict
func (*CollegeDistrictDataAccessObject) InsertOne(collegeDistrict *CollegeDistrict) {
	_, err := orm.Table(CollegeDistrictDAO.TableName()).InsertOne(collegeDistrict)
	logger.LogIfError(err)
}

// UpdateOne updates a CollegeDistrict
func (*CollegeDistrictDataAccessObject) UpdateOne(collegeDistrict *CollegeDistrict) {
	_, err := orm.Table(CollegeDistrictDAO.TableName()).ID(collegeDistrict.ID).Update(collegeDistrict)
	logger.LogIfError(err)
}

// FindByID finds a collegeDistrict according to an ID
func (*CollegeDistrictDataAccessObject) FindByID(id int) (CollegeDistrict, bool) {
	var collegeDistrict CollegeDistrict
	has, err := orm.Table(CollegeDistrictDAO.TableName()).ID(id).Get(&collegeDistrict)
	logger.LogIfError(err)
	return collegeDistrict, has
}

//DeleteByID delete a collegeDistrict according to its ID
func (*CollegeDistrictDataAccessObject) DeleteByID(id int) {
	var collegeDistrict CollegeDistrict
	_, err := orm.Table(CollegeDistrictDAO.TableName()).ID(id).Delete(&collegeDistrict)
	logger.LogIfError(err)
}

// DeleteByCID deletes all college district according to an college ID
func (*CollegeDistrictDataAccessObject) DeleteByCID(cid int) {
	var buf CollegeDistrict
	_, err := orm.Table(CollegeDistrictDAO.TableName()).
		Where("college_id=?", cid).
		Unscoped().Delete(&buf)
	logger.LogIfError(err)
}

// FindByCID find all college district according to an college ID
func (*CollegeDistrictDataAccessObject) HasCreate(cid int, name string) bool {
	var buf CollegeDistrict
	has, err := orm.Table(CollegeDistrictDAO.TableName()).
		Where("college_id=?", cid).And("name=?", name).
		Get(&buf)
	logger.LogIfError(err)
	return has;
}