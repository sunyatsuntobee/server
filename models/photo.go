package models

import (
	"time"

	"github.com/sunyatsuntobee/server/logger"
)

// Photo Model
type Photo struct {
	ID             int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	URL            string    `xorm:"url VARCHAR(45) NOTNULL"`
	TookTime       time.Time `xorm:"took_time DATETIME NOTNULL"`
	TookLocation   string    `xorm:"took_location VARCHAR(45) NOTNULL"`
	ReleaseTime    time.Time `xorm:"release_time DATETIME"`
	Category       string    `xorm:"category VARCHAR(45) NOTNULL"`
	Likes          int       `xorm:"likes INT NOTNULL"`
	RejectReason   string    `xorm:"reject_reason VARCHAR(100)"`
	PhotographerID int       `xorm:"photographer_id INT NOTNULL INDEX(photographer_id_idx)"`
}

type PhotoDataAccessObject struct{}

const PhotoTableName string = "photos"

var PhotoDAO *PhotoDataAccessObject

func (*PhotoDataAccessObject) FindAll() []Photo {
	photos := make([]Photo, 0)
	err := orm.Table(PhotoTableName).Find(&photos)
	logger.LogIfError(err)
	return photos
}

func (*PhotoDataAccessObject) InsertOne(photo *Photo) {
	_, err := orm.Table(PhotoTableName).InsertOne(photo)
	logger.LogIfError(err)
}

func (*PhotoDataAccessObject) UpdateByID(id int, photo *Photo) {
	_, err := orm.Table(PhotoTableName).ID(id).Update(photo)
	logger.LogIfError(err)
}

func (*PhotoDataAccessObject) DeleteByID(id int) {
	var photo Photo
	_, err := orm.Table(PhotoTableName).ID(id).Unscoped().Delete(&photo)
	logger.LogIfError(err)
}
