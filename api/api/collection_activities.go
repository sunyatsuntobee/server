package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/models"
)

func initCollectionActivitiesRouter(router *mux.Router) {
	// GET /activities
	router.HandleFunc("/api/activities", activitiesGetHandler()).
		Methods(http.MethodGet)
}

func activitiesGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		w.Header().Set("Access-Control-Allow-Origin", cor)
		if req.FormValue("oid") == "" {

		} else {
			oid, _ := strconv.Atoi(req.FormValue("oid"))
			data := models.ActivityDAO.FindByOID(oid)
			formatter.JSON(w, http.StatusOK, data)
		}

	}

}
