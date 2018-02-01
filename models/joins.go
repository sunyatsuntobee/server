package models

// PhotoLiveFull contains all information for an Entity PhotoLive
type PhotoLiveFull struct {
	PhotoLive           PhotoLive     `xorm:"extends"`
	Organization        Organization  `xorm:"extends"`
	Activity            Activity      `xorm:"extends"`
	ActivityStage       ActivityStage `xorm:"extends"`
	Manager             User          `xorm:"extends"`
	PhotographerManager User          `xorm:"extends"`
	Supervisor          User          `xorm:"extends"`
}

// OrganizationFull contains all information for an Entity Organization
type OrganizationFull struct {
	Organization Organization           `xorm:"extends"`
	Contactor    User                   `xorm:"extends"`
	Department   OrganizationDepartment `xorm:"extends"`
	LoginLog     OrganizationLoginLog   `xorm:"extends"`
	Activity     Activity               `xorm:"extends"`
	Stage        ActivityStage          `xorm:"extends"`
}

// ActivityFull contains all information for an Entity Activity
type ActivityFull struct {
	Activity Activity      `xorm:"extends"`
	Stage    ActivityStage `xorm:"extends"`
}
