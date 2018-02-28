package models

import "time"

// UsersFocusUsers Model
type UsersFocusUsers struct {
	ID            int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	UserID        int       `xorm:"user_id INT NOTNULL INDEX(user_id_idx)"`
	FocusedUserID int       `xorm:"focused_user_id INT NOTNULL INDEX(liked_user_id_idx)"`
	Timestamp     time.Time `xorm:"timestamp DATETIME NOTNULL" json:"timestamp"`
}

// NewUsersUsers creates a new user-user relationship
func NewUsersFocusUsers(userID int, focusedUserID int) *UsersFocusUsers {
	return &UsersFocusUsers{
		UserID:        userID,
		FocusedUserID: focusedUserID,
		Timestamp:     time.Now(),
	}
}
