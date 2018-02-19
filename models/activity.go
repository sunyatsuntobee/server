package models

import "github.com/sunyatsuntobee/server/logger"

// Activity Model
type Activity struct {
	ID             int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	Name           string `xorm:"name VARCHAR(20) NOTNULL" json:"name"`
	Description    string `xorm:"description VARCHAR(200) NOTNULL" json:"description"`
	Category       string `xorm:"category VARCHAR(10) NOTNULL" json:"category"`
	PosterURL      string `xorm:"poster_url VARCHAR(50)" json:"poster_url"`
	LogoURL        string `xorm:"logo_url VARCHAR(50)" json:"logo_url"`
	OrganizationID int    `xorm:"organization_id INT NOTNULL INDEX(organization_id_idx)" json:"organization_id"`
}

// ActivityDataAccessObject provides database access for Model Activity
type ActivityDataAccessObject struct{}

// ActivityDAO instance of ActivityDataAccessObject
var ActivityDAO *ActivityDataAccessObject

// NewActivity creates a new activity
func NewActivity(name string, description string, category string,
	posterURL string, logoURL string, organizationID int) *Activity {
	return &Activity{
		Name:           name,
		Description:    description,
		Category:       category,
		PosterURL:      posterURL,
		LogoURL:        logoURL,
		OrganizationID: organizationID,
	}
}

// TableName returns table name
func (*ActivityDataAccessObject) TableName() string {
	return "activities"
}

// UpdateOne updates an activity
func (*ActivityDataAccessObject) UpdateOne(activity *Activity) {
	_, err := orm.Table(ActivityDAO.TableName()).ID(activity.ID).
		Update(activity)
	logger.LogIfError(err)
}

// FindByID finds an activity according to an ID
func (*ActivityDataAccessObject) FindByID(id int) (Activity, bool) {
	var activity Activity
	has, err := orm.Table(ActivityDAO.TableName()).ID(id).Get(&activity)
	logger.LogIfError(err)
	return activity, has
}

// FindFullByID finds joined activities according to an ID
func (*ActivityDataAccessObject) FindFullByID(id int) []ActivityFull {
	l := make([]ActivityFull, 0)
	err := orm.Table(ActivityDAO.TableName()).
		Join("INNER", ActivityStageDAO.TableName(),
			"activities.id=activity_stages.activity_id").
		Join("INNER", OrganizationDAO.TableName(),
			"activities.organization_id=organizations.id").
		Where("activities.id=?", id).Asc("stage_num").
		Find(&l)
	logger.LogIfError(err)
	return l
}

// FindFullByOID finds joined activities according to an organization ID
func (*ActivityDataAccessObject) FindFullByOID(oid int) []ActivityFull {
	activities := make([]ActivityFull, 0)
	err := orm.Table(ActivityDAO.TableName()).
		Join("INNER", ActivityStageDAO.TableName(),
			"activities.id=activity_stages.activity_id").
		Join("INNER", OrganizationDAO.TableName(),
			"activities.organization_id=organizations.id").
		Where("organizations.id=?", oid).Asc("stage_num").
		Find(&activities)
	logger.LogIfError(err)
	return activities
}

// FindFullAll finds all joined activities
func (*ActivityDataAccessObject) FindFullAll() []ActivityFull {
	activities := make([]ActivityFull, 0)
	err := orm.Table(ActivityDAO.TableName()).
		Join("INNER", ActivityStageDAO.TableName(),
			"activities.id=activity_stages.activity_id").
		Join("INNER", OrganizationDAO.TableName(),
			"activities.organization_id=organizations.id").
		Asc("stage_num").
		Find(&activities)
	logger.LogIfError(err)
	return activities
}

// FindByOID finds activities according to an organization ID
func (*ActivityDataAccessObject) FindByOID(oid int) []Activity {
	activities := make([]Activity, 0)
	err := orm.Table(ActivityDAO.TableName()).Where("organization_id=?", oid).
		Find(&activities)
	logger.LogIfError(err)
	return activities
}
