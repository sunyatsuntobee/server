package models

import "github.com/sunyatsuntobee/server/logger"

// Activity Model
type Activity struct {
	ID             int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	Name           string `xorm:"name VARCHAR(45) NOTNULL" json:"name"`
	Description    string `xorm:"description VARCHAR(100) NOTNULL" json:"description"`
	Category       string `xorm:"category VARCHAR(45) NOTNULL" json:"category"`
	PosterURL      string `xorm:"poster_url VARCHAR(45)" json:"poster_url"`
	LogoURL        string `xorm:"logo_url VARCHAR(45)" json:"logo_url"`
	OrganizationID int    `xorm:"organization_id INT NOTNULL INDEX(organization_id_idx)" json:"organization_id"`
}

type ActivityDataAccessObject struct{}

var ActivityDAO *ActivityDataAccessObject

func NewActivity(name string, description string, category string, poster_url string, logo_url string, organization_id int) {
	return &Activity{Name:name, Description:description, Category:category, PosterURL:poster_url, LogoURL:logo_url, OrganizationID:organization_id}
}

func (*ActivityDataAccessObject) TableName() string {
	return "activities"
}

func (*ActivityDataAccessObject) UpdateOne(activity *Activity) {
	_, err := orm.Table(ActivityDAO.TableName()).ID(activity.ID).
		Update(activity)
	logger.LogIfError(err)
}

func (*ActivityDataAccessObject) FindByID(id int) (Activity, bool) {
	var activity Activity
	has, err := orm.Table(ActivityDAO.TableName()).ID(id).Get(&activity)
	logger.LogIfError(err)
	return activity, has
}

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

func (*ActivityDataAccessObject) FindByOID(oid int) []Activity {
	activities := make([]Activity, 0)
	err := orm.Table(ActivityDAO.TableName()).Where("organization_id=?", oid).
		Find(&activities)
	logger.LogIfError(err)
	return activities
}
