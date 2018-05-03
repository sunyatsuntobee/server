package models

import (
	"time"

	"github.com/sunyatsuntobee/server/logger"
)

// UsersParticipateOrganizations Model
type UsersParticipateOrganizations struct {
	ID             int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	UserID         int       `xorm:"user_id INT NOTNULL INDEX(contact_id_idx)" json:"user_id"`
	OrganizationID int       `xorm:"organization_id INT NOTNULL INDEX(organization_id_idx)" json:"organization_id"`
	Privilege	   int 		 `xorm:"privilege INT NOTNULL" json:"privilege"`
	Applying	   bool		 `xorm:"applying BOOL" json:"applying"`
	Timestamp      time.Time `xorm:"timestamp DATETIME NOTNULL" json:"timestamp"`
}

// UsersParticipateOrganizationsDataAccessObject provides database access for Model
// UsersParticipateOrganizations
type UsersParticipateOrganizationsDataAccessObject struct{}

// UsersParticipateOrganizationsDAO instance of
// UsersParticipateOrganizationsDataAccessObject
var UsersParticipateOrganizationsDAO *UsersParticipateOrganizationsDataAccessObject

// TableName returns table name
func (*UsersParticipateOrganizationsDataAccessObject) TableName() string {
	return "users_participate_organizations"
}

// NewUsersParticipateOrganizations creates a new relationship
func NewUsersParticipateOrganizations(userID int,
	organizationID int) *UsersParticipateOrganizations {
	return &UsersParticipateOrganizations{
		UserID:         userID,
		OrganizationID: organizationID,
	}
}

// DeleteByOID deletes all contactors according to an organization ID
func (*UsersParticipateOrganizationsDataAccessObject) DeleteByOID(oid int) {
	var buf UsersParticipateOrganizations
	_, err := orm.Table(UsersParticipateOrganizationsDAO.TableName()).
		Where("organization_id=?", oid).Unscoped().Delete(&buf)
	logger.LogIfError(err)
}

// InsertOne inserts a new relationship
func (*UsersParticipateOrganizationsDataAccessObject) InsertOne(
	r *UsersParticipateOrganizations) {
	_, err := orm.Table(UsersParticipateOrganizationsDAO.TableName()).
		InsertOne(r)
	logger.LogIfError(err)
}
