package models

import "time"

// UsersFocusActivities Model
type UsersFocusActivities struct {
	ID         int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	UserID     int       `xorm:"user_id INT NOTNULL INDEX(user_id_idx)"`
	ActivityID int       `xorm:"activity_id INT NOTNULL INDEX(activity_id_idx)"`
	Timestamp  time.Time `xorm:"timestamp DATETIME NOTNULL" json:"timestamp"`
}

// NewUsersFocusActivities creates a new user-activity relationship
func NewUsersFocusActivities(userID int, activityID int) *UsersFocusActivities {
	return &UsersFocusActivities{
		UserID:     userID,
		ActivityID: activityID,
		Timestamp:  time.Now(),
	}
}
