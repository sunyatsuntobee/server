package models

import "time"

// UsersFocusOrganizations Model
type UsersFocusOrganizations struct {
	ID             int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	UserID         int       `xorm:"user_id NOTNULL INDEX(user_id_idx)" json:"user_id"`
	OrganizationID int       `xorm:"organization_id INT NOTNULL INDEX(organization_id_idx)" json:"organization_id"`
	Timestamp      time.Time `xorm:"timestamp DATETIME NOTNULL" json:"timestamp"`
}

// NewUsersOrganizations creates a new user-organization relationship
func NewUsersFocusOrganizations(userID int,
	organizationID int) *UsersFocusOrganizations {
	return &UsersFocusOrganizations{
		UserID:         userID,
		OrganizationID: organizationID,
		Timestamp:      time.Now(),
	}
}
