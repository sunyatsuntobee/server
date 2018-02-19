package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/models"
)

func initCollectionUsersRouter(router *mux.Router) {
	// GET /users
	router.HandleFunc("/api/users", usersGetHandler()).
		Methods(http.MethodGet)

	// PUT /users/{ID}
	router.HandleFunc("/api/users/{ID}", usersPutHandler()).
		Methods(http.MethodPut)
}

func usersPutHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		id, _ := strconv.Atoi(req.FormValue("id"))
		user, _ := models.UserDAO.FindByID(id)
		user.Username = req.FormValue("username")
		user.Location = req.FormValue("location")
		user.Camera = req.FormValue("camera")
		user.Description = req.FormValue("description")
		user.Occupation = req.FormValue("occupation")
		user.College = req.FormValue("collage")
		models.UserDAO.UpdateOne(&user)
		formatter.JSON(w, http.StatusCreated, nil)
	}

}

func usersGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if req.FormValue("id") == "" {
			users := models.UserDAO.FindAll()
			formatter.JSON(w, http.StatusOK, users)
		}
	}

}
