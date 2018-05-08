package models

import (
	"time"

	"github.com/sunyatsuntobee/server/logger"
)

// MomentComment Model
type MomentComment struct {
	ID        int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	MomentID  int       `xorm:"moment_id INT NOTNULL INDEX(fk_moment_comments_moment_id_idx)" json:"moment_id"`
	UserID    int       `xorm:"user_id INT NOTNULL INDEX(fk_moment_comments_user_id_idx)" json:"user_id"`
	Content   string    `xorm:"content VARCHAR(200) NOTNULL" json:"content"`
	Timestamp time.Time `xorm:"timestamp DATETIME NOTNULL" json:"timestamp"`
}

// MomentCommentFull joins with the user table
type MomentCommentFull struct {
	ID        int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	MomentID  int       `xorm:"moment_id INT NOTNULL INDEX(fk_moment_comments_moment_id_idx)" json:"moment_id"`
	User      User      `xorm:"extends" json:"user"`
	Content   string    `xorm:"content VARCHAR(200) NOTNULL" json:"content"`
	Timestamp time.Time `xorm:"timestamp DATETIME NOTNULL" json:"timestamp"`
}

// MomentCommentDataAccessObject provides database access for
// Model MomentComment
type MomentCommentDataAccessObject struct{}

// MomentCommentDAO instance of MomentCommentDataAccessObject
var MomentCommentDAO *MomentCommentDataAccessObject

// TableName returns table name
func (*MomentCommentDataAccessObject) TableName() string {
	return "moments_comments"
}

// FindFullByMomentID finds full comments by moment id
func (*MomentCommentDataAccessObject) FindFullByMomentID(
	momentID int) MomentComment {
	var momentComment MomentComment
	err := orm.Table(MomentCommentDAO.TableName()).
		Join("INNER", UserDAO.TableName(),
			MomentCommentDAO.TableName()+".user_id="+
				UserDAO.TableName()+".id").
		Where("moment_id=?", momentID).
		Find(&momentComment)
	logger.LogIfError(err)
	return momentComment
}

// InsertOne inserts a momentComment
func (*MomentCommentDataAccessObject) InsertOne(momentComment *MomentComment) {
	_, err := orm.Table(MomentCommentDAO.TableName()).InsertOne(momentComment)
	logger.LogIfError(err)
}

//DeleteByID delete a momentComment by its ID
func (*MomentCommentDataAccessObject) DeleteByID(id int) {
	var momentComment MomentComment
	_, err := orm.Table(MomentCommentDAO.TableName()).
		ID(id).Delete(&momentComment)
	logger.LogIfError(err)
}
