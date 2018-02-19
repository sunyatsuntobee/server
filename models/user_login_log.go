package models

import "time"

// UserLoginLog Model
type UserLoginLog struct {
	ID            int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	LoginTime     time.Time `xorm:"login_time DATETIME NOTNULL"`
	LoginLocation string    `xorm:"login_location VARCHAR(45) NOTNULL"`
	LoginDevice   string    `xorm:"login_device VARCHAR(45) NOTNULL"`
	UserID        int       `xorm:"user_id INT NOTNULL INDEX(user_id_idx)"`
}

func NewUserLoginLog(login_time time.Time, login_location string,
					 login_device string, user_id int) {
	return &UserLoginLog{LoginTime:login_time, LoginLocation:login_location,
						 LoginDevice:login_device, UserID:user_id}
}