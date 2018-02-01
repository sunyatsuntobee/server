package controllers

import "github.com/gorilla/mux"

func initConsoleRouter(router *mux.Router) {
	initConsolePhotosRouter(router)
	initConsolePhotoLivesRouter(router)
	initConsoleOrganizationsRouter(router)
}
