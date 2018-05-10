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
	ActivityID int       `xorm:"activity_id INT NOTNULL INDEX(activity_id_idx)" json:"activity_id"`
}

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

//updateOne
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
	err := orm.Table(ActivityStageDAO.TableName()).
		Where("start_time>=?", start).And("end_time<?", end).
		Join("INNER", ActivityDAO.TableName(),
			"activity_stages.activity_id=activities.id").
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


