package models

import "github.com/sunyatsuntobee/server/logger"

// PhotoLivesSupervisors Model
type PhotoLivesSupervisors struct {
	ID           int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	PhotoLiveID  int `xorm:"photo_live_id INT NOTNULL INDEX(photo_live_id_idx)"`
	SupervisorID int `xorm:"supervisor_id INT NOTNULL INDEX(supervisor_id_idx)"`
}

type PhotoLivesSupervisorsDataAccessObject struct{}

var PhotoLivesSupervisorsDAO *PhotoLivesSupervisorsDataAccessObject

func (*PhotoLivesSupervisorsDataAccessObject) TableName() string {
	return "photo_lives_supervisors"
}

func NewPhotoLivesSupervisors(photoLiveID int,
	supervisorID int) *PhotoLivesSupervisors {
	return &PhotoLivesSupervisors{
		PhotoLiveID:  photoLiveID,
		SupervisorID: supervisorID,
	}
}

func (*PhotoLivesSupervisorsDataAccessObject) InsertOne(
	r *PhotoLivesSupervisors) {
	_, err := orm.Table(PhotoLivesSupervisorsDAO.TableName()).InsertOne(r)
	logger.LogIfError(err)
}

func (*PhotoLivesSupervisorsDataAccessObject) ClearByPLID(plID int) {
	var buf PhotoLivesSupervisors
	_, err := orm.Table(PhotoLivesSupervisorsDAO.TableName()).
		Where("photo_live_id=?", plID).
		Unscoped().Delete(&buf)
	logger.LogIfError(err)
}
