package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/logger"
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
	router.HandleFunc("/api/users", usersCreateHandler()).
		Methods(http.MethodPost)
}
func usersCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		decoder := json.NewDecoder(req.Body)
		var user models.User
		err := decoder.Decode(&user)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}
		user.Password = util.MD5Hash(user.Password)
		user.CreateTime = time.Now()
		user.VIP = false

		_, flagPhone := models.UserDAO.FindByPhone(user.Phone)
		if flagPhone == true {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "此号码已被使用", nil))
			return
		}

		models.UserDAO.InsertOne(&user)

		formatter.JSON(w, http.StatusCreated,
			NewJSON("Created", "注册成功", user))
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
