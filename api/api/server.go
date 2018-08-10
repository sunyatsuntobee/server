package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/sunyatsuntobee/server/models"
)

const (
	cor string = "*"
)

// NewServer creates a new Negroni Server for RESTful api
func NewServer() *negroni.Negroni {
	n := negroni.Classic()
	router := mux.NewRouter()

	InitRouter(router)

	n.UseHandler(router)

	InitSql();

	return n
}

// InitRouter init router for RESTful API
func InitRouter(router *mux.Router) {
	// Resources files
	router.PathPrefix("/res").Handler(http.StripPrefix("/res/",
		http.FileServer(http.Dir(resDir))))

	initAuthorizationRouter(router)
	initCollectionUsersRouter(router)
	initCollectionActivitiesRouter(router)
	initCollectionActivityStagesRouter(router)
	initCollectionPhotoLivesRouter(router)
	initCollectionPhotosRouter(router)
	initCollectionOrganizationsRouter(router)
	initCollectionMomentsRouter(router)
	initCollectionCollegesRouter(router)
}

func optionsHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", cor)
		w.WriteHeader(http.StatusNoContent)
	}

}

func InitSql() {
	data := models.ActivityStageDAO.FindAll()
	for _,v := range data {
		activity,_ := models.ActivityDAO.FindByID(v.ActivityID)
		if v.WechatURL == "" {
			v.WechatURL =  activity.WechatURL
		}
			
		models.ActivityStageDAO.UpdateOne(&v)
	}
	activities := models.ActivityDAO.FindAll()
	for _,v := range activities {
		if v.School == "" {
			v.School =  "中山大学"
		}	
		models.ActivityDAO.UpdateOne(&v)
	}
}
