package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/logger"
	"github.com/sunyatsuntobee/server/models"
	"github.com/sunyatsuntobee/server/util"
)

func initCollectionPhotosRouter(router *mux.Router) {
	url := "/api/photos"
	router.HandleFunc(url, photosPostHandler()).Methods(http.MethodPost)
	router.HandleFunc(url+"/{ID}/photo", photosUploadHandler()).
		Methods(http.MethodPatch)
	router.HandleFunc(url, photosGetHandler()).Methods(http.MethodGet)
	router.HandleFunc(url, photosPutHandler()).Methods(http.MethodPut)
}

func photosUploadHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		id, _ := strconv.Atoi(mux.Vars(req)["ID"])
		photo, has := models.PhotoDAO.FindByID(id)
		if !has {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "照片对象不存在", nil))
			return
		}
		file, header, err := req.FormFile("photo")
		logger.LogIfError(err)
		defer file.Close()
		name := strings.Split(header.Filename, ".")
		path := resDir + "/photos/" + mux.Vars(req)["ID"] + "." + name[1]
		url := "/res/photos/" + mux.Vars(req)["ID"] + "." + name[1]
		util.SaveMultipartFile(path, file)
		photo.URL = url
		models.PhotoDAO.UpdateOne(&photo)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("created", "照片上传成功", photo))
	}

}

func photosPostHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		decoder := json.NewDecoder(req.Body)
		var data models.Photo
		err := decoder.Decode(&data)
		photo := models.NewPhoto(
			data.TookTime,
			data.TookLocation,
			data.PhotographerID)
		logger.LogIfError(err)
		models.PhotoDAO.InsertOne(photo)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("created", "照片创建成功", photo))
	}

}

func photosGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if req.FormValue("category") == "" {
			data := models.PhotoDAO.FindFullAll()
			formatter.JSON(w, http.StatusOK, data)
		} else {
			data := models.PhotoDAO.FindFullByCategory(
				req.FormValue("category"))
			formatter.JSON(w, http.StatusOK, data)
		}
	}

}

func photosPutHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		id, _ := strconv.Atoi(req.FormValue("id"))
		old, _ := models.PhotoDAO.FindByID(id)
		old.URL = req.FormValue("url")
		old.TookTime, _ = time.Parse(time.RFC3339, req.FormValue("took_time"))
		old.TookLocation = req.FormValue("took_location")
		old.ReleaseTime, _ = time.Parse(time.RFC3339,
			req.FormValue("release_time"))
		old.Category = req.FormValue("category")
		old.Likes, _ = strconv.Atoi(req.FormValue("likes"))
		old.RejectReason = req.FormValue("reject_reason")
		old.PhotographerID, _ = strconv.Atoi(req.FormValue("photographer_id"))
		models.PhotoDAO.UpdateOne(&old)
		formatter.JSON(w, http.StatusCreated, nil)
	}

}
