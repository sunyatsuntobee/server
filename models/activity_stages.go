package models

import "time"

type Activity_stages struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	StageNum int `xorm:"stage_num INT NOTNULL"`
	StartTime time.Time `xorm:"start_time DATETIME NOTNULL"`
	EndTime time.Time `xorm:"end_time DATETIME NOTNULL"`
	Location string `xorm:"location VARCHAR(45) NOTNULL"`
	Content string `xorm:"content VARCHAR(45) NOTNULL"`
	ActivityID int `xorm:"activity_id INT NOTNULL INDEX(activity_id_idx)"`
}
