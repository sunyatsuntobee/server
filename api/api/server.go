package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

const (
	cor string = "*"
)

func NewServer() *negroni.Negroni {
	n := negroni.Classic()
	router := mux.NewRouter()

	InitRouter(router)

	n.UseHandler(router)

	return n
}

func InitRouter(router *mux.Router) {
	initAuthorizationRouter(router)
	initCollectionUsersRouter(router)
	initCollectionActivitiesRouter(router)
	initCollectionActivityStagesRouter(router)
	initCollectionPhotoLivesRouter(router)
	initCollectionPhotosRouter(router)
	initCollectionOrganizationsRouter(router)
}

func optionsHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", cor)
		w.WriteHeader(http.StatusNoContent)
	}

}
