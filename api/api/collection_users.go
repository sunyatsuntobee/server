package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/models"
)

func initCollectionUsersRouter(router *mux.Router) {
	// GET /users
	router.HandleFunc("/api/users", usersGetHandler()).
		Methods(http.MethodGet)
}

func usersGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		w.Header().Set("Access-Control-Allow-Origin", cor)
		if req.FormValue("id") == "" {
			users := models.UserDAO.FindAll()
			formatter.JSON(w, http.StatusOK, users)
		}
	}

}
