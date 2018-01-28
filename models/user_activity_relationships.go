package models

type User_activity_relationships struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	UserID int `xorm:"user_id INT NOTNULL INDEX(user_id_idx)"`
	ActivityID int `xorm:"activity_id INT NOTNULL(activity_id_idx)"`
}