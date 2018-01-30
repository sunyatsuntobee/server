package controllers

import (
	"html/template"
)

type layout struct {
	Title   string
	Content template.HTML
}

type contact struct {
	Name  string
	Phone string
}

type photoLive struct {
	ID                       int       `json:"id"`
	ActivityName             string    `json:"activity_name"`
	OrganizationName         string    `json:"organization_name"`
	Location                 string    `json:"location"`
	StartTime                string    `json:"start_time"`
	EndTime                  string    `json:"end_time"`
	ExpectMembers            int       `json:"expect_members"`
	AdProgress               string    `json:"ad_progress"`
	ManagerName              string    `json:"manager_name"`
	ManagerPhone             string    `json:"manager_phone"`
	PhotographerManagerName  string    `json:"photographer_manager_name"`
	PhotographerManagerPhone string    `json:"photographer_manager_phone"`
	Supervisors              []contact `json:"supervisors"`
}
