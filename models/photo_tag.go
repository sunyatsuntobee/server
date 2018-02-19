package models

import "github.com/sunyatsuntobee/server/logger"

// PhotoTag Model
type PhotoTag struct {
	ID      int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	Tag     string `xorm:"tag VARCHAR(20) NOTNULL"`
	PhotoID int    `xorm:"photo_id INT NOTNULL UNIQUE INDEX(photo_id_idx)"`
}

// NewPhotoTag creates a new photo tag
func NewPhotoTag(tag string, photoID int) *PhotoTag {
	return &PhotoTag{
		Tag:     tag,
		PhotoID: photoID,
	}
}

// PhotoTagDataAccessObject provides database access for Model
// PhotoTag
type PhotoTagDataAccessObject struct{}

// PhotoTagDAO instance of PhotoTagDataAccessObject
var PhotoTagDAO *PhotoTagDataAccessObject

// TableName returns table name
func (*PhotoTagDataAccessObject) TableName() string {
	return "photo_tags"
}

// FindAll finds all photo tags
func (*PhotoTagDataAccessObject) FindAll() []PhotoTag {
	l := make([]PhotoTag, 0)
	err := orm.Table(PhotoTagDAO.TableName()).Find(&l)
	logger.LogIfError(err)
	return l
}

// InsertOne inserts a photo tag
func (*PhotoTagDataAccessObject) InsertOne(photoTag *PhotoTag) {
	_, err := orm.Table(PhotoTagDAO.TableName()).InsertOne(photoTag)
	logger.LogIfError(err)
}

// UpdateByID updates a photo tag
func (*PhotoTagDataAccessObject) UpdateByID(photoTag *PhotoTag) {
	_, err := orm.Table(PhotoTagDAO.TableName()).ID(photoTag.ID).
		Update(&photoTag)
	logger.LogIfError(err)
}

// DeleteByID deletes a photo tag by ID
func (*PhotoTagDataAccessObject) DeleteByID(id int) {
	var photoTag PhotoTag
	_, err := orm.Table(PhotoTagDAO.TableName()).ID(id).Delete(&photoTag)
	logger.LogIfError(err)
}
