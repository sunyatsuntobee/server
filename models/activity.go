package models

// Activity Model
type Activity struct {
	ID             int     `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	Name           string  `xorm:"name VARCHAR(45) NOTNULL"`
	Description    string  `xorm:"description VARCHAR(100) NOTNULL"`
	Category       string  `xorm:"category VARCHAR(45) NOTNULL"`
	PosterUrl      string  `xorm:"poster_url VARCHAR(45)"`
	Logo           []uint8 `xorm:"logo BLOB"`
	OrganizationID int     `xorm:"organization_id INT NOTNULL INDEX(organization_id_idx)"`
}
