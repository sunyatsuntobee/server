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
	Camera      string    `xorm:"camera VARCHAR(45)"`
	Description string    `xorm:"description VARCHAR(45)"`
	Occupation  string    `xorm:"occupation VARCHAR(45)"`
	Collage     string    `xorm:"collage VARCHAR(45)"`
}

type UserDataAccessObject struct{}

const UserTableName string = "users"

var UserDAO *UserDataAccessObject

func (*UserDataAccessObject) FindAll() []User {
	l := make([]User, 0)
	err := orm.Table(UserTableName).Find(&l)
	logger.LogIfError(err)
	return l
}

func (*UserDataAccessObject) InsertOne(user *User) {
	_, err := orm.Table(UserTableName).InsertOne(user)
	logger.LogIfError(err)
}

func (*UserDataAccessObject) UpdateByID(id int, user *User) {
	_, err := orm.Table(UserTableName).ID(id).Update(user)
	logger.LogIfError(err)
}

func (*UserDataAccessObject) FindByPhone(phone string) (User, bool) {
	var user User
	has, err := orm.Table(UserTableName).Where("phone=?", phone).Get(&user)
	logger.LogIfError(err)
	return user, has
}
