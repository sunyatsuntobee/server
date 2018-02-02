package models

import "time"

// ActivityStage Model
type ActivityStage struct {
	ID         int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	StageNum   int       `xorm:"stage_num INT NOTNULL"`
	StartTime  time.Time `xorm:"start_time DATETIME NOTNULL"`
	EndTime    time.Time `xorm:"end_time DATETIME NOTNULL"`
	Location   string    `xorm:"location VARCHAR(45) NOTNULL"`
	Content    string    `xorm:"content VARCHAR(200) NOTNULL"`
	ActivityID int       `xorm:"activity_id INT NOTNULL INDEX(activity_id_idx)"`
}

type ActivityStageDataAccessObject struct{}

var ActivityStageDAO *ActivityStageDataAccessObject

func (*ActivityStageDataAccessObject) TableName() string {
	return "activity_stages"
}
