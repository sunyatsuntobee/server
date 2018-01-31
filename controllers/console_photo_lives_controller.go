package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/models"
)

func initConsolePhotoLivesRouter(router *mux.Router) {
	router.HandleFunc("/photolives", consolePhotoLivesGetHandler()).
		Methods(http.MethodGet)
	router.HandleFunc("/photolives/add", consolePhotoLivesAddGetHandler()).
		Methods(http.MethodGet)
}

func consolePhotoLivesGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if req.FormValue("id") == "" {
			photoLives := make([]PhotoLiveDetail, 0)
			data := PhotoLiveDetail{
				PhotoLive: &models.PhotoLive{
					ID:            0,
					ExpectMembers: 100,
					AdProgress:    "谈判中111111111111111111111111111111111111111111111111111111111111111111111111111111111",
				},
				Organization: &models.Organization{
					ID:      0,
					Name:    "中大团委",
					LogoURL: "/static/assets/tobee.png",
				},
				Activity: &models.Activity{
					ID:   0,
					Name: "1758晚会",
				},
				ActivityStage: &models.ActivityStage{
					ID:        0,
					StartTime: time.Now(),
					EndTime:   time.Now(),
					Location:  "中大某地",
				},
				Manager: &models.User{
					ID:       1,
					Username: "张铁林",
					Phone:    "12345678901",
				},
				PhotographerManager: &models.User{
					ID:       1,
					Username: "张铁林",
					Phone:    "12345678901",
				},
				Supervisors: []*models.User{
					&models.User{
						ID:       1,
						Username: "张铁林",
						Phone:    "12345678901",
					},
					&models.User{
						ID:       1,
						Username: "张铁林",
						Phone:    "12345678901",
					},
				},
			}
			photoLives = append(photoLives, data)
			formatter.HTML(w, http.StatusOK,
				"console/photo_lives/photo_lives_list", photoLives)
		} else {
			id, _ := strconv.Atoi(req.FormValue("id"))
			data := PhotoLiveDetail{
				PhotoLive: &models.PhotoLive{
					ID:            id,
					ExpectMembers: 100,
					AdProgress:    "谈判中111111111111111111111111111111111111111111111111111111111111111111111111111111111",
				},
				Organization: &models.Organization{
					ID:      0,
					Name:    "中大团委",
					LogoURL: "/static/assets/tobee.png",
				},
				Activity: &models.Activity{
					ID:   0,
					Name: "1758晚会",
				},
				ActivityStage: &models.ActivityStage{
					ID:        0,
					StartTime: time.Now(),
					EndTime:   time.Now(),
					Location:  "中大某地",
				},
				Manager: &models.User{
					ID:       1,
					Username: "张铁林",
					Phone:    "12345678901",
				},
				PhotographerManager: &models.User{
					ID:       1,
					Username: "张铁林",
					Phone:    "12345678901",
				},
				Supervisors: []*models.User{
					&models.User{
						ID:       1,
						Username: "张铁林",
						Phone:    "12345678901",
					},
					&models.User{
						ID:       1,
						Username: "张铁林",
						Phone:    "12345678901",
					},
				},
			}
			formatter.HTML(w, http.StatusOK,
				"console/photo_lives/photo_live", data)
		}
	}

}

func consolePhotoLivesAddGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK,
			"console/photo_lives/photo_live_add", nil)
	}

}
