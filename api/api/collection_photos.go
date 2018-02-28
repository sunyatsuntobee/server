package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/logger"
	"github.com/sunyatsuntobee/server/models"
	"github.com/sunyatsuntobee/server/util"
)

func initCollectionPhotosRouter(router *mux.Router) {
	url := "/api/photos"

	// POST /photos
	router.HandleFunc(url,
		photosPostHandler()).Methods(http.MethodPost)

	// PATCH /photos/{ID}/photo
	router.HandleFunc(url+"/{ID}/photo",
		photosUploadHandler()).Methods(http.MethodPatch)

	// GET /photos{?category,oid}
	router.HandleFunc(url,
		photosGetHandler()).Methods(http.MethodGet)

	// PUT /photos/{ID}
	router.HandleFunc(url+"/{ID}",
		photosPutHandler()).Methods(http.MethodPut)
}

func photosGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		category := req.FormValue("category")
		oidStr := req.FormValue("oid")
		if category == "" && oidStr == "" {
			// Both null
		} else if category == "" {
			// Category is null, oid specified
		} else if oidStr == "" {
			// OID is null, category specified
		} else {
			// Both specified
		}
	}

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
			data.UserID)
		logger.LogIfError(err)
		models.PhotoDAO.InsertOne(photo)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("created", "照片创建成功", photo))
	}

}

func photosPutHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		var photo models.Photo
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&photo)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}
		id, _ := strconv.Atoi(mux.Vars(req)["ID"])
		old, _ := models.PhotoDAO.FindByID(id)
		photo.ID = id
		photo.URL = old.URL
		models.PhotoDAO.UpdateOne(&photo)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("created", "修改照片信息成功", photo))
	}

}
