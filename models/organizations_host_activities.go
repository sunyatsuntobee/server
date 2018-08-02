package models

import (
	"time"

	"github.com/sunyatsuntobee/server/logger"
)

// OrganizationsHostActivities Model
type OrganizationsHostActivities struct {
	ID             int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	OrganizationID int       `xorm:"organization_id INT NOTNULL INDEX(organization_id_idx)" json:"organization_id"`
    ActivityID     int       `xorm:"activity_id INT NOTNULL INDEX(activity_id_idx)" json:"activity_id"`
	Timestamp      time.Time `xorm:"timestamp DATETIME NOTNULL" json:"timestamp"`
}


// OrganizationsHostActivitiesFull Model
type OrganizationsHostActivitiesFull struct {
	ID             int          `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	Organization   Organization `xorm:"extends" json:"user"`
	HostedActivity Activity     `xorm:"extends" json:"activity"`
	Timestamp      time.Time    `xorm:"timestamp DATETIME NOTNULL" json:"timestamp"`
}

// OrganizationsHostActivitiesDataAccessObject provides database access for
// Model OrganizationsHostActivities
type OrganizationsHostActivitiesDataAccessObject struct{}

// OrganizationsHostActivitiesDAO instance of OrganizationsHostActivitiesDataAccessObject
var OrganizationsHostActivitiesDAO *OrganizationsHostActivitiesDataAccessObject

// NewOrganizationsHostActivities creates a new organization-activity relationship
func NewOrganizationsHostActivities(organizationID int, activityID int) *OrganizationsHostActivities {
	return &OrganizationsHostActivities{
		OrganizationID: organizationID,
		ActivityID:     activityID,
		Timestamp:      time.Now(),
	}
}

// TableName returns table name
func (*OrganizationsHostActivitiesDataAccessObject) TableName() string {
	return "organizations_host_activities"
}


// FindActivitiesByOrganizationID finds all full models by an organization ID
func (*OrganizationsHostActivitiesDataAccessObject) FindActivitiesByOrganizationID(
	organizationID int) []OrganizationsHostActivitiesFull {
	l := make([]OrganizationsHostActivitiesFull, 0)
	err := orm.Table(OrganizationsHostActivitiesDAO.TableName()).
		Join("INNER", OrganizationDAO.TableName(),
		    OrganizationsHostActivitiesDAO.TableName()+".organization_id="+
				OrganizationDAO.TableName()+".id").
		Join("INNER", ActivityDAO.TableName(),
		    OrganizationsHostActivitiesDAO.TableName()+".activity_id="+
				ActivityDAO.TableName()+".id").
		Where("organization_id=?", organizationID).
		Find(&l)
	logger.LogIfError(err)
	return l
}


// FindOrganizationsByActivityID finds all full models by an activity ID
func (*OrganizationsHostActivitiesDataAccessObject) FindOrganizationsByActivityID(
	activityID int) []OrganizationsHostActivitiesFull {
	l := make([]OrganizationsHostActivitiesFull, 0)
	err := orm.Table(OrganizationsHostActivitiesDAO.TableName()).
		Join("INNER", OrganizationDAO.TableName(),
		    OrganizationsHostActivitiesDAO.TableName()+".organization_id="+
				OrganizationDAO.TableName()+".id").
		Join("INNER", ActivityDAO.TableName(),
		    OrganizationsHostActivitiesDAO.TableName()+".activity_id="+
				ActivityDAO.TableName()+".id").
		Where("activity_id=?", activityID).
		Find(&l)
	logger.LogIfError(err)
	return l
}


// InsertOne insert a organization-activity relationship
func (*OrganizationsHostActivitiesDataAccessObject) InsertOne(
	organizationsHostActivities *OrganizationsHostActivities) {
	_, err := orm.Table(OrganizationsHostActivitiesDAO.TableName()).
		InsertOne(organizationsHostActivities)
	logger.LogIfError(err)
}

//DeleteByID delete a organization-activity relationship by its ID
func (*OrganizationsHostActivitiesDataAccessObject) DeleteByID(id int) {
	var organizationsHostActivities OrganizationsHostActivities
	_, err := orm.Table(OrganizationsHostActivitiesDAO.TableName()).
		ID(id).Delete(&organizationsHostActivities)
	logger.LogIfError(err)
}

//FindByID find a organization-activity relationship by its ID
func (*OrganizationsHostActivitiesDataAccessObject) FindByID(id int) (OrganizationsHostActivities, bool) {
	var oha OrganizationsHostActivities
	has, err := orm.Table(OrganizationsHostActivitiesDAO.TableName()).ID(id).Get(&oha)
	logger.LogIfError(err)
	return oha, has
}

