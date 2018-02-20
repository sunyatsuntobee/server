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
		var flagusername bool
		var flagpassword bool
		var rightpassword string
		var id int
		req.ParseForm()
		typ, _ := strconv.Atoi(req.FormValue("type"))

		switch typ {
		case 0:
			administrator, flagusername := models.AdministratorDAO.FindByUsername(req.FormValue("username"))
			if flagusername == true {
				rightpassword = util.MD5Hash(administrator.Password)
				if rightpassword != req.FormValue("password") {
					flagpassword = false
				}
				id = administrator.ID
			}
		case 1:
			user, flagusername := models.UserDAO.FindByPhone(req.FormValue("username"))
			if flagusername == true {
				rightpassword = util.MD5Hash(user.Password)
				if rightpassword != req.FormValue("password") {
					flagpassword = false
				}
				id = user.ID
			}
		case 2:
			organization, flagusername := models.OrganizationDAO.FindByPhone(req.FormValue("username"))
			if flagusername == true {
				rightpassword = util.MD5Hash(organization.Password)
				if rightpassword != req.FormValue("password") {
					flagpassword = false
				}
				id = organization.ID
			}
		}

		if flagusername == false {
			formatter.JSON(w, http.StatusBadRequest, util.Error{
				Msg: "Username not found",
			})
			return
		}

		if flagpassword == false {
			formatter.JSON(w, http.StatusBadRequest, util.Error{
				Msg: "Password not right",
			})
			return
		}

		Signed := util.NewJWT(id, typ)
		formatter.JSON(w, http.StatusOK, Signed)
	}
}
