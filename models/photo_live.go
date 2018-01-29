package models

import "github.com/sunyatsuntobee/server/logger"

// PhotoLive Model
type PhotoLive struct {
	ID                    int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	ExpectMembers         int    `xorm:"expect_members INT NOTNULL"`
	AdProgress            string `xorm:"ad_progress VARCHAR(100) NOTNULL"`
	ActivityStageID       int    `xorm:"activity_stage_id INT"`
	ManagerID             int    `xorm:"manager_id INT"`
	PhotographerManagerID int    `xorm:"photographer_manager_id INT"`
}

type PhotoLiveDataAccessObject struct{}

const PhotoLiveTableName string = "photo_lives"

var PhotoLiveDAO *PhotoLiveDataAccessObject

func (*PhotoLiveDataAccessObject) InsertOne(photoLive *PhotoLive) {
	_, err := orm.Table(PhotoLiveTableName).InsertOne(photoLive)
	logger.LogIfError(err)
}

func (*PhotoLiveDataAccessObject) FindAll() []PhotoLive {
	l := make([]PhotoLive, 0)
	err := orm.Table(PhotoLiveTableName).Find(&l)
	logger.LogIfError(err)
	return l
}

func (*PhotoLiveDataAccessObject) UpdateByID(id int, photoLive *PhotoLive) {
	_, err := orm.Table(PhotoLiveTableName).ID(id).Update(photoLive)
	logger.LogIfError(err)
}

func (*PhotoLiveDataAccessObject) DeleteByID(id int) {
	var photoLive PhotoLive
	_, err := orm.Table(PhotoLiveTableName).ID(id).Unscoped().Delete(&photoLive)
	logger.LogIfError(err)
}
