package controllers

import (
	"html/template"

	"github.com/sunyatsuntobee/server/models"
)

type layout struct {
	Title   string
	Content template.HTML
}

type contact struct {
	Name  string
	Phone string
}

type PhotoLiveDetail struct {
	PhotoLive           *models.PhotoLive
	Organization        *models.Organization
	Activity            *models.Activity
	ActivityStage       *models.ActivityStage
	Manager             *models.User
	PhotographerManager *models.User
	Supervisors         []*models.User
}

type OrganizationDetail struct {
	Organization *models.Organization
	Contactors   []*models.User
	Departments  []*models.OrganizationDepartment
	LoginLogs    []*models.OrganizationLoginLog
	Activities   []*ActivityDetail
}

type ActivityDetail struct {
	Activity *models.Activity
	Stages   []*models.ActivityStage
}
