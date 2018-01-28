package models

type Photo_lives struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	ExpectMembers int `xorm:"expect_members INT NOTNULL"`
	AdProgress string `xorm:"ad_progress VARCHAR(45) NOTNULL"`
	ActivityStageID int `xorm:"activity_stage_id INT"`
	ManagerID int `xorm:"manager_id INT"`
	PhotographerManagerID int `xorm:"photographer_manager_id INT"`
}
