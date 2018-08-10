package models

import (
	"github.com/sunyatsuntobee/server/logger"
	"fmt"
)

// Activity Model
type Activity struct {
	ID                 int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	ShortName          string `xorm:"short_name VARCHAR(20)" json:"short_name"`
	Name               string `xorm:"name VARCHAR(50) NOTNULL" json:"name"`
	Description        string `xorm:"description VARCHAR(200) NOTNULL" json:"description"`
	Category           string `xorm:"category VARCHAR(10) NOTNULL" json:"category"`
	CollegeDistrict    int    `xorm:"college_district INT NOTNULL" json:"college_district"`
	School			   string `xorm:"school VARCHAR(45) " json:"school"`
	Range			   string `xorm:"range VARCHAR(45)" json:"range"`
	Type			   string `xorm:"type VARCHAR(45)" json:"type"`
	PosterURL          string `xorm:"poster_url VARCHAR(50)" json:"poster_url"`
	LogoURL            string `xorm:"logo_url VARCHAR(50)" json:"logo_url"`
	WechatURL          string `xorm:"wechat_url VARCHAR(50)" json:"wechat_url"`
	LiveURL			   string `xorm:"live_url VARCHAR(50)" json:"live_url"`
	//SignUpURL		   string `xorm:"signup_url VARCHAR(50)" json:"signup_url"`
	SportsMedals       string `xorm:"sports_medals VARCHAR(50)" json:"sports_medals"`
	PublicServiceHours string `xorm:"public_service_hours VARCHAR(50)" json:"public_service_hours"`
	Prize              string `xorm:"prize VARCHAR(100)" json:"prize"`
	OtherPrize         string `xorm:"other_prize VARCHAR(100)" json:"other_prize"`
	AttentionNum       int    `xorm:"attention_num INT NOTNULL" json:"attention_num"`
}

// ActivityDataAccessObject provides database access for Model Activity
type ActivityDataAccessObject struct{}

// ActivityDAO instance of ActivityDataAccessObject
var ActivityDAO *ActivityDataAccessObject

// NewActivity creates a new activity
func NewActivity(shortName, name string, description string, category string,
	posterURL string, logoURL string, wechatURL string,
	college_district int,
	sportsMedals string, publicServiceHours string, prize string,
	otherPrize string, attentionNum int, liveURL string) *Activity {
	return &Activity{
		ShortName:          shortName,
		Name:               name,
		Description:        description,
		Category:           category,
	    CollegeDistrict:    college_district,
		PosterURL:          posterURL,
		LogoURL:            logoURL,
		WechatURL:          wechatURL,
		SportsMedals:       sportsMedals,
		PublicServiceHours: publicServiceHours,
		Prize:              prize,
		OtherPrize:         otherPrize,
		AttentionNum:       attentionNum,
		LiveURL:            liveURL,
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

/*
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
}*/

/*
// FindByOID finds activities according to an organization ID
func (*ActivityDataAccessObject) FindByOID(oid int) []ActivityAndStage {
	l := make([]ActivityAndStage, 0)
	err := orm.Table(ActivityDAO.TableName()).Where("organization_id=?", oid).Find(&l)
	for i := 0; i < len(l); i++ {
		l[i].Stages  = make([]ActivityStage, 0)
		err = orm.Table(ActivityStageDAO.TableName()).Where("activity_id=?", l[i].Activity.ID).Find(&l[i].Stages)
	}
	logger.LogIfError(err)
	return l
}
*/

// InsertOne inserts an activity
func (*ActivityDataAccessObject) InsertOne(activity *Activity) {
	fmt.Println("start")
	_, err := orm.Table(ActivityDAO.TableName()).InsertOne(activity)
	fmt.Println("end")
	logger.LogIfError(err)
}

// FindFullByAID finds activities and stages according to an actID
func (*ActivityDataAccessObject) FindFullByAID(id int) (ActivityAndStage, bool) {
	var l ActivityAndStage

	has, err := orm.Table(ActivityDAO.TableName()).ID(id).Get(&l.Activity)

	l.Stages = make([]ActivityStage, 0)
	err = orm.Table(ActivityStageDAO.TableName()).Where("activity_id=?", id).
		Find(&l.Stages)
	logger.LogIfError(err)
	return l, has
}

// FindAll finds all activities
func (*ActivityDataAccessObject) FindAll() []Activity {
	activities := make([]Activity, 0)
	err := orm.Table(ActivityDAO.TableName()).
		Find(&activities)
	logger.LogIfError(err)
	return activities
}

// FindByCDID find all activities with CollegeDistrict ID
func (*ActivityDataAccessObject) FindByCDID(cdid int) []Activity {
	activities := make([]Activity, 0)
	err := orm.Table(ActivityDAO.TableName()).Where("college_district=?", cdid).
	Find(&activities)
	logger.LogIfError(err)
	return activities;
}
