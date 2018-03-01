package models

import "time"

// UsersFollowOrganizations Model
type UsersFollowOrganizations struct {
	ID             int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	UserID         int       `xorm:"user_id NOTNULL INDEX(user_id_idx)" json:"user_id"`
	OrganizationID int       `xorm:"organization_id INT NOTNULL INDEX(organization_id_idx)" json:"organization_id"`
	Timestamp      time.Time `xorm:"timestamp DATETIME NOTNULL" json:"timestamp"`
}

// NewUsersFollowOrganizations creates a new user-organization relationship
func NewUsersFollowOrganizations(userID int,
	organizationID int) *UsersFollowOrganizations {
	return &UsersFollowOrganizations{
		UserID:         userID,
		OrganizationID: organizationID,
		Timestamp:      time.Now(),
	}
}
