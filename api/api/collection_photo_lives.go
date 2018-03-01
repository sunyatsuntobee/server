package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/logger"
	"github.com/sunyatsuntobee/server/models"
)

func initCollectionPhotoLivesRouter(router *mux.Router) {
	url := "/api/photo_lives"

	// POST /photolives
	router.HandleFunc(url,
		photoLivesPostHandler()).Methods(http.MethodPost)

	// PUT /photolives/{ID}
	router.HandleFunc(url+"/{ID}",
		photoLivesPutHandler()).Methods(http.MethodPut)
}

func photoLivesPostHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		var photoLive models.PhotoLive
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&photoLive)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}
		models.PhotoLiveDAO.InsertOne(&photoLive)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("created", "创建照片直播成功", photoLive))
	}

}

func photoLivesPutHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		ID, _ := strconv.Atoi(mux.Vars(req)["ID"])
		var photoLive models.PhotoLive
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&photoLive)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}
		photoLive.ID = ID
		models.PhotoLiveDAO.UpdateOne(&photoLive)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("created", "修改照片直播成功", photoLive))
	}

}
