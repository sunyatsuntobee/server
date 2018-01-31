package controllers

import (
	"net/http"

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
			organizations := make([]models.Organization, 0)
			data := models.Organization{
				ID:          0,
				Name:        "自嗨社",
				Collage:     "深圳中学",
				Description: "简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介",
				LogoURL:     "/static/assets/tobee.png",
			}
			organizations = append(organizations, data)
			formatter.HTML(w, http.StatusOK,
				"console/organizations/organizations_list", organizations)
		} else {
			data := OrganizationDetail{
				Organization: &models.Organization{
					ID:          0,
					Name:        "自嗨社",
					Collage:     "深圳中学",
					Description: "介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简介简",
					LogoURL:     "/static/assets/tobee.png",
				},
				Contactors: []*models.User{
					&models.User{
						ID:       0,
						Username: "r0beRT",
						Location: "广东省深圳市",
						Phone:    "12345678901",
					},
				},
			}
			formatter.HTML(w, http.StatusOK,
				"console/organizations/organization", data)
		}
	}

}
