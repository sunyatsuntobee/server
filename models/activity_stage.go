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
	Location   string    `xorm:"location VARCHAR(20) NOTNULL" json:"location"`
	Content    string    `xorm:"content VARCHAR(400) NOTNULL" json:"content"`
	ActivityID int       `xorm:"activity_id INT NOTNULL INDEX(activity_id_idx)" json:"activity_id"`
}

type ActivityStageDataAccessObject struct{}

var ActivityStageDAO *ActivityStageDataAccessObject

func NewActivityStage(stageNum int, startTime time.Time, endTime time.Time,
	location string, content string, activityId int) *ActivityStage {
	return &ActivityStage{
		StageNum:   stageNum,
		StartTime:  startTime,
		EndTime:    endTime,
		Location:   location,
		Content:    content,
		ActivityID: activityId,
	}
}

func (*ActivityStageDataAccessObject) TableName() string {
	return "activity_stages"
}

func (*ActivityStageDataAccessObject) InsertOne(stage *ActivityStage) {
	_, err := orm.Table(ActivityStageDAO.TableName()).Insert(stage)
	logger.LogIfError(err)
}

func (*ActivityStageDataAccessObject) FindByAID(aid int) []ActivityStage {
	l := make([]ActivityStage, 0)
	err := orm.Table(ActivityStageDAO.TableName()).Where("activity_id=?", aid).
		Find(&l)
	logger.LogIfError(err)
	return l
}

func (*ActivityStageDataAccessObject) DeleteByAID(aid int) {
	var buf ActivityStage
	_, err := orm.Table(ActivityStageDAO.TableName()).
		Where("activity_id=?", aid).Unscoped().Delete(&buf)
	logger.LogIfError(err)
}
