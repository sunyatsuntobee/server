package models

import (
	"fmt"
	"reflect"
)

// PhotoLiveFull contains all information for an Entity PhotoLive
type PhotoLiveFull struct {
	PhotoLive             `xorm:"extends"`
	ActivityStage         `xorm:"extends"`
	Activity              `xorm:"extends"`
	Organization          `xorm:"extends"`
	Manager               User `xorm:"extends"`
	PhotographerManager   User `xorm:"extends"`
	PhotoLivesSupervisors `xorm:"extends"`
	Supervisor            User `xorm:"extends"`
}

// OrganizationFull contains all information for an Entity Organization
type OrganizationFull struct {
	Organization                  `xorm:"extends"`
	UsersParticipateOrganizations `xorm:"extends"`
	Contactor                     User                   `xorm:"extends"`
	Department                    OrganizationDepartment `xorm:"extends"`
	Activity                      `xorm:"extends"`
	Stage                         ActivityStage `xorm:"extends"`
}

// ActivityFull contains all information for an Entity Activity
type ActivityFull struct {
	Activity      `xorm:"extends" json:"activity"`
	ActivityStage `xorm:"extends" json:"activity_stage"`
	Organizations []Organization  `xorm:"extends" json:"organizations"`
}

// ActivityStageFull contains all information for an Entity Activity Stage
type ActivityStageFull struct {	
	Activity      `xorm:"extends" json:"activity"`
	ActivityStage `xorm:"extends"`
}

type ActivityAndStage struct {	
	Activity      `xorm:"extends" json:"activity"`
	//stages []ActivityStage `xorm:"extends"`
	Stages []ActivityStage `xorm:"extends" json:"activity_stages"`
}

// PhotoFull contains all information for an Entity Photo
type PhotoFull struct {
	Photo `xorm:"extends" json:"photo"`
	User  `xorm:"extends" json:"user"`
}

// UniqueAt deletes duplicated rows according to a column
func UniqueAt(col string, from interface{}) interface{} {
	fromVal := reflect.ValueOf(from)
	toVal := reflect.MakeSlice(reflect.TypeOf(from), 0, 0)
	for i := 0; i < fromVal.Len(); i++ {
		ele := fromVal.Index(i)
		contain := false
		for j := 0; j < toVal.Len(); j++ {
			if fmt.Sprint(toVal.Index(j).FieldByName(col)) ==
				fmt.Sprint(ele.FieldByName(col)) {
				contain = true
				break
			}
		}
		if !contain {
			toVal = reflect.Append(toVal, ele)
		}
	}
	return toVal.Interface()
}
