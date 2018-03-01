package models

import "github.com/sunyatsuntobee/server/logger"

// UsersLikeMoments Model
type UsersLikeMoments struct {
	ID       int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	UserID   int `xorm:"user_id INT NOTNULL INDEX(fk_users_like_moments_user_id_idx)" json:"user_id"`
	MomentID int `xorm:"moment_id INT NOTNULL INDEX(fk_users_like_moments_moment_id_idx)" json:"moment_id"`
}

// UsersLikeMomentsFull Model
type UsersLikeMomentsFull struct {
	ID       int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	User     `xorm:"extends" json:"user"`
	MomentID int `xorm:"moment_id INT NOTNULL INDEX(fk_users_like_moments_moment_id_idx)" json:"moment_id"`
}

// UsersLikeMomentsDataAccessObject provides database access for
// Model UsersLikeMoments
type UsersLikeMomentsDataAccessObject struct{}

// UsersLikeMomentsDAO instance of UsersLikeMomentsDataAccessObject
var UsersLikeMomentsDAO *UsersLikeMomentsDataAccessObject

// TableName returns table name
func (*UsersLikeMomentsDataAccessObject) TableName() string {
	return "users_like_moments"
}

// FindFullByMomentID finds all full models by a moment ID
func (*UsersLikeMomentsDataAccessObject) FindFullByMomentID(
	id int) []UsersLikeMomentsFull {
	l := make([]UsersLikeMomentsFull, 0)
	err := orm.Table(UsersLikeMomentsDAO.TableName()).
		Join("INNER", UserDAO.TableName(),
			UsersLikeMomentsDAO.TableName()+".user_id="+
				UserDAO.TableName()+".id").
		Join("INNER", MomentDAO.TableName(),
			UsersLikeMomentsDAO.TableName()+".moment_id="+
				MomentDAO.TableName()+".id").
		Where("moment_id=?", id).
		Find(&l)
	logger.LogIfError(err)
	return l
}

// InsertOne insert a user_like_moment relationship
func (*UsersLikeMomentsDataAccessObject) InsertOne(
	usersLikeMoments *UsersLikeMoments) {
	_, err := orm.Table(UsersLikeMomentsDAO.TableName()).
		InsertOne(usersLikeMoments)
	logger.LogIfError(err)
}

//DeleteByID delete a user_like_moment relationship by its ID
func (*UsersLikeMomentsDataAccessObject) DeleteByID(id int) {
	var users_like_moments UsersLikeMoments
	_, err := orm.Table(UsersLikeMomentsDAO.TableName()).
		ID(id).Delete(&users_like_moments)
	logger.LogIfError(err)
}
