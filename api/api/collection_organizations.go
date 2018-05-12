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

func initCollectionOrganizationsRouter(router *mux.Router) {
	url := "/api/organizations"

	// PUT /organizations/{ID}
	router.HandleFunc(url+"/{ID}",
		organizationsPutHandler()).Methods(http.MethodPut)

	// POST /organizations
	router.HandleFunc(url,
		organizationsCreateHandler()).Methods(http.MethodPost)

	// GET /organizations
	router.HandleFunc(url,
		organizationsGetHandler()).Methods(http.MethodGet)

	// POST /handle_applicants_and_members
	router.HandleFunc("/api/handle_applicants_and_members",
		organizationsApplyandMembersManageHandler()).Methods(http.MethodPost)
}

func organizationsGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		organizations := models.OrganizationDAO.FindAll()
		formatter.JSON(w, http.StatusOK,
			NewJSON("OK", "获取社团列表成功", organizations))
	}

}

func organizationsCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		decoder := json.NewDecoder(req.Body)
		var organization models.Organization
		err := decoder.Decode(&organization)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}
		organization.LogoURL = ""
		organization.Password = util.MD5Hash(organization.Password)
		_, flagPhone := models.OrganizationDAO.FindByPhone(organization.Phone)
		if flagPhone == true {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "此号码已被使用", nil))
			return
		}

		models.OrganizationDAO.InsertOne(&organization)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("Created", "注册成功", organization))
	}
}

func organizationsPutHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		var data models.Organization
		id, _ := strconv.Atoi(mux.Vars(req)["ID"])
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&data)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}
		old, has := models.OrganizationDAO.FindByID(id)
		if !has {
			formatter.JSON(w, http.StatusBadRequest, util.Error{
				Msg: "entity not found",
			})
		}
		old.Name = data.Name
		old.Phone = data.Phone
		old.College = data.College
		old.Description = data.Description
		models.OrganizationDAO.UpdateOne(&old)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("created", "修改社团信息成功", old))
	}

}

func organizationsApplyandMembersManageHandler() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		decoder := json.NewDecoder(r.Body)
		var usersParticipateOrganizations models.UsersParticipateOrganizations
		err := decoder.Decode(&usersParticipateOrganizations)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", err))
			return
		}
		usersParticipateOrganizations.Timestamp = time.Now()
		usersParticipateOrganizations.Applying = false;
		models.UsersParticipateOrganizationsDAO.InsertOne(&usersParticipateOrganizations)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("Created", "管理社团成员成功", usersParticipateOrganizations))
	}
}