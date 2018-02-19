package models

import (
	"time"

	"github.com/sunyatsuntobee/server/logger"
)

// User Model
type User struct {
	ID          int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	Username    string    `xorm:"username VARCHAR(20) NOTNULL" json:"username"`
	Phone       string    `xorm:"phone VARCHAR(20) NOTNULL" json:"phone"`
	Password    string    `xorm:"password VARCHAR(50) NOTNULL" json:"password"`
	Location    string    `xorm:"location VARCHAR(50) NOTNULL" json:"location"`
	CreateTime  time.Time `xorm:"create_time TIMESTAMP NOTNULL CREATED" json:"create_time"`
	VIP         bool      `xorm:"vip INT NOTNULL" json:"vip"`
	AvatarURL   string    `xorm:"avatar_url VARCHAR(45)" json:"avatar_url"`
	Camera      string    `xorm:"camera VARCHAR(45)" json:"camera"`
	Description string    `xorm:"description VARCHAR(45)" json:"description"`
	Occupation  string    `xorm:"occupation VARCHAR(45)" json:"occupation"`
	College     string    `xorm:"college VARCHAR(45)" json:"college"`
}

type UserDataAccessObject struct{}

var UserDAO *UserDataAccessObject

func NewUser(username string, phone string, password string, location string,
	create_time time.Time, vip bool, avatar_url string, camera string,
	description string, occupation string, college string) {
	return &User{Username: username, Phone: phone, Password: password, Location: location,
		CreateTime: create_time, VIP: vip, AvatarURL: avatar_url, Camera: camera,
		Description: description, Occupation: occupation, College: college}
}

func (*UserDataAccessObject) TableName() string {
	return "users"
}

func (*UserDataAccessObject) FindAll() []User {
	l := make([]User, 0)
	err := orm.Table(UserDAO.TableName()).Find(&l)
	logger.LogIfError(err)
	return l
}

func (*UserDataAccessObject) InsertOne(user *User) {
	_, err := orm.Table(UserDAO.TableName()).InsertOne(user)
	logger.LogIfError(err)
}

func (*UserDataAccessObject) UpdateOne(user *User) {
	_, err := orm.Table(UserDAO.TableName()).ID(user.ID).Update(user)
	logger.LogIfError(err)
}

func (*UserDataAccessObject) FindByID(id int) (User, bool) {
	var user User
	has, err := orm.Table(UserDAO.TableName()).Where("id=?", id).Get(&user)
	logger.LogIfError(err)
	return user, has
}

func (*UserDataAccessObject) FindByPhone(phone string) (User, bool) {
	var user User
	has, err := orm.Table(UserDAO.TableName()).Where("phone=?", phone).Get(&user)
	logger.LogIfError(err)
	return user, has
}
