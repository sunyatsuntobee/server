package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/models"
	"github.com/sunyatsuntobee/server/util"
)

func initCollectionOrganizationsRouter(router *mux.Router) {
	url := "/api/organizations"
	router.HandleFunc(url, organizationsPutHandler()).Methods(http.MethodPut)
	router.HandleFunc(url+"/{ID}/departments",
		organizationsDepartmentsDeleteHandler()).Methods(http.MethodDelete)
	router.HandleFunc(url+"/{ID}/departments",
		organizationsDepartmentsPostHandler()).Methods(http.MethodPost)
	router.HandleFunc(url+"/{ID}/contacts",
		organizationsContactsDeleteHandler()).Methods(http.MethodDelete)
	router.HandleFunc(url+"/{ID}/contacts",
		organizationsContactsPostHandler()).Methods(http.MethodPost)
}

func organizationsContactsDeleteHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		id, _ := strconv.Atoi(mux.Vars(req)["ID"])
		models.OrganizationsContactorsDAO.DeleteByOID(id)
		formatter.JSON(w, http.StatusNoContent, nil)
	}

}

func organizationsContactsPostHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		oid, _ := strconv.Atoi(req.FormValue("organization_id"))
		cid, _ := strconv.Atoi(req.FormValue("contact_id"))
		models.OrganizationsContactorsDAO.InsertOne(
			&models.OrganizationsContactors{
				OrganizationID: oid,
				ContactorID:    cid,
			})
		formatter.JSON(w, http.StatusCreated, nil)
	}

}

func organizationsPutHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		id, _ := strconv.Atoi(req.FormValue("id"))
		old, has := models.OrganizationDAO.FindByID(id)
		if !has {
			formatter.JSON(w, http.StatusBadRequest, util.Error{
				Msg: "entity not found",
			})
		}
		old.Name = req.FormValue("name")
		old.Phone = req.FormValue("phone")
		old.Password = req.FormValue("password")
		old.College = req.FormValue("collage")
		path := "static/assets/" + req.FormValue("id") + ".png"
		util.SaveBase64AsPNG(req.FormValue("logo_url"), path)
		old.LogoURL = "/" + path

		old.Description = req.FormValue("description")
		models.OrganizationDAO.UpdateOne(&old)
		formatter.JSON(w, http.StatusCreated, nil)
	}

}

func organizationsDepartmentsDeleteHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(req)["ID"])
		models.OrganizationDepartmentDAO.DeleteByOID(id)
		formatter.JSON(w, http.StatusNoContent, nil)
	}

}

func organizationsDepartmentsPostHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		oid, _ := strconv.Atoi(req.FormValue("organization_id"))
		department := models.OrganizationDepartment{
			Name:           req.FormValue("name"),
			OrganizationID: oid,
		}
		models.OrganizationDepartmentDAO.InsertOne(&department)
		formatter.JSON(w, http.StatusCreated, nil)
	}

}
