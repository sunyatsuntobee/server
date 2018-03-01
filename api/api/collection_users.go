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

	// GET /users/{ID}
	router.HandleFunc(url+"/{ID}",
		usersGetByIDHandler()).Methods(http.MethodGet)

	// GET /users_follow_users{?user_id,followed_user_id}
	router.HandleFunc("/api/users_follow_users",
		usersGetFollowHandler()).Methods(http.MethodGet)

	// POST /users_follow_users
	router.HandleFunc("/api/users_follow_users",
		usersFollowCreateHandler()).Methods(http.MethodPost)

	// DELETE /users_follow_users/{ID}
	router.HandleFunc("/api/users_follow_organizations/{ID}",
		usersFollowUsersDeleteHandler()).Methods(http.MethodDelete)

	// GET /users_follow_organizations{?user_id,organization_id}
	router.HandleFunc("/api/users_follow_organizations",
		organizationsGetFollowHandler()).Methods(http.MethodGet)

	// POST /users_follow_organizations
	router.HandleFunc("/api/users_follow_users",
		usersFollowOrganizationsCreateHandler()).Methods(http.MethodPost)

	// DELETE /users_follow_organizations/{ID}
	router.HandleFunc("/api/users_follow_organizations/{ID}",
		usersFollowOrganizationsDeleteHandler()).Methods(http.MethodDelete)

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
func usersFollowUsersDeleteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		usersFollowUsersIdInt, _ := strconv.Atoi(mux.Vars(req)["ID"])
		models.UsersFollowUsersDAO.DeleteByID(usersFollowUsersIdInt)
		formatter.JSON(w, http.StatusNoContent, nil)
	}
}
func usersFollowOrganizationsDeleteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		usersFollowOrganizationsIdInt, _ := strconv.Atoi(mux.Vars(req)["ID"])
		models.UsersFollowUsersDAO.DeleteByID(usersFollowOrganizationsIdInt)
		formatter.JSON(w, http.StatusNoContent, nil)
	}
}
func usersFollowOrganizationsCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		decoder := json.NewDecoder(req.Body)
		var usersFollowOrganizations models.UsersFollowOrganizations
		err := decoder.Decode(&usersFollowOrganizations)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}
		usersFollowOrganizations.Timestamp = time.Now()
		models.UsersFollowOrganizationsDAO.InsertOne(&usersFollowOrganizations)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("Created", "关注社团成功", usersFollowOrganizations))
	}
}

func usersFollowCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		decoder := json.NewDecoder(req.Body)
		var usersFollowUsers models.UsersFollowUsers
		err := decoder.Decode(&usersFollowUsers)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}
		usersFollowUsers.Timestamp = time.Now()
		models.UsersFollowUsersDAO.InsertOne(&usersFollowUsers)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("Created", "关注用户成功", usersFollowUsers))
	}
}
func usersGetFollowHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		userId := req.FormValue("user_id")
		followedUserId := req.FormValue("followed_user_id")
		userIdInt, _ := strconv.Atoi(userId)
		followedUserIdInt, _ := strconv.Atoi(followedUserId)
		if userId != "" && followedUserId != "" {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "输入参数错误", nil))
		}
		if userId != "" && followedUserId == "" {
			FollowedUserList :=
				models.UsersFollowUsersDAO.FindFullByUserID(userIdInt)
			formatter.JSON(w, http.StatusOK,
				NewJSON("OK", "获取该用户关注的用户列表成功", FollowedUserList))
		}
		if userId == "" && followedUserId != "" {
			FollowedUserList :=
				models.UsersFollowUsersDAO.FindFullByFollowedUserID(followedUserIdInt)
			formatter.JSON(w, http.StatusOK,
				NewJSON("OK", "获取关注该用户的用户列表成功", FollowedUserList))
		}
	}
}
func organizationsGetFollowHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		userId := req.FormValue("user_id")
		followedOrganizationId := req.FormValue("organization_id")
		userIdInt, _ := strconv.Atoi(userId)
		followedOrganizationIdInt, _ :=
			strconv.Atoi(followedOrganizationId)
		if userId != "" && followedOrganizationId != "" {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "输入参数错误", nil))
		}
		if userId != "" && followedOrganizationId == "" {
			OrganizaitionList :=
				models.UsersFollowOrganizationsDAO.FindFullByUserID(userIdInt)
			formatter.JSON(w, http.StatusOK,
				NewJSON("OK", "获取该用户关注的社团列表成功", OrganizaitionList))
		}
		if userId == "" && followedOrganizationId != "" {
			UserList :=
				models.UsersFollowOrganizationsDAO.FindFullByOrganizationID(
					followedOrganizationIdInt)
			formatter.JSON(w, http.StatusOK,
				NewJSON("OK", "获取关注该社团的用户列表成功", UserList))
		}
	}
}
func usersGetByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		idInt, _ := strconv.Atoi(mux.Vars(req)["ID"])
		user, hasUser := models.UserDAO.FindByID(idInt)
		if hasUser == false {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "用户对象不存在", nil))
			return
		} else {
			formatter.JSON(w, http.StatusOK,
				NewJSON("OK", "获取对应id的用户成功", user))
		}
	}
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
