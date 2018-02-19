package models

// UsersActivities Model
type UsersActivities struct {
	ID         int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	UserID     int `xorm:"user_id INT NOTNULL INDEX(user_id_idx)"`
	ActivityID int `xorm:"activity_id INT NOTNULL INDEX(activity_id_idx)"`
}

// NewUsersActivities creates a new user-activity relationship
func NewUsersActivities(userID int, activityID int) *UsersActivities {
	return &UsersActivities{
		UserID:     userID,
		ActivityID: activityID,
	}
}
