package models

type Photo_tags struct {
	ID int `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	Tag string `xorm:"tag VARCHAR(45) NOTNULL"`
	PhotoID int `xorm:"photo_id INT NOTNULL UNIQUE INDEX(photo_id_UNIQUE)"`
}
