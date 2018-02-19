package models

import "github.com/sunyatsuntobee/server/logger"

// PhotoLiveSupervisorRelationship Model
type PhotoLiveSupervisorRelationship struct {
	ID           int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	PhotoLiveID  int `xorm:"photo_live_id INT NOTNULL INDEX(photo_live_id_idx)"`
	SupervisorID int `xorm:"supervisor_id INT NOTNULL INDEX(supervisor_id_idx)"`
}

type PhotoLiveSupervisorRDataAccessObject struct{}

var PhotoLiveSupervisorRDAO *PhotoLiveSupervisorRDataAccessObject

func NewTableName(photo_live_id int, supervisor_id int) {
	return &TableName{PhotoLiveID:photo_live_id, SupervisorID:supervisor_id}
}

func (*PhotoLiveSupervisorRDataAccessObject) TableName() string {
	return "photo_live_supervisor_relationships"
}

func (*PhotoLiveSupervisorRDataAccessObject) InsertOne(
	r *PhotoLiveSupervisorRelationship) {
	_, err := orm.Table(PhotoLiveSupervisorRDAO.TableName()).InsertOne(r)
	logger.LogIfError(err)
}

func (*PhotoLiveSupervisorRDataAccessObject) ClearByPLID(plID int) {
	var buf PhotoLiveSupervisorRelationship
	_, err := orm.Table(PhotoLiveSupervisorRDAO.TableName()).
		Where("photo_live_id=?", plID).
		Unscoped().Delete(&buf)
	logger.LogIfError(err)
}
