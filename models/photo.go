package models

import (
	"time"

	"github.com/sunyatsuntobee/server/logger"
)

// Photo Model
type Photo struct {
	ID             int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	URL            string    `xorm:"url VARCHAR(50) NOTNULL" json:"url"`
	TookTime       time.Time `xorm:"took_time DATETIME NOTNULL" json:"took_time"`
	TookLocation   string    `xorm:"took_location VARCHAR(50) NOTNULL" json:"took_location"`
	ReleaseTime    time.Time `xorm:"release_time DATETIME" json:"release_time"`
	Category       string    `xorm:"category VARCHAR(20) NOTNULL" json:"category"`
	Likes          int       `xorm:"likes INT NOTNULL" json:"likes"`
	RejectReason   string    `xorm:"reject_reason VARCHAR(200)" json:"reject_reason"`
	UserID         int       `xorm:"user_id INT NOTNULL INDEX(photographer_id_idx)" json:"user_id"`
	OrganizationID int       `xorm:"organization_id INT INDEX(fk_photos_organization_id_idx)" json:"organization_id"`
}

// NewPhoto creates a new photo
func NewPhoto(tookTime time.Time, tookLocation string,
	userID int) *Photo {
	return &Photo{
		URL:          "",
		TookTime:     tookTime,
		TookLocation: tookLocation,
		ReleaseTime:  time.Time{},
		Category:     "未审核",
		Likes:        0,
		RejectReason: "",
		UserID:       userID,
	}
}

// PhotoDataAccessObject provides database access for Model Photo
type PhotoDataAccessObject struct{}

// PhotoDAO instance of PhotoDataAccessObject
var PhotoDAO *PhotoDataAccessObject

// TableName returns table name
func (*PhotoDataAccessObject) TableName() string {
	return "photos"
}

// FindByOID finds photos by organization ID
func (*PhotoDataAccessObject) FindByOID(oid int) []Photo {
	photos := make([]Photo, 0)
	err := orm.Table(PhotoDAO.TableName()).Where("organization_id=?", oid).
		Find(&photos)
	logger.LogIfError(err)
	return photos
}

// FindAll finds all photos
func (*PhotoDataAccessObject) FindAll() []Photo {
	photos := make([]Photo, 0)
	err := orm.Table(PhotoDAO.TableName()).Find(&photos)
	logger.LogIfError(err)
	return photos
}

// InsertOne inserts a photo
func (*PhotoDataAccessObject) InsertOne(photo *Photo) {
	_, err := orm.Table(PhotoDAO.TableName()).InsertOne(photo)
	logger.LogIfError(err)
}

// UpdateOne updates a photo
func (*PhotoDataAccessObject) UpdateOne(photo *Photo) {
	_, err := orm.Table(PhotoDAO.TableName()).ID(photo.ID).Update(photo)
	logger.LogIfError(err)
}

// DeleteByID deletes a photo according to ID
func (*PhotoDataAccessObject) DeleteByID(id int) {
	var photo Photo
	_, err := orm.Table(PhotoDAO.TableName()).ID(id).Delete(&photo)
	logger.LogIfError(err)
}

// FindByID finds a photo by ID
func (*PhotoDataAccessObject) FindByID(id int) (Photo, bool) {
	var photo Photo
	has, err := orm.Table(PhotoDAO.TableName()).ID(id).Get(&photo)
	logger.LogIfError(err)
	return photo, has
}

// FindByCategory finds photos by category
func (*PhotoDataAccessObject) FindByCategory(category string) []Photo {
	l := make([]Photo, 0)
	err := orm.Table(PhotoDAO.TableName()).Where("category=", category).Find(&l)
	logger.LogIfError(err)
	return l
}

// FindFullByCategory finds joined photos according to category
func (*PhotoDataAccessObject) FindFullByCategory(category string) []PhotoFull {
	l := make([]PhotoFull, 0)
	err := orm.Table(PhotoDAO.TableName()).
		Join("INNER", UserDAO.TableName(), "photos.photographer_id=users.id").
		Where("photos.category=?", category).Find(&l)
	logger.LogIfError(err)
	return l
}

// FindFullAll finds all joined photos
func (*PhotoDataAccessObject) FindFullAll() []PhotoFull {
	l := make([]PhotoFull, 0)
	err := orm.Table(PhotoDAO.TableName()).
		Join("INNER", UserDAO.TableName(), "photos.photographer_id=users.id").
		Find(&l)
	logger.LogIfError(err)
	return l
}

// FindFullAllWithoutUnchecked finds all checked joined photos
func (*PhotoDataAccessObject) FindFullAllWithoutUnchecked() []PhotoFull {
	l := make([]PhotoFull, 0)
	err := orm.Table(PhotoDAO.TableName()).
		Join("INNER", UserDAO.TableName(), "photos.photographer_id=users.id").
		Where("photos.category!=?", "未审核").Find(&l)
	logger.LogIfError(err)
	return l
}

// FindFullByID finds a joined photo by ID
func (*PhotoDataAccessObject) FindFullByID(id int) (PhotoFull, bool) {
	var photo PhotoFull
	has, err := orm.Table(PhotoDAO.TableName()).
		Join("INNER", UserDAO.TableName(), "photos.photographer_id=users.id").
		Where("photos.id=?", id).Get(&photo)
	logger.LogIfError(err)
	return photo, has
}
