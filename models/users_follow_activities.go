package models

import "time"

// UsersFollowActivities Model
type UsersFollowActivities struct {
	ID         int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	UserID     int       `xorm:"user_id INT NOTNULL INDEX(user_id_idx)"`
	ActivityID int       `xorm:"activity_id INT NOTNULL INDEX(activity_id_idx)"`
	Timestamp  time.Time `xorm:"timestamp DATETIME NOTNULL" json:"timestamp"`
}

// NewUsersFollowActivities creates a new user-activity relationship
func NewUsersFollowActivities(userID int, activityID int) *UsersFollowActivities {
	return &UsersFollowActivities{
		UserID:     userID,
		ActivityID: activityID,
		Timestamp:  time.Now(),
	}
}
