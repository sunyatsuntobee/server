package api

import (
	"fmt"
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
		var id int
		req.ParseForm()
		inputUsername := req.FormValue("username")
		inputPassword := req.FormValue("password")
		typ, _ := strconv.Atoi(req.FormValue("type"))

		switch typ {
		case 0:
			var administrator models.Administrator
			administrator, flagUsername =
				models.AdministratorDAO.FindByUsername(inputUsername)
			if flagUsername == true {
				fmt.Println(administrator.Password, util.MD5Hash(inputPassword))
				if administrator.Password != util.MD5Hash(inputPassword) {
					flagPassword = false
				}
				id = administrator.ID
			}
		case 1:
			var user models.User
			user, flagUsername =
				models.UserDAO.FindByPhone(inputUsername)
			if flagUsername == true {
				if user.Password != util.MD5Hash(inputPassword) {
					flagPassword = false
				}
				id = user.ID
			}
		case 2:
			var organization models.Organization
			organization, flagUsername =
				models.OrganizationDAO.FindByPhone(inputUsername)
			if flagUsername == true {
				if organization.Password != util.MD5Hash(inputPassword) {
					flagPassword = false
				}
				id = organization.ID
			}
		}

		if flagUsername == false {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "用户名不存在", nil))
			return
		}
		if flagPassword == false {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "密码错误", nil))
			return
		}

		signed := util.NewJWT(id, typ)
		formatter.JSON(w, http.StatusOK, NewJSON("OK", "验证成功",
			util.JWTMessage{
				Token: signed,
			}))
	}
}
