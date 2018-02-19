package models

// UserActivityRelationship Model
type UserActivityRelationship struct {
	ID         int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	UserID     int `xorm:"user_id INT NOTNULL INDEX(user_id_idx)"`
	ActivityID int `xorm:"activity_id INT NOTNULL INDEX(activity_id_idx)"`
}

func NewUserActivityRelationship(user_id int, activity_id int) {
	return &UserActivityRelationship{UserID:user_id, ActivityID:activity_id}
}
