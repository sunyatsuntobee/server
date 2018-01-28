package models

type Photo_live_supervisor_relationships struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	PhotoLiveID int `xorm:"photo_live_id INT NOTNULL INDEX(photo_live_id_idx)"`
	PhotoSuperVisorRelationshipsCol int `xorm:"photo_supervisor_relationshipscol INT NOTNULL INDEX(supervisor_id_idx)"`
}