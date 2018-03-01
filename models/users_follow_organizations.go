package models

import (
	"time"

	"github.com/sunyatsuntobee/server/logger"
)

// UsersFollowOrganizations Model
type UsersFollowOrganizations struct {
	ID             int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	UserID         int       `xorm:"user_id NOTNULL INDEX(user_id_idx)" json:"user_id"`
	OrganizationID int       `xorm:"organization_id INT NOTNULL INDEX(organization_id_idx)" json:"organization_id"`
	Timestamp      time.Time `xorm:"timestamp DATETIME NOTNULL" json:"timestamp"`
}

// UsersFollowOrganizationsFull Model
type UsersFollowOrganizationsFull struct {
	ID                   int          `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	User                 User         `xorm:"extends" json:"user"`
	FollowedOrganization Organization `xorm:"extends" json:"followed_organization"`
	Timestamp            time.Time    `xorm:"timestamp DATETIME NOTNULL" json:"timestamp"`
}

// UsersFollowOrganizationsDataAccessObject provides database access for
// Model UsersFollowOrganizations
type UsersFollowOrganizationsDataAccessObject struct{}

// UsersFollowOrganizationsDAO instance of
// UsersFollowOrganizationsDataAccessObject
var UsersFollowOrganizationsDAO *UsersFollowOrganizationsDataAccessObject

// NewUsersFollowOrganizations creates a new user-organization relationship
func NewUsersFollowOrganizations(userID int,
	organizationID int) *UsersFollowOrganizations {
	return &UsersFollowOrganizations{
		UserID:         userID,
		OrganizationID: organizationID,
		Timestamp:      time.Now(),
	}
}

// TableName returns table name
func (*UsersFollowOrganizationsDataAccessObject) TableName() string {
	return "users_follow_organizations"
}

// FindFullByUserID finds all full models by a user ID
func (*UsersFollowOrganizationsDataAccessObject) FindFullByUserID(
	id int) []UsersFollowOrganizationsFull {
	l := make([]UsersFollowOrganizationsFull, 0)
	err := orm.Table(UsersFollowOrganizationsDAO.TableName()).
		Join("INNER", UserDAO.TableName(),
			UsersFollowOrganizationsDAO.TableName()+".user_id="+
				UserDAO.TableName()+".id").
		Join("INNER", OrganizationDAO.TableName(),
			UsersFollowOrganizationsDAO.TableName()+".organization_id="+
				OrganizationDAO.TableName()+".id").
		Where("user_id=?", id).
		Find(&l)
	logger.LogIfError(err)
	return l
}

// FindFullByOrganizationID finds a user_follow_organization relationship by organization id
func (*UsersFollowOrganizationsDataAccessObject) FindFullByOrganizationID(
	followedID int) []UsersFollowOrganizationsFull {
	l := make([]UsersFollowOrganizationsFull, 0)
	err := orm.Table(UsersFollowOrganizationsDAO.TableName()).
		Join("INNER", UserDAO.TableName(),
			UsersFollowOrganizationsDAO.TableName()+".user_id="+
				UserDAO.TableName()+".id").
		Join("INNER", OrganizationDAO.TableName(),
			UsersFollowOrganizationsDAO.TableName()+".organization_id="+
				OrganizationDAO.TableName()+".id").
		Where("organization_id=?", followedID).
		Find(&l)
	logger.LogIfError(err)
	return l
}

// InsertOne insert a user-organization relationship
func (*UsersFollowOrganizationsDataAccessObject) InsertOne(
	usersFollowOrganizations *UsersFollowOrganizations) {
	_, err := orm.Table(UsersFollowOrganizationsDAO.TableName()).
		InsertOne(usersFollowOrganizations)
	logger.LogIfError(err)
}

//DeleteByID delete a user-organization relationship by its ID
func (*UsersFollowOrganizationsDataAccessObject) DeleteByID(id int) {
	var usersFollowOrganizations UsersFollowOrganizations
	_, err := orm.Table(UsersFollowOrganizationsDAO.TableName()).
		ID(id).Delete(&usersFollowOrganizations)
	logger.LogIfError(err)
}
