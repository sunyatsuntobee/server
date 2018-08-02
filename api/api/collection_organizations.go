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
	//添加认证
	router.HandleFunc(url+"/{ID}",
	    handlerSecure(organizationsPutHandler())).Methods(http.MethodPut)

	// POST /organizations
	router.HandleFunc(url,
		organizationsCreateHandler()).Methods(http.MethodPost)

	// GET /organizations
	router.HandleFunc(url,
		organizationsGetHandler()).Methods(http.MethodGet)

	// GET /organizations{?aid} //由于数据库更改，已经添加了根据活动ID获取社团的api
	router.HandleFunc(url,
		organizationsGetHandler()).Methods(http.MethodGet)
	
	// GET /spread_organizations
	router.HandleFunc(url+"/spread_organizations",
	    organizationsGetSpreadHandler()).Methods(http.MethodGet)

	// POST /handle_applicants_and_members
	//添加认证
	router.HandleFunc("/api/handle_applicants_and_members",
		handlerSecure(organizationsApplyandMembersManageHandler())).Methods(http.MethodPost)

	// PATCH /organizations/{ID}/coin
	router.HandleFunc(url+"/{ID}/coin",
		organizationPatchCoinHandler()).Methods(http.MethodPatch)
		
	// POST /organizations_host_activities
	router.HandleFunc("/api/organizations_host_activities",
		organizationsHostActivitiesCreateHandler()).Methods(http.MethodPost)
	
	// PATCH关于推广社团时间天数设置的路由
}

func organizationsHostActivitiesCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		decoder := json.NewDecoder(req.Body)
		var organizationsHostActivities models.OrganizationsHostActivities
		err := decoder.Decode(&organizationsHostActivities)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}
		organizationsHostActivities.Timestamp = time.Now()

		models.OrganizationsHostActivitiesDAO.InsertOne(&organizationsHostActivities)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("Created", "关注活动成功", organizationsHostActivities))
	}
}

func organizationPatchCoinHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
			req.ParseForm()
			ID, _ := strconv.Atoi(mux.Vars(req)["ID"])
			_, has := models.OrganizationDAO.FindByID(ID)
			if !has {
				formatter.JSON(w, http.StatusBadRequest,
					NewJSON("bad request", "该社团不存在", nil))
				return
			}
	
			var data models.Organization
			decoder := json.NewDecoder(req.Body)
			err := decoder.Decode(&data)
			if err != nil {
				logger.E.Println(err)
				formatter.JSON(w, http.StatusBadRequest,
					NewJSON("bad request", "数据格式错误", nil))
				return
			}
			
			data.ID = ID
	
			models.OrganizationDAO.UpdateOne(&data)
	
			newData, _ := models.UserDAO.FindByID(ID) 
			formatter.JSON(w, http.StatusCreated,
				NewJSON("created", "修改社团积分成功", newData))
		
	}
}

func organizationsGetSpreadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		organizations := make([]models.Organization, 0)

		allOrganizations := models.OrganizationDAO.FindAll()
		num := len(allOrganizations)
		for i := 0; i < num; i++ {
			if allOrganizations[i].SpreadFlag == false {
				continue
			}
			k := time.Now()
			startTime := allOrganizations[i].SpreadStartTime
			spreadDay := allOrganizations[i].SpreadDay
			strHour := strconv.Itoa(spreadDay*24)+"h"
			d, _ := time.ParseDuration(strHour)
			endTime := startTime.Add(d)

			if k.After(endTime) {
				continue
			} else {
				tempOrganization := allOrganizations[i]
			    organizations = append(organizations, tempOrganization)
			}
		}
		formatter.JSON(w, http.StatusOK,
		    NewJSON("OK", "获取被推广的社团列表成功", organizations))
	}
}

func organizationsGetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if req.FormValue("aid") == "" {
			organizations := models.OrganizationDAO.FindAll()
		    formatter.JSON(w, http.StatusOK,
			NewJSON("OK", "获取社团列表成功", organizations))
		} else {
			aid, _ := strconv.Atoi(req.FormValue("aid"))
			data := models.OrganizationsHostActivitiesDAO.FindOrganizationsByActivityID(aid)
		
			formatter.JSON(w, http.StatusOK,
				NewJSON("OK", "根据活动id获取社团列表成功", data))
		}
		
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

		//获取创建社团的邀请码
		models.OrganizationDAO.InsertOne(&organization)
		organization.InvitationCode = "SYSU"+strconv.Itoa(organization.ID)
		models.OrganizationDAO.UpdateOne(&organization)


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
	
		//进行管理员权限认证
		claims := util.ParseClaims(getTokenString(req))
        if int(claims["type"].(float64)) != 0 {
			formatter.JSON(w, http.StatusUnauthorized,
				NewJSON("Unauthorized", "需要此社团管理员权限", nil))
			return
		}
		adms, has := models.AdministratorDAO.FindById(int(claims["aud"].(float64)))
		if !has {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("Bad request", "不存在此管理员", nil))
			return
		}
		if adms.OrganizationID != id {
			formatter.JSON(w, http.StatusUnauthorized,
				NewJSON("Unauthorized", "需要此社团管理员权限", nil))
			return
		}

		//

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
		var t models.UsersParticipateOrganizations
		err := decoder.Decode(&t)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", err))
			return
		}

		//进行管理员权限认证
		claims := util.ParseClaims(getTokenString(r))
        if int(claims["type"].(float64)) != 0 {
			formatter.JSON(w, http.StatusUnauthorized,
				NewJSON("Unauthorized", "需要管理员权限", nil))
			return
		}
		adms, has := models.AdministratorDAO.FindById(int(claims["aud"].(float64)))
		if !has {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("Bad request", "不存在此管理员", nil))
			return
		}
		if adms.OrganizationID != t.OrganizationID {
			formatter.JSON(w, http.StatusUnauthorized,
				NewJSON("Unauthorized", "需要此社团管理员权限", nil))
			return
		}

		
		data := models.UsersParticipateOrganizationsDAO.FindByUOID(t.UserID, t.OrganizationID)
		data.Privilege = t.Privilege
		data.Applying = false;
		models.UsersParticipateOrganizationsDAO.UpdateOne(&data)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("Created", "管理社团成员成功", data))
	}
}