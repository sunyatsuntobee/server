package models

// Activity Model
type Activity struct {
	ID             int    `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR"`
	Name           string `xorm:"name VARCHAR(45) NOTNULL"`
	Description    string `xorm:"description VARCHAR(100) NOTNULL"`
	Category       string `xorm:"category VARCHAR(45) NOTNULL"`
	PosterUrl      string `xorm:"poster_url VARCHAR(45)"`
	LogoURL        string `xorm:"logo_url VARCHAR(45)"`
	OrganizationID int    `xorm:"organization_id INT NOTNULL INDEX(organization_id_idx)"`
}

type ActivityDataAccessObject struct{}

var ActivityDAO *ActivityDataAccessObject

func (*ActivityDataAccessObject) TableName() string {
	return "activities"
}

func (*ActivityDataAccessObject) FindFullByOID(oid int) []ActivityFull {
	activities := make([]ActivityFull, 0)
	orm.Table(ActivityDAO.TableName()).
		Join("INNER", ActivityStageDAO.TableName(),
			"activities.id=activity_stages.activity_id").
		Join("INNER", OrganizationDAO.TableName(),
			"activities.organization_id=organizations.id").
		Asc("stage_num").
		Find(&activities)
	return activities
}
