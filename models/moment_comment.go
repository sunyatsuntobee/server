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

// MomentCommentDataAccessObject provides database access for
// Model MomentComment
type MomentCommentDataAccessObject struct{}

// MomentCommentDAO instance of MomentCommentDataAccessObject
var MomentCommentDAO *MomentCommentDataAccessObject

// TableName returns table name
func (*MomentCommentDataAccessObject) TableName() string {
	return "moment_comment"
}

func FindByMomentID(moment_id int) MomentComment {
	var moment_comment MomentComment
	err := orm.Table(MomentCommentDAO.TableName()).
		Where("moment_commet.moment_id=?", moment_id).
		Find(&moment_comment)
	logger.LogIfError(err)
	return moment_comment
}

// InsertOne inserts a moment_comment
func (*MomentCommentDataAccessObject) InsertOne(moment_comment *MomentComment) {
	_, err := orm.Table(MomentCommentDAO.TableName()).InsertOne(moment_comment)
	logger.LogIfError(err)
}

//DeleteByID delete a moment_comment by its ID
func (*MomentCommentDataAccessObject) DeleteByID(id int) {
	var moment_comment MomentComment
	_, err := orm.Table(MomentCommentDAO.TableName()).
		ID(id).Delete(&moment_comment)
	logger.LogIfError(err)
}
