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
		var rightpassword string
		var id int
		req.ParseForm()
		typ, _ := strconv.Atoi(req.FormValue("type"))

		switch typ {
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

		switch typ {
		case 0:
			administrator, _ := models.AdministratorDAO.FindByUsername(req.FormValue("username"))
			rightpassword = util.MD5Hash(administrator.Password)
			id = administrator.ID
		case 1:
			user, _ := models.UserDAO.FindByPhone(req.FormValue("username"))
			rightpassword = util.MD5Hash(user.Password)
			id = user.ID
		case 2:
			organization, _ := models.OrganizationDAO.FindByPhone(req.FormValue("username"))
			rightpassword = util.MD5Hash(organization.Password)
			id = organization.ID
		}

		inputpassword := req.FormValue("password")

		if rightpassword != inputpassword {
			formatter.JSON(w, http.StatusBadRequest, util.Error{
				Msg: "Password not right",
			})
		}

		Signed := util.NewJWT(id, typ)
		formatter.JSON(w, http.StatusOK, Signed)
	}
}
