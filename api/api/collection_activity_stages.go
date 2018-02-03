package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/models"
)

func initCollectionActivityStagesRouter(router *mux.Router) {
	// GET /activitystages
	router.HandleFunc("/api/activitystages", activityStagesGetHandler()).
		Methods(http.MethodGet)
}

func activityStagesGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		w.Header().Set("Access-Control-Allow-Origin", cor)
		if req.FormValue("aid") == "" {

		} else {
			aid, _ := strconv.Atoi(req.FormValue("aid"))
			data := models.ActivityStageDAO.FindByAID(aid)
			formatter.JSON(w, http.StatusOK, data)
		}
	}

}
