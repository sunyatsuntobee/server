package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/models"
	"github.com/sunyatsuntobee/server/util"
)

func initAuthorizationRouter(router *mux.Router) {
	router.HandleFunc("/api/auth", authHandler()).
		Methods(http.MethodGet)
}

func authHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var flag bool
		var right_password string
		req.ParseForm()
		//users := models.UserDAO.FindAll()
		type_, _ := strconv.Atoi(req.FormValue("type"))

		switch type_ {
		case 0:
			_, flag = models.AdministratorDAO.FindByUsername(req.FormValue("username"))
		case 1:
			_, flag = models.UserDAO.FindByPhone(req.FormValue("username"))
		case 2:
			_, flag = models.OrganizationDAO.FindByPhone(req.FormValue("username"))
		}

		if flag == false {
			formatter.JSON(w, http.StatusBadRequest, util.Error{
				Msg: "Username not found",
			})
		}

		switch type_ {
		case 0:
			administrator, _ := models.AdministratorDAO.FindByUsername(req.FormValue("username"))
			right_password := administrator.Password
		case 1:
			user, _ := models.UserDAO.FindByPhone(req.FormValue("username"))
			right_password := user.Password
		case 2:
			organization, _ := models.OrganizationDAO.FindByPhone(req.FormValue("username"))
			right_password := organization.Password
		}

		inputpassword := req.FormValue("password")

		if right_password != inputpassword {
			formatter.JSON(w, http.StatusBadRequest, util.Error{
				Msg: "Password not right",
			})
		}

		id, _ := strconv.Atoi(req.FormValue("id"))
		Signed := util.NewJWT(id, type_)
		formatter.JSON(w, http.StatusOK, Signed)
	}
}
