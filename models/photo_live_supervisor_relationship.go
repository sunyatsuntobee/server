package models

// PhotoLiveSupervisorRelationship Model
type PhotoLiveSupervisorRelationship struct {
	ID           int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	PhotoLiveID  int `xorm:"photo_live_id INT NOTNULL INDEX(photo_live_id_idx)"`
	SupervisorID int `xorm:"supervisor_id INT NOTNULL INDEX(supervisor_id_idx)"`
}
