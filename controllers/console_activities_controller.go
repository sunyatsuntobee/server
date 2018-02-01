package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/models"
)

func initConsoleActivitiesRouter(router *mux.Router) {
	router.HandleFunc("/activities", consoleActivitiesGetHandler()).
		Methods(http.MethodGet)
}

func consoleActivitiesGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		oid, _ := strconv.Atoi(req.FormValue("oid"))
		data := models.ActivityDAO.FindFullByOID(oid)
		formatter.HTML(w, http.StatusOK, "console/activities/activities", data)
	}

}
