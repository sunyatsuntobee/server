package controllers

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func initIndexRouter(router *mux.Router) {
	// router.HandleFunc("/", secureHandler(indexHandler()))
	router.HandleFunc("/", indexHandler())
	router.HandleFunc("/profile", indexProfileHandler())
	router.HandleFunc("/photolives", indexPhotoLivesHandler())
}

func indexHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		html, _ := ioutil.ReadFile("./views/index/photos.html")
		formatter.HTML(w, http.StatusOK, "index/photos", template.HTML(string(html)))
	}

}

func indexProfileHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		html, _ := ioutil.ReadFile("./views/index/profile.html")
		formatter.HTML(w, http.StatusOK, "index/profile", template.HTML(string(html)))
	}

}

func indexPhotoLivesHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		photoLives := make([]photoLive, 0)
		test := photoLive{
			ID:                       0,
			ActivityName:             "1758晚会",
			OrganizationName:         "团委",
			Location:                 "中大某地",
			StartTime:                time.Now(),
			EndTime:                  time.Now(),
			ExpectMembers:            100,
			AdProgress:               "谈判中",
			ManagerName:              "张铁林",
			ManagerPhone:             "12345678901",
			PhotographerManagerName:  "张铁林",
			PhotographerManagerPhone: "12345678901",
			Supervisors: []contact{
				contact{Name: "张铁林", Phone: "12345678901"},
				contact{Name: "张铁林", Phone: "12345678901"},
			},
		}
		photoLives = append(photoLives, test)
		formatter.HTML(w, http.StatusOK, "index/photo_lives", photoLives)
	}

}
