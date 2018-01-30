package controllers

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func initIndexRouter(router *mux.Router) {
	// router.HandleFunc("/", secureHandler(indexHandler()))
	router.HandleFunc("/", indexHandler())
	router.HandleFunc("/profile", indexProfileHandler())
	router.HandleFunc("/photolives", indexPhotoLivesHandler())
	router.HandleFunc("/photolives/add", indexPhotoLivesAddHandler())
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
			StartTime:                "2012年01月01日 21:00",
			EndTime:                  "2019年01月01日 21:00",
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

func indexPhotoLivesAddHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "index/photo_lives_add", nil)
	}

}
