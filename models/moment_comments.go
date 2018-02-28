package models

import "time"

// MomentComment Model
type MomentComment struct {
	ID        int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	MomentID  int       `xorm:"moment_id INT NOTNULL INDEX(fk_moment_comments_moment_id_idx)" json:"moment_id"`
	UserID    int       `xorm:"user_id INT NOTNULL INDEX(fk_moment_comments_user_id_idx)" json:"user_id"`
	Content   string    `xorm:"content VARCHAR(200) NOTNULL" json:"content"`
	Timestamp time.Time `xorm:"timestamp DATETIME NOTNULL" json:"timestamp"`
}
