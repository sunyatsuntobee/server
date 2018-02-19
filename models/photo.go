package models

import (
	"time"

	"github.com/sunyatsuntobee/server/logger"
)

// Photo Model
type Photo struct {
	ID             int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	URL            string    `xorm:"url VARCHAR(59) NOTNULL" json:"url"`
	TookTime       time.Time `xorm:"took_time DATETIME NOTNULL" json:"took_time"`
	TookLocation   string    `xorm:"took_location VARCHAR(50) NOTNULL" json:"took_location"`
	ReleaseTime    time.Time `xorm:"release_time DATETIME" json:"release_time"`
	Category       string    `xorm:"category VARCHAR(20) NOTNULL" json:"category"`
	Likes          int       `xorm:"likes INT NOTNULL" json:"likes"`
	RejectReason   string    `xorm:"reject_reason VARCHAR(200)" json:"reject_reason"`
	PhotographerID int       `xorm:"photographer_id INT NOTNULL INDEX(photographer_id_idx)" json:"photographer_id"`
}

type PhotoDataAccessObject struct{}

var PhotoDAO *PhotoDataAccessObject

func NewPhoto(url string, tookTime time.Time, tookLocation string,
	releaseTime time.Time, category string, likes int,
	rejectReason string, photographerId int) *Photo {
	return &Photo{
		URL:            url,
		TookTime:       tookTime,
		TookLocation:   tookLocation,
		ReleaseTime:    releaseTime,
		Category:       category,
		Likes:          likes,
		RejectReason:   rejectReason,
		PhotographerID: photographerId,
	}
}

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

func (*PhotoDataAccessObject) UpdateOne(photo *Photo) {
	_, err := orm.Table(PhotoDAO.TableName()).ID(photo.ID).Update(photo)
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

func (*PhotoDataAccessObject) FindFullAllWithoutUnchecked() []PhotoFull {
	l := make([]PhotoFull, 0)
	err := orm.Table(PhotoDAO.TableName()).
		Join("INNER", UserDAO.TableName(), "photos.photographer_id=users.id").
		Where("photos.category!=?", "未审核").Find(&l)
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
