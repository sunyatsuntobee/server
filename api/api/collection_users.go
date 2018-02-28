package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/logger"
	"github.com/sunyatsuntobee/server/models"
	"github.com/sunyatsuntobee/server/util"
)

func initCollectionUsersRouter(router *mux.Router) {
	url := "/api/users"

	// GET /users
	router.HandleFunc(url,
		usersGetHandler()).Methods(http.MethodGet)

	// POST /users
	router.HandleFunc(url,
		usersCreateHandler()).Methods(http.MethodPost)

	// PATCH /users/{ID}/avatar
	router.HandleFunc(url+"/{ID}/avatar",
		usersUploadAvatarHandler()).Methods(http.MethodPatch)

	// PATCH /users/{ID}/background
	router.HandleFunc(url+"/{ID}/background",
		usersUploadBackgroundHandler()).Methods(http.MethodPatch)
}
func usersUploadAvatarHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		id, _ := strconv.Atoi(mux.Vars(req)["ID"])
		user, has := models.UserDAO.FindByID(id)
		if !has {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "用户对象不存在", nil))
			return
		}
		file, header, err := req.FormFile("avatar")
		logger.LogIfError(err)
		defer file.Close()
		name := strings.Split(header.Filename, ".")
		path := resDir + "/avatars/" + mux.Vars(req)["ID"] + "." + name[1]
		url := "/res/avatars/" + mux.Vars(req)["ID"] + "." + name[1]
		util.SaveMultipartFile(path, file)
		user.AvatarURL = url
		models.UserDAO.UpdateOne(&user)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("created", "用户头像上传成功", user))
	}
}
func usersUploadBackgroundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		id, _ := strconv.Atoi(mux.Vars(req)["ID"])
		user, has := models.UserDAO.FindByID(id)
		if !has {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "用户对象不存在", nil))
			return
		}
		file, header, err := req.FormFile("background")
		logger.LogIfError(err)
		defer file.Close()
		name := strings.Split(header.Filename, ".")
		path := resDir + "/backgrounds/" + mux.Vars(req)["ID"] + "." + name[1]
		url := "/res/backgrounds/" + mux.Vars(req)["ID"] + "." + name[1]
		util.SaveMultipartFile(path, file)
		user.BackgroundURL = url
		models.UserDAO.UpdateOne(&user)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("created", "用户背景图像上传成功", user))
	}
}
func usersGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		userList := models.UserDAO.FindAll()
		formatter.JSON(w, http.StatusOK,
			NewJSON("OK", "获取用户列表成功", userList))
	}

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
		user.VIP = 0
		user.AvatarURL = ""
		user.BackgroundURL = ""

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
