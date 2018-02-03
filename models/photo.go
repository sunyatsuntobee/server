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

var PhotoDAO *PhotoDataAccessObject

func (*PhotoDataAccessObject) TableName() string {
	return "photos"
}

func (*PhotoDataAccessObject) FindAll() []Photo {
	photos := make([]Photo, 0)
	err := orm.Table(PhotoDAO.TableName()).Find(&photos)
	logger.LogIfError(err)
	return photos
}

func (*PhotoDataAccessObject) InsertOne(photo *Photo) {
	_, err := orm.Table(PhotoDAO.TableName()).InsertOne(photo)
	logger.LogIfError(err)
}

func (*PhotoDataAccessObject) UpdateByID(id int, photo *Photo) {
	_, err := orm.Table(PhotoDAO.TableName()).ID(id).Update(photo)
	logger.LogIfError(err)
}

func (*PhotoDataAccessObject) DeleteByID(id int) {
	var photo Photo
	_, err := orm.Table(PhotoDAO.TableName()).ID(id).Unscoped().Delete(&photo)
	logger.LogIfError(err)
}

func (*PhotoDataAccessObject) FindByID(id int) (Photo, bool) {
	var photo Photo
	has, err := orm.Table(PhotoDAO.TableName()).ID(id).Get(&photo)
	logger.LogIfError(err)
	return photo, has
}

func (*PhotoDataAccessObject) FindByCategory(category string) []Photo {
	l := make([]Photo, 0)
	err := orm.Table(PhotoDAO.TableName()).Where("category=", category).Find(&l)
	logger.LogIfError(err)
	return l
}

func (*PhotoDataAccessObject) FindFullByCategory(category string) []PhotoFull {
	l := make([]PhotoFull, 0)
	err := orm.Table(PhotoDAO.TableName()).
		Join("INNER", UserDAO.TableName(), "photos.photographer_id=users.id").
		Where("photos.category=?", category).Find(&l)
	logger.LogIfError(err)
	return l
}

func (*PhotoDataAccessObject) FindFullAll() []PhotoFull {
	l := make([]PhotoFull, 0)
	err := orm.Table(PhotoDAO.TableName()).
		Join("INNER", UserDAO.TableName(), "photos.photographer_id=users.id").
		Find(&l)
	logger.LogIfError(err)
	return l
}

func (*PhotoDataAccessObject) FindFullByID(id int) (PhotoFull, bool) {
	var photo PhotoFull
	has, err := orm.Table(PhotoDAO.TableName()).
		Join("INNER", UserDAO.TableName(), "photos.photographer_id=users.id").
		Where("photos.id=?", id).Get(&photo)
	logger.LogIfError(err)
	return photo, has
}
