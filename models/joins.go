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
	Organization            `xorm:"extends"`
	OrganizationsContactors `xorm:"extends"`
	Contactor               User                   `xorm:"extends"`
	Department              OrganizationDepartment `xorm:"extends"`
	LoginLog                OrganizationLoginLog   `xorm:"extends"`
	Activity                `xorm:"extends"`
	Stage                   ActivityStage `xorm:"extends"`
}

// ActivityFull contains all information for an Entity Activity
type ActivityFull struct {
	Activity     `xorm:"extends" json:"activity"`
	Stage        ActivityStage `xorm:"extends" json:"activity_stage"`
	Organization `xorm:"extends" json:"organization"`
}

// ActivityStageFull contains all information for an Entity Activity Stage
type ActivityStageFull struct {
	ActivityStage `xorm:"extends" json:"activity_stage"`
	Activity      `xorm:"extends" json:"activity"`
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
