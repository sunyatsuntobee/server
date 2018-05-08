package models

import (
	"time"
	"github.com/sunyatsuntobee/server/logger"
)

// UsersFollowActivities Model
type UsersFollowActivities struct {
	ID         int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	UserID     int       `xorm:"user_id INT NOTNULL INDEX(user_id_idx)" json:"user_id"`
	ActivityID int       `xorm:"activity_id INT NOTNULL INDEX(activity_id_idx)" json:"activity_id"`
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

func (*UsersFollowActivitiesDataAccessObject) TableName() string {
	return "users_follow_activities"
}

type UsersFollowActivitiesDataAccessObject struct{}

var UsersFollowActivitiesDAO *UsersFollowActivitiesDataAccessObject


func (*UsersFollowActivitiesDataAccessObject) InsertOne(
	r *UsersFollowActivities) {
	_, err := orm.Table(UsersFollowActivitiesDAO.TableName()).
		InsertOne(r)
	logger.LogIfError(err)
}

//DeleteByID delete a user-activity relationship by its ID
func (*UsersFollowActivitiesDataAccessObject) DeleteByID(id int) {
	var usersFollowActivities UsersFollowActivities
	_, err := orm.Table(UsersFollowActivitiesDAO.TableName()).
		ID(id).Delete(&usersFollowActivities)
	logger.LogIfError(err)
}
