package models

// OrganizationHonor Model
type OrganizationHonor struct {
	ID             int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	Honor          string `xorm:"honor VARCHAR(50) NOTNULL" json:"honor"`
	OrganizationID int    `xorm:"organization_id INT NOTNULL INDEX(fk_organization_honors_organization_id_idx)" json:"organization_id"`
}
