package models

import (
	"time"

	"github.com/sunyatsuntobee/server/logger"
)

// User Model
type User struct {
	ID          int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	Username    string    `xorm:"username VARCHAR(45) NOTNULL"`
	Phone       string    `xorm:"phone PK VARCHAR(45) NOTNULL"`
	Password    string    `xorm:"password VARCHAR(45) NOTNULL"`
	Location    string    `xorm:"location VARCHAR(45) NOTNULL"`
	CreateTime  time.Time `xorm:"create_time TIMESTAMP NOTNULL CREATED"`
	VIP         bool      `xorm:"vip INT NOTNULL"`
	AvatarURL   string    `xorm:"avatar_url VARCHAR(45)"`
	Camera      string    `xorm:"camera VARCHAR(45)"`
	Description string    `xorm:"description VARCHAR(45)"`
	Occupation  string    `xorm:"occupation VARCHAR(45)"`
	Collage     string    `xorm:"collage VARCHAR(45)"`
}

type UserDataAccessObject struct{}

var UserDAO *UserDataAccessObject

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

func (*UserDataAccessObject) UpdateByID(id int, user *User) {
	_, err := orm.Table(UserDAO.TableName()).ID(id).Update(user)
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
