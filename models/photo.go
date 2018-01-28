package models

import "time"

// Photo Model
type Photo struct {
	ID             int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	URL            string    `xorm:"url VARCHAR(45) NOTNULL"`
	TookTime       time.Time `xorm:"took_time DATETIME NOTNULL"`
	TookLocation   string    `xorm:"took_location VARCHAR(45) NOTNULL"`
	ReleaseTime    time.Time `xorm:"release_time DATETIME"`
	Category       string    `xorm:"category VARCHAR(45) NOTNULL"`
	Likes          int       `xorm:"likes INT NOTNULL"`
	PhotographerID int       `xorm:"photographer_id INT NOTNULL INDEX(photographer_id_idx)"`
}
