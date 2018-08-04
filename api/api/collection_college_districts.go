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

func initCollectionCollegeDistrictsRouter(router *mux.Router) {
	url := "/api/college_districts"

	// GET /college_districts
	router.HandleFunc(url,
		collegeDistrictsGetHandler()).Methods(http.MethodGet)

	// GET /college_districts/{ID}
	router.HandleFunc(url+"/{ID}",
	    collegeDistrictsGetByIDHandler()).Methods(http.MethodGet)

	// PUT /college_districts/{ID}
	router.HandleFunc(url+"/{ID}",
		collegeDistrictsPutHandler()).Methods(http.MethodPut)
		
	// POST /college_districts
	router.HandleFunc(url,
		collegeDistrictsCreateHandler()).Methods(http.MethodPost)
	
	//DELETE /college_districts/{ID}
	router.HandleFunc(url+"/{ID}",
		collegeDistrictsDeleteHandler()).Methods(http.MethodDelete)
	
}

func collegeDistrictsGetByIDHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        req.ParseForm()
		IDInt, _ := strconv.Atoi(mux.Vars(req)["ID"])
		collegeDistrict, hasCollegeDistrict := models.CollegeDistrictDAO.FindByID(IDInt)
		if hasCollegeDistrict == false {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "校区对象不存在", nil))
		} else {
			formatter.JSON(w, http.StatusOK,
				NewJSON("OK", "获取对应ID的校区成功", collegeDistrict))
		}
	}
}

func collegeDistrictsGetHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
		collegeDistrictList := models.CollegeDistrictDAO.FindAll()
		formatter.JSON(w, http.StatusOK,
			NewJSON("OK", "获取校区列表成功", collegeDistrictList))
	}
}

func collegeDistrictsPutHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        
		var data models.CollegeDistricts
		id, _ := strconv.Atoi(mux.Vars(req)["ID"])
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&data)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}

		old, has := models.CollegeDistrictDAO.FindByID(id)
		if !has {
			formatter.JSON(w, http.StatusBadRequest, util.Error{
				Msg: "entity not found",
			})
		}
		old.Name = data.Name
		old.CollegeID = data.CollegeID
		models.CollegeDistrictDAO.UpdateOne(&old)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("created", "修改校区信息成功", old))
	}
}

func collegeDistrictsCreateHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		decoder := json.NewDecoder(req.Body)
		var collegeDistrict models.CollegeDistrict
		err := decoder.Decode(&collegeDistrict)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}
		
	    _, hasCollege := models.CollegeDAO.FindByID(collegeDistrict.CollegeID)
		if !hasCollege {
            formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "不存在此学校", nil))
			return
		}

		hasCollegeDistrict := models.CollegeDistrictDAO.HasCreate(collegeDistrict.CollegeID, collegeDistrict.Name)
		if hasCollegeDistrict {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "此校区已被创建", nil))
			return
		}
		
		models.CollegeDistrictDAO.InsertOne(&collegeDistrict)

		formatter.JSON(w, http.StatusCreated,
			NewJSON("Created", "添加成功", college))
	}
}

func collegeDistrictsDeleteHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		id, _ := strconv.Atoi(mux.Vars(req)["ID"])

		college,has := models.CollegeDistrictDAO.FindByID(id)

		if has == false {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "不存在这个校区", nil))
		}
		
		models.CollegeDistrictDAO.DeleteByID(id)
		formatter.JSON(w, http.StatusNoContent, nil)
	}
}





