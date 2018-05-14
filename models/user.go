package models

import (
	"time"

	"github.com/sunyatsuntobee/server/logger"
)

// User Model
type User struct {
	ID              int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	OpenID          string    `xorm:"openid VARCHAR(50) NOTNULL" json:"openid"`
	Phone           string    `xorm:"phone VARCHAR(20) NOTNULL" json:"phone"`
	Password        string    `xorm:"password VARCHAR(50) NOTNULL" json:"password"`
	Username        string    `xorm:"username VARCHAR(20) NOTNULL" json:"username"`
	Nickname        string    `xorm:"nickname VARCHAR(20) NOTNULL" json:"nickname"`
	Email           string    `xorm:"email VARCHAR(50) NOTNULL" json:"email"`
	AvatarURL       string    `xorm:"avatar_url VARCHAR(45)" json:"avatar_url"`
	CreateTime      time.Time `xorm:"create_time DATETIME NOTNULL" json:"create_time"`
	CityID          int       `xorm:"city_id int NOTNULL INDEX(fk_users_city_id_idx)" json:"city_id"`
	VIP             int       `xorm:"vip INT NOTNULL" json:"vip"`
	Camera          string    `xorm:"camera VARCHAR(50)" json:"camera"`
	Description     string    `xorm:"description VARCHAR(200)" json:"description"`
	College         string    `xorm:"college VARCHAR(50) NOTNULL" json:"college"`
	CollegeDistrict string    `xorm:"college_district VARCHAR(20) NOTNULL" json:"college_district"`
	EnrollTime      int       `xorm:"enroll_time INT" json:"enroll_time"`
	Institute       string    `xorm:"institute VARCHAR(50)" json:"institute"`
	Astrology       string    `xorm:"astrology VARCHAR(10)" json:"astrology"`
	QQ              string    `xorm:"qq VARCHAR(10)" json:"qq"`
	BackgroundURL   string    `xorm:"background_url VARCHAR(50)" json:"background_url"`
}

// NewUser creates a new user
func NewUser(phone, password, username, nickname, email string,
	cityID, vip int, camera, description, college string,
	enrollTime int, institute, astrology, qq string) *User {
	return &User{
		Phone:       phone,
		Password:    password,
		Username:    username,
		Nickname:    nickname,
		Email:       email,
		CreateTime:  time.Now(),
		CityID:      cityID,
		VIP:         vip,
		Camera:      camera,
		Description: description,
		College:     college,
		EnrollTime:  enrollTime,
		Institute:   institute,
		Astrology:   astrology,
		QQ:          qq,
	}
}

// UserDataAccessObject provides database access for Model User
type UserDataAccessObject struct{}

// UserDAO instance of UserDataAccessObject
var UserDAO *UserDataAccessObject

// TableName returns table name
func (*UserDataAccessObject) TableName() string {
	return "users"
}

// FindAll finds all users
func (*UserDataAccessObject) FindAll() []User {
	l := make([]User, 0)
	err := orm.Table(UserDAO.TableName()).Find(&l)
	logger.LogIfError(err)
	return l
}

// InsertOne inserts a user
func (*UserDataAccessObject) InsertOne(user *User) {
	_, err := orm.Table(UserDAO.TableName()).InsertOne(user)
	logger.LogIfError(err)
}

// UpdateOne updates a user
func (*UserDataAccessObject) UpdateOne(user *User) {
	_, err := orm.Table(UserDAO.TableName()).ID(user.ID).Update(user)
	logger.LogIfError(err)
}

// FindByID finds a user by ID
func (*UserDataAccessObject) FindByID(id int) (User, bool) {
	var user User
	has, err := orm.Table(UserDAO.TableName()).Where("id=?", id).Get(&user)
	logger.LogIfError(err)
	return user, has
}

// FindByPhone finds a user by phone number
func (*UserDataAccessObject) FindByPhone(phone string) (User, bool) {
	var user User
	has, err := orm.Table(UserDAO.TableName()).Where("phone=?", phone).Get(&user)
	logger.LogIfError(err)
	return user, has
}

func (*UserDataAccessObject) FindByOpenid(
	openid string) (User, bool) {
	var user User
	has, err := orm.Table(UserDAO.TableName()).
		Where("openid=?", openid).Get(&user)
	logger.LogIfError(err)
	return user, has
}
