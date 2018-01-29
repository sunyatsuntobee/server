package controllers

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/models"
	"github.com/sunyatsuntobee/server/util"
)

func initLoginRouter(router *mux.Router) {
	router.HandleFunc("/login", loginGetHandler()).Methods(http.MethodGet)
	router.HandleFunc("/login", loginPostHandler()).Methods(http.MethodPost)
}

func loginGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		html, _ := ioutil.ReadFile("./views/login.html")
		page := layout{
			Title:   "登录 - 图蜂后台管理系统",
			Content: template.HTML(string(html)),
		}
		formatter.HTML(w, http.StatusOK, "layout", page)
	}

}

func loginPostHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		phone := req.FormValue("phone")
		password := req.FormValue("password")
		typ, _ := strconv.Atoi(req.FormValue("type"))
		switch typ {
		case 0:
			// entity, has := models.AdministratorDAO.FindByPhone(phone)
			break
		case 1:
			user, has := models.UserDAO.FindByPhone(phone)
			if !has {
				formatter.JSON(w, http.StatusBadRequest, util.Error{
					Msg: "User not found",
				})
			} else if password != util.MD5Hash(user.Password) {
				formatter.JSON(w, http.StatusBadRequest, util.Error{
					Msg: "Password incorrect",
				})
			} else {
				token := util.NewJWT(user.ID, typ)
				formatter.JSON(w, http.StatusCreated, util.JWTMessage{
					Token: token,
				})
			}
			break
		case 2:
			// entity, has := models.OrganzationDAO.FindByPhone(phone)
			break
		default:
			return
		}
	}

}
