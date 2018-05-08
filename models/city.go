package models

// City Model
type City struct {
	ID         int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	Name       string `xorm:"name VARCHAR(20) NOTNULL" json:"name"`
	ProvinceID int    `xorm:"province_id INT NOTNULL INDEX(fk_cities_province_id_idx)" json:"province_id"`
}
