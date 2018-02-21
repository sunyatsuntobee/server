package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/models"
	"github.com/sunyatsuntobee/server/util"
)

func initCollectionUsersRouter(router *mux.Router) {
	// GET /users
	router.HandleFunc("/api/users", usersGetHandler()).
		Methods(http.MethodGet)

	// PUT /users/{ID}
	router.HandleFunc("/api/users/{ID}", usersPutHandler()).
		Methods(http.MethodPut)

	// POST /Create a  new user/
	router.HandleFunc("/api/users", usersCreatHandler()).
		Methods(http.MethodPost)
}
func usersCreatHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var flagPhone bool = false
		req.ParseForm()
		postUsername := req.FormValue("username")
		postPassword := util.MD5Hash(req.FormValue("password"))
		postPhone := req.FormValue("phone")
		postLocation := req.FormValue("location")
		postCreateTime := time.Now()
		postVipString := req.FormValue("vip")
		postAvatarUrl := req.FormValue("avatar_url")
		postCamera := req.FormValue("camera")
		postDescription := req.FormValue("description")
		postOccupation := req.FormValue("occupation")
		postCollege := req.FormValue("college")
		postVipBool, _ := strconv.ParseBool(postVipString)

		_, flagPhone =
			models.UserDAO.FindByPhone(postPhone)
		if flagPhone == true {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "此号码已被使用", nil))
			return
		}

		user := models.NewUser(postUsername,
			postPhone,
			postPassword,
			postLocation,
			postCreateTime,
			postVipBool,
			postAvatarUrl,
			postCamera,
			postDescription,
			postOccupation,
			postCollege,
		)

		models.UserDAO.InsertOne(&user)

		formatter.JSON(w, http.StatusCreated, NewJSON("Created", "注册成功", user))
	}
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
		user.College = req.FormValue("college")
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
