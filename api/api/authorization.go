package api

import (
	"net/http"
	"strconv"
  "io/ioutil"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/models"
	"github.com/sunyatsuntobee/server/util"
)

func initAuthorizationRouter(router *mux.Router) {
  // GET /openid{?code}
  router.HandleFunc("/api/openid", openidHandler()).
    Methods(http.MethodGet)

	router.HandleFunc("/api/auth", authHandler()).
		Methods(http.MethodGet)
}

func openidHandler() http.HandlerFunc {
  return func(w http.ResponseWriter, req *http.Request) {
    req.ParseForm()
    code := req.FormValue("code")
    if code == "" {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "获取失败", nil))
      return
    }
    resp, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=wxe394906b703e6f6d&secret=4265e01cb21d28e15e54957dc06ceaa2&js_code=" + code + "&grant_type=authorization_code")
    if err != nil {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("Weixin Server Error", "Weixin获取失败", nil))
      return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    formatter.JSON(w, http.StatusOK,
      NewJSON("OK", "获取成功", string(body)))
  }
}

func authHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		flagOpenid := true
		flagPassword := true
		var id int
		req.ParseForm()
		inputOpenid := req.FormValue("openid")
		inputPassword := req.FormValue("password")
		typ, _ := strconv.Atoi(req.FormValue("type"))
		switch typ {
		case 0:
			var administrator models.Administrator
			administrator, flagOpenid =
				models.AdministratorDAO.FindByOpenid(inputOpenid)
			if flagOpenid == true {
				if administrator.Password != util.MD5Hash(inputPassword) && administrator.Password != inputPassword{
					flagPassword = false
				}
				id = administrator.ID
			}
		case 1:
			var user models.User
			user, flagOpenid =
				models.UserDAO.FindByOpenid(inputOpenid)
			if flagOpenid == true {
				if user.Password != util.MD5Hash(inputPassword) && user.Password != inputPassword{
					flagPassword = false
				}
				id = user.ID
			}
		case 2:
			var organization models.Organization
			organization, flagOpenid =
				models.OrganizationDAO.FindByOpenid(inputOpenid)
			if flagOpenid == true {
				if organization.Password != util.MD5Hash(inputPassword) && organization.Password != inputPassword{
					flagPassword = false
				}
				id = organization.ID
			}
		}

		if flagOpenid == false {
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
