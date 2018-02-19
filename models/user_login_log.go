package models

import "time"

// UserLoginLog Model
type UserLoginLog struct {
	ID            int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	LoginTime     time.Time `xorm:"login_time DATETIME NOTNULL"`
	LoginLocation string    `xorm:"login_location VARCHAR(50) NOTNULL"`
	LoginDevice   string    `xorm:"login_device VARCHAR(20) NOTNULL"`
	UserID        int       `xorm:"user_id INT NOTNULL INDEX(user_id_idx)"`
}
