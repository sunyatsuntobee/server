package models

import "time"

// UsersSignActivities Relationship Model
type UsersSignActivities struct {
	ID         int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	UserID     int       `xorm:"user_id INT NOTNULL INDEX(fk_users_sign_activities_user_id_idx)" json:"user_id"`
	ActivityID int       `xorm:"activity_id INT NOTNULL INDEX(fk_users_sign_activities_activity_id_idx)" json:"activity_id"`
	Timestamp  time.Time `xorm:"timestamp DATETIME NOTNULL" json:"timestamp"`
}
