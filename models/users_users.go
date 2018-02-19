package models

// UserUserRelationship Model
type UsersUsers struct {
	ID          int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	UserID      int `xorm:"user_id INT NOTNULL INDEX(user_id_idx)"`
	LikedUserID int `xorm:"liked_user_id INT NOTNULL INDEX(liked_user_id_idx)"`
}

func NewUsersUsers(userID int, likedUserID int) *UsersUsers {
	return &UsersUsers{
		UserID:      userID,
		LikedUserID: likedUserID,
	}
}
