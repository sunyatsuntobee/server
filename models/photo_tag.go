package models

import "github.com/sunyatsuntobee/server/logger"

// PhotoTag Model
type PhotoTag struct {
	ID      int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	Tag     string `xorm:"tag VARCHAR(20) NOTNULL"`
	PhotoID int    `xorm:"photo_id INT NOTNULL UNIQUE INDEX(photo_id_idx)"`
}

type PhotoTagDataAccessObject struct{}

const PhotoTagTableName string = "photo_tags"

var PhotoTagDAO *PhotoTagDataAccessObject

func NewPhotoTag(tag string, photo_id int) {
	return &PhotoTag{Tag:tag, PhotoID:photo_id}
}

func (*PhotoTagDataAccessObject) FindAll() []PhotoTag {
	l := make([]PhotoTag, 0)
	err := orm.Table(PhotoTagTableName).Find(&l)
	logger.LogIfError(err)
	return l
}

func (*PhotoTagDataAccessObject) InsertOne(photoTag *PhotoTag) {
	_, err := orm.Table(PhotoTagTableName).InsertOne(photoTag)
	logger.LogIfError(err)
}

func (*PhotoTagDataAccessObject) UpdateByID(id int, photoTag *PhotoTag) {
	_, err := orm.Table(PhotoTagTableName).ID(id).Update(&photoTag)
	logger.LogIfError(err)
}

func (*PhotoTagDataAccessObject) DeleteByID(id int) {
	var photoTag PhotoTag
	_, err := orm.Table(PhotoTagTableName).ID(id).Unscoped().Delete(&photoTag)
	logger.LogIfError(err)
}
