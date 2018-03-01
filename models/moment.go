package models

import (
	"time"

	"github.com/sunyatsuntobee/server/logger"
)

// Moment Model
type Moment struct {
	ID             int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	UserID         int       `xorm:"user_id INT NOTNULL INDEX(fk_moments_user_id_idx)" json:"user_id"`
	OrganizationID int       `xorm:"organization_id INT INDEX(fk_moments_organization_id_idx)" json:"organization_id"`
	Timestamp      time.Time `xorm:"timestamp DATETIME NOTNULL" json:"timestamp"`
	Content        string    `xorm:"content VARCHAR(400) NOTNULL" json:"content"`
	Photo1ID       int       `xorm:"photo_1_id INT NOTNULL INDEX(fk_moments_photo_1_id_idx)" json:"photo_1_id"`
	Photo2ID       int       `xorm:"photo_2_id INT INDEX(fk_moments_photo_2_id_idx)" json:"photo_2_id"`
	Photo3ID       int       `xorm:"photo_3_id INT INDEX(fk_moments_photo_3_id_idx)" json:"photo_3_id"`
	Photo4ID       int       `xorm:"photo_4_id INT INDEX(fk_moments_photo_4_id_idx)" json:"photo_4_id"`
	Photo5ID       int       `xorm:"photo_5_id INT INDEX(fk_moments_photo_5_id_idx)" json:"photo_5_id"`
	Photo6ID       int       `xorm:"photo_6_id INT INDEX(fk_moments_photo_6_id_idx)" json:"photo_6_id"`
	Photo7ID       int       `xorm:"photo_7_id INT INDEX(fk_moments_photo_7_id_idx)" json:"photo_7_id"`
	Photo8ID       int       `xorm:"photo_8_id INT INDEX(fk_moments_photo_8_id_idx)" json:"photo_8_id"`
	Photo9ID       int       `xorm:"photo_9_id INT INDEX(fk_moments_photo_9_id_idx)" json:"photo_9_id"`
}

// MomentDataAccessObject provides database access for
// Model Moment
type MomentDataAccessObject struct{}

// MomentDAO instance of MomentDataAccessObject
var MomentDAO *MomentDataAccessObject

// TableName returns table name
func (*MomentDataAccessObject) TableName() string {
	return "moment"
}

// FindByUserID finds a moment by user id
func (*MomentDataAccessObject) FindByUserID(user_id int) Moment {
	var moment Moment
	err := orm.Table(MomentDAO.TableName()).
		Where("moment.user_id=?", user_id).
		Find(&moment)
	logger.LogIfError(err)
	return moment
}

// InsertOne inserts a moment
func (*MomentDataAccessObject) InsertOne(moment *Moment) {
	_, err := orm.Table(MomentDAO.TableName()).InsertOne(moment)
	logger.LogIfError(err)
}
