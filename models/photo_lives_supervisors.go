package models

import "github.com/sunyatsuntobee/server/logger"

// PhotoLivesSupervisors Model
type PhotoLivesSupervisors struct {
	ID           int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	PhotoLiveID  int `xorm:"photo_live_id INT NOTNULL INDEX(photo_live_id_idx)"`
	SupervisorID int `xorm:"supervisor_id INT NOTNULL INDEX(supervisor_id_idx)"`
}

// PhotoLivesSupervisorsDataAccessObject provides database access for
// Model PhotoLivesSupervisorsDataAccessObject
type PhotoLivesSupervisorsDataAccessObject struct{}

// PhotoLivesSupervisorsDAO instance of PhotoLivesSupervisorsDataAccessObject
var PhotoLivesSupervisorsDAO *PhotoLivesSupervisorsDataAccessObject

// TableName returns table name
func (*PhotoLivesSupervisorsDataAccessObject) TableName() string {
	return "photo_lives_supervisors"
}

// NewPhotoLivesSupervisors creates a new photo live-
// supervisor relationship
func NewPhotoLivesSupervisors(photoLiveID int,
	supervisorID int) *PhotoLivesSupervisors {
	return &PhotoLivesSupervisors{
		PhotoLiveID:  photoLiveID,
		SupervisorID: supervisorID,
	}
}

// InsertOne inserts a relationship
func (*PhotoLivesSupervisorsDataAccessObject) InsertOne(
	r *PhotoLivesSupervisors) {
	_, err := orm.Table(PhotoLivesSupervisorsDAO.TableName()).InsertOne(r)
	logger.LogIfError(err)
}

// ClearByPLID deletes all supervisors of a photo live
func (*PhotoLivesSupervisorsDataAccessObject) ClearByPLID(plID int) {
	var buf PhotoLivesSupervisors
	_, err := orm.Table(PhotoLivesSupervisorsDAO.TableName()).
		Where("photo_live_id=?", plID).Delete(&buf)
	logger.LogIfError(err)
}
