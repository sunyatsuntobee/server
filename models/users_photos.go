package models

// UsersPhotos Model
type UsersPhotos struct {
	ID           int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	UserID       int `xorm:"user_id INT NOTNULL INDEX(user_id_idx)"`
	LikedPhotoID int `xorm:"liked_photo_id INT NOTNULL INDEX(liked_photo_id_idx)"`
}

// NewUsersPhotos creates a new user-photo relationship
func NewUsersPhotos(userID int, likedPhotoID int) *UsersPhotos {
	return &UsersPhotos{
		UserID:       userID,
		LikedPhotoID: likedPhotoID,
	}
}
