package models

import (
	"time"

	"github.com/sunyatsuntobee/server/logger"
)

// ActivityStage Model
type ActivityStage struct {
	ID         int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	StageNum   int       `xorm:"stage_num INT NOTNULL" json:"stage_num"`
	StartTime  time.Time `xorm:"start_time DATETIME NOTNULL" json:"start_time"`
	EndTime    time.Time `xorm:"end_time DATETIME NOTNULL" json:"end_time"`
	Location   string    `xorm:"location VARCHAR(50) NOTNULL" json:"location"`
	Content    string    `xorm:"content VARCHAR(400) NOTNULL" json:"content"`
	WechatURL  string 	 `xorm:"wechat_url VARCHAR(50)" json:"wechat_url"`
	SignUpURL  string    `xorm:"signup_url VARCHAR(50)" json:"signup_url"`
	ActivityID int       `xorm:"activity_id INT NOTNULL INDEX(activity_id_idx)" json:"activity_id"`
}
//参与环节可以额外增加时间段， 分活动环节，核心环节，参与环节（之后有瞬间时间点）
//管理员的总数等于社团的部门数+3
//几个按钮都是管理员的权限

// NewActivityStage creates a new activity stage
func NewActivityStage(stageNum int, startTime time.Time, endTime time.Time,
	location string, content string, wechatURL string, activityID int) *ActivityStage {
	return &ActivityStage{
		StageNum:   stageNum,
		StartTime:  startTime,
		EndTime:    endTime,
		Location:   location,
		Content:    content,
		WechatURL:  wechatURL,
		ActivityID: activityID,
	}
}

// ActivityStageDataAccessObject provides database access for ActivityStage
type ActivityStageDataAccessObject struct{}

// ActivityStageDAO instance of ActivityStageDataAccessObject
var ActivityStageDAO *ActivityStageDataAccessObject

// TableName returns the table name
func (*ActivityStageDataAccessObject) TableName() string {
	return "activity_stages"
}

// InsertOne inserts a new activity stage
func (*ActivityStageDataAccessObject) InsertOne(stage *ActivityStage) {
	_, err := orm.Table(ActivityStageDAO.TableName()).Insert(stage)
	logger.LogIfError(err)
}

// FindByAID finds activity stages of an activity
func (*ActivityStageDataAccessObject) FindByAID(aid int) []ActivityStage {
	l := make([]ActivityStage, 0)
	err := orm.Table(ActivityStageDAO.TableName()).Where("activity_id=?", aid).
		Find(&l)
	
	logger.LogIfError(err)
	return l
}

// FindAll find all activity stages
func (*ActivityStageDataAccessObject) FindAll() []ActivityStage {
	l := make([]ActivityStage, 0)
	err := orm.Table(ActivityStageDAO.TableName()).Find(&l)

	logger.LogIfError(err)
	return l
}

// DeleteByAID deletes all activity stages of an activity
func (*ActivityStageDataAccessObject) DeleteByAID(aid int) {
	var buf ActivityStage
	_, err := orm.Table(ActivityStageDAO.TableName()).
		Where("activity_id=?", aid).Delete(&buf)
	logger.LogIfError(err)
}

//UpdateOne update an activity stage
func (*ActivityStageDataAccessObject) UpdateOne(activityStage *ActivityStage) {
	_, err := orm.Table(ActivityStageDAO.TableName()).ID(activityStage.ID).
		Update(activityStage)
	logger.LogIfError(err)
}

// FindFullByDay finds all joined activity stages on a day
func (*ActivityStageDataAccessObject) FindFullByDay(
	date time.Time) []ActivityStageFull {

	startDate := time.Date(date.Year(), date.Month(), date.Day(),
		0, 0, 0, 0, time.Local)
	endDate := startDate.AddDate(0, 0, 1)
	start := startDate.Format(mysqlTimeFormat)
	end := time.Date(endDate.Year(), endDate.Month(), endDate.Day(),
		0, 0, 0, 0, time.Local).Format(mysqlTimeFormat)
	result := make([]ActivityStageFull, 0)

	err := orm.Table(ActivityDAO.TableName()).
		Where("activity_stages.activity_id=activities.id").
		Join("INNER", ActivityStageDAO.TableName(),
		"start_time>=?", start).And("end_time<?", end).
		Find(&result)
	logger.LogIfError(err)
	return result
}

// FindFullByMonth finds all joined activity stages in a month
func (*ActivityStageDataAccessObject) FindFullByMonth(
	date time.Time) [][]ActivityStageFull {

	start := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.Local)
	end := start.AddDate(0, 1, 0).AddDate(0, 0, -1)
	result := make([][]ActivityStageFull, 0)
	for i := start.Day(); i <= end.Day(); i++ {
		curDate := time.Date(date.Year(), date.Month(), i, 0, 0, 0, 0,
			time.Local)
		result = append(result, ActivityStageDAO.FindFullByDay(curDate))
	}

	return result
}

// FindByID finds activityStage by ID
func (*ActivityStageDataAccessObject) FindByID(id int) (ActivityStage, bool) {
	var result ActivityStage
	has, err := orm.Table(ActivityStageDAO.TableName()).ID(id).Get(&result) 
	logger.LogIfError(err)
	return result, has
}