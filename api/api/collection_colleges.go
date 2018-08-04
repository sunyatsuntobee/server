package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	
	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/logger"
	"github.com/sunyatsuntobee/server/models"
	"github.com/sunyatsuntobee/server/util"
)

func initCollectionCollegesRouter(router *mux.Router) {
	url := "/api/colleges"

	// GET /colleges
	router.HandleFunc(url,
		collegesGetHandler()).Methods(http.MethodGet)

	// GET /colleges/{ID}
	router.HandleFunc(url+"/{ID}",
	    collegesGetByIDHandler()).Methods(http.MethodGet)

	// PUT /colleges/{ID}
	router.HandleFunc(url+"/{ID}",
		collegesPutHandler()).Methods(http.MethodPut)
		
	// POST /colleges
	router.HandleFunc(url,
		collegesCreateHandler()).Methods(http.MethodPost)
	
	//DELETE /colleges/{ID}
	router.HandleFunc(url+"/{ID}",
		collegesDeleteHandler()).Methods(http.MethodDelete)
	
}

func collegesGetByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
        req.ParseForm()
		IDInt, _ := strconv.Atoi(mux.Vars(req)["ID"])
		college, hasCollege := models.CollegeDAO.FindByID(IDInt)
		if hasCollege == false {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "学校对象不存在", nil))
		} else {
			formatter.JSON(w, http.StatusOK,
				NewJSON("OK", "获取对应ID的学校成功", college))
		}
	}
}

func collegesGetHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
		collegeList := models.CollegeDAO.FindAll()
		formatter.JSON(w, http.StatusOK,
			NewJSON("OK", "获取学校列表成功", collegeList))
	}
}

func collegesPutHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        
		var data models.College
		id, _ := strconv.Atoi(mux.Vars(req)["ID"])
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&data)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}

		old, has := models.CollegeDAO.FindByID(id)
		if !has {
			formatter.JSON(w, http.StatusBadRequest, util.Error{
				Msg: "entity not found",
			})
		}
		old.Name = data.Name
		old.LogoURL = data.LogoURL
		old.ImageURL = data.ImageURL
		models.CollegeDAO.UpdateOne(&old)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("created", "修改学校信息成功", old))
	}
}


func collegesCreateHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		decoder := json.NewDecoder(req.Body)
		var college models.College
		err := decoder.Decode(&college)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}
		college.LogoURL = ""
		college.ImageURL = ""

		if college.Name != "" {
			_, flagName := models.CollegeDAO.FindByName(college.Name)
			if flagName {
				formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "此学校已被添加", nil))
			    return
			}
		}
		
		models.CollegeDAO.InsertOne(&college)

		formatter.JSON(w, http.StatusCreated,
			NewJSON("Created", "添加成功", college))
	}
}

func collegesDeleteHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		collegeIDInt, _ := strconv.Atoi(mux.Vars(req)["ID"])

		_,has := models.CollegeDAO.FindByID(collegeIDInt)

		if has == false {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "不存在这所学校", nil))
		}
		
		models.CollegeDAO.DeleteByID(collegeIDInt)
		models.CollegeDistrictDAO.DeleteByCID(collegeIDInt)//顺便删除此学校所有校区
		formatter.JSON(w, http.StatusNoContent, nil)
	}
}



