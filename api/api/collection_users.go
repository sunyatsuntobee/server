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

	// GET /users_follow_users{?user_ID,followed_user_ID}
	router.HandleFunc("/api/users_follow_users",
		usersGetFollowHandler()).Methods(http.MethodGet)

	// POST /users_follow_users
	router.HandleFunc("/api/users_follow_users",
		usersFollowCreateHandler()).Methods(http.MethodPost)

	// DELETE /users_follow_users/{ID}
	router.HandleFunc("/api/users_follow_organizations/{ID}",
		usersFollowUsersDeleteHandler()).Methods(http.MethodDelete)

	// GET /users_follow_organizations{?user_ID,organization_ID}
	router.HandleFunc("/api/users_follow_organizations",
		organizationsGetFollowHandler()).Methods(http.MethodGet)

	// POST /users_follow_organizations
	router.HandleFunc("/api/users_follow_organizations",
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

	// POST /users_apply_organizations
	router.HandleFunc("/api/users_apply_organizations",
		usersApplyOrganizationsCreateHandler()).Methods(http.MethodPost)
	
	// POST /users_follow_activities
	router.HandleFunc("/api/users_follow_activities",
		usersFollowActivitiesCreateHandler()).Methods(http.MethodPost)

	// DELETE /users_follow_activities/ID
	router.HandleFunc("/api/users_follow_activities/{ID}",
		usersFollowActivitiesDeleteHandler()).Methods(http.MethodDelete)
}
func usersFollowUsersDeleteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		usersFollowUsersIDInt, _ := strconv.Atoi(mux.Vars(req)["ID"])
		models.UsersFollowUsersDAO.DeleteByID(usersFollowUsersIDInt)
		formatter.JSON(w, http.StatusNoContent, nil)
	}
}
func usersFollowOrganizationsDeleteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		usersFollowOrganizationsIDInt, _ := strconv.Atoi(mux.Vars(req)["ID"])
		models.UsersFollowUsersDAO.DeleteByID(usersFollowOrganizationsIDInt)
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

func usersApplyOrganizationsCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		decoder := json.NewDecoder(req.Body)
		var usersParticipateOrganizations models.UsersParticipateOrganizations
		err := decoder.Decode(&usersParticipateOrganizations)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", err))
			return
		}
		usersParticipateOrganizations.Timestamp = time.Now()
		usersParticipateOrganizations.Privilege = 1;
		usersParticipateOrganizations.Applying = true;
		models.UsersParticipateOrganizationsDAO.InsertOne(&usersParticipateOrganizations)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("Created", "加入社团成功", usersParticipateOrganizations))
	}
}

func usersFollowActivitiesCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		decoder := json.NewDecoder(req.Body)
		var usersFollowActivities models.UsersFollowActivities
		err := decoder.Decode(&usersFollowActivities)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}
		usersFollowActivities.Timestamp = time.Now()
		models.UsersFollowActivitiesDAO.InsertOne(&usersFollowActivities)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("Created", "关注活动成功", usersFollowActivities))
	}
}
func usersFollowActivitiesDeleteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		usersFollowActivitiesIDInt, _ := strconv.Atoi(mux.Vars(req)["ID"])
		models.UsersFollowActivitiesDAO.DeleteByID(usersFollowActivitiesIDInt)
		formatter.JSON(w, http.StatusNoContent, nil)
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
		userID := req.FormValue("user_ID")
		followedUserID := req.FormValue("followed_user_ID")
		userIDInt, _ := strconv.Atoi(userID)
		followedUserIDInt, _ := strconv.Atoi(followedUserID)
		if userID != "" && followedUserID != "" {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "输入参数错误", nil))
		}
		if userID != "" && followedUserID == "" {
			FollowedUserList :=
				models.UsersFollowUsersDAO.FindFullByUserID(userIDInt)
			formatter.JSON(w, http.StatusOK,
				NewJSON("OK", "获取该用户关注的用户列表成功", FollowedUserList))
		}
		if userID == "" && followedUserID != "" {
			FollowedUserList :=
				models.UsersFollowUsersDAO.FindFullByFollowedUserID(followedUserIDInt)
			formatter.JSON(w, http.StatusOK,
				NewJSON("OK", "获取关注该用户的用户列表成功", FollowedUserList))
		}
	}
}
func organizationsGetFollowHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		userID := req.FormValue("user_ID")
		followedOrganizationID := req.FormValue("organization_ID")
		userIDInt, _ := strconv.Atoi(userID)
		followedOrganizationIDInt, _ :=
			strconv.Atoi(followedOrganizationID)
		if userID != "" && followedOrganizationID != "" {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "输入参数错误", nil))
		}
		if userID != "" && followedOrganizationID == "" {
			OrganizaitionList :=
				models.UsersFollowOrganizationsDAO.FindFullByUserID(userIDInt)
			formatter.JSON(w, http.StatusOK,
				NewJSON("OK", "获取该用户关注的社团列表成功", OrganizaitionList))
		}
		if userID == "" && followedOrganizationID != "" {
			UserList :=
				models.UsersFollowOrganizationsDAO.FindFullByOrganizationID(
					followedOrganizationIDInt)
			formatter.JSON(w, http.StatusOK,
				NewJSON("OK", "获取关注该社团的用户列表成功", UserList))
		}
	}
}
func usersGetByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		IDInt, _ := strconv.Atoi(mux.Vars(req)["ID"])
		user, hasUser := models.UserDAO.FindByID(IDInt)
		if hasUser == false {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "用户对象不存在", nil))
		} else {
			formatter.JSON(w, http.StatusOK,
				NewJSON("OK", "获取对应ID的用户成功", user))
		}
	}
}

func usersUploadAvatarHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		ID, _ := strconv.Atoi(mux.Vars(req)["ID"])
		user, has := models.UserDAO.FindByID(ID)
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
		ID, _ := strconv.Atoi(mux.Vars(req)["ID"])
		user, has := models.UserDAO.FindByID(ID)
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
