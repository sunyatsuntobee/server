package controllers

import "github.com/gorilla/mux"

func initConsoleRouter(router *mux.Router) {
	initConsolePhotoLivesRouter(router)
	initConsoleOrganizationsRouter(router)
	initConsoleActivitiesRouter(router)
}
