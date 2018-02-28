package models

// UsersLikeMoments Model
type UsersLikeMoments struct {
	ID       int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	UserID   int `xorm:"user_id INT NOTNULL INDEX(fk_users_like_moments_user_id_idx)" json:"user_id"`
	MomentID int `xorm:"moment_id INT NOTNULL INDEX(fk_users_like_moments_moment_id_idx)" json:"moment_id"`
}
