package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/models"
)

func initConsoleOrganizationsRouter(router *mux.Router) {
	router.HandleFunc("/organizations", consoleOrganizationsGetHandler()).
		Methods(http.MethodGet)
}

func consoleOrganizationsGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if req.FormValue("id") == "" {
			data := models.OrganizationDAO.FindAll()
			formatter.HTML(w, http.StatusOK,
				"console/organizations/organizations_list", data)
		} else {
			id, _ := strconv.Atoi(req.FormValue("id"))
			data := models.OrganizationFullDAO.FindFullByID(id)
			formatter.HTML(w, http.StatusOK,
				"console/organizations/organization", data)
		}
	}

}
