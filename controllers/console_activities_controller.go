package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func initConsoleActivitiesRouter(router *mux.Router) {
	router.HandleFunc("/activities", consoleActivitiesGetHandler()).
		Methods(http.MethodGet)
}

func consoleActivitiesGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		oid, _ := strconv.Atoi(req.FormValue("oid"))
		formatter.HTML(w, http.StatusOK, "console/activities/activities", nil)
	}

}
