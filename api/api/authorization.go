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
		var flagUsername bool
		var flagPassword bool
		var rightPassword string
		var id int
		req.ParseForm()
		var inputUsername string = req.FormValue("username")
		var inputPassword string = req.FormValue("password")
		typ, _ := strconv.Atoi(req.FormValue("type"))

		switch typ {
		case 0:
			administrator, flagUsername :=
				models.AdministratorDAO.FindByUsername(inputUsername)
			if flagUsername == true {
				rightPassword = util.MD5Hash(administrator.Password)
				if rightPassword != inputPassword {
					flagPassword = false
				}
				id = administrator.ID
			}
		case 1:
			user, flagUsername :=
				models.UserDAO.FindByPhone(inputUsername)
			if flagUsername == true {
				rightPassword = util.MD5Hash(user.Password)
				if rightPassword != inputPassword {
					flagPassword = false
				}
				id = user.ID
			}
		case 2:
			organization, flagUsername :=
				models.OrganizationDAO.FindByPhone(inputUsername)
			if flagUsername == true {
				rightPassword = util.MD5Hash(organization.Password)
				if rightPassword != inputPassword {
					flagPassword = false
				}
				id = organization.ID
			}
		}

		if flagUsername == false {
			formatter.JSON(w, http.StatusBadRequest, util.Error{
				Msg: "Username not found",
			})
			return
		}
		if flagPassword == false {
			formatter.JSON(w, http.StatusBadRequest, util.Error{
				Msg: "Password not right",
			})
			return
		}

		signed := util.NewJWT(id, typ)
		formatter.JSON(w, http.StatusOK, signed)
	}
}
