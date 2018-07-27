package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	//"strings"
	//"time"
	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/logger"
	"github.com/sunyatsuntobee/server/models"
	"github.com/sunyatsuntobee/server/util"
)

func initCollectionAddressListsRouter(router *mux.Router) {
	url := "/api/addresslists"

	// GET /addresslists/{ID}
	router.HandleFunc(url+"/{ID}",
	    addressListsGetByIDHandler()).Methods(http.MethodGet)

	// PUT /addresslists/{ID}
	router.HandleFunc(url+"/{ID}",
		addressListsPutHandler()).Methods(http.MethodPut)

	// POST /addresslists
	router.HandleFunc(url,
		addressListsCreateHandler()).Methods(http.MethodPost)

	// GET /addresslists
	router.HandleFunc(url,
		addressListsGetHandler()).Methods(http.MethodGet)

	// PATCH /addresslists/{ID}
	router.HandleFunc(url+"/{ID}",
		addressListsPatchHandler()).Methods(http.MethodPatch)
}

func addressListsCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		decoder := json.NewDecoder(req.Body)

		var addressList models.DepartmentAddressList

		err := decoder.Decode(&addressList)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}


		_, flagName := models.DepartmentAddressListDAO.FindByName(addressList.Name)
		if flagName == true {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "通讯录中已包含该姓名的用户", nil))
			return
		}

		models.DepartmentAddressListDAO.InsertOne(&addressList)

		formatter.JSON(w, http.StatusCreated,
			NewJSON("Created", "注册成功", addressList))
	}
}

func addressListsPatchHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		ID, _ := strconv.Atoi(mux.Vars(req)["ID"])
		_, has := models.DepartmentAddressListDAO.FindByID(ID)
		if !has {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "用户对象不存在", nil))
			return
		}

		var data models.DepartmentAddressList
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&data)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}

		data.ID = ID
		models.DepartmentAddressListDAO.UpdateOne(&data)

		newAddressList, _ := models.DepartmentAddressListDAO.FindByID(ID) //这样可以得到最新的吗

		formatter.JSON(w, http.StatusCreated,
			NewJSON("created", "用户修改自身通讯录信息成功", newAddressList))
	}
}

func addressListsGetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		addressList := models.DepartmentAddressListDAO.FindAll()
		formatter.JSON(w, http.StatusOK,
			NewJSON("OK", "获取通讯录表格成功", addressList))
	}
}

func addressListsGetByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		ID, _ := strconv.Atoi(mux.Vars(req)["ID"])
		addressList, has := models.DepartmentAddressListDAO.FindByID(ID)

		if !has {
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "该通讯录条目不存在", nil))
			return
		}
		
		formatter.JSON(w, http.StatusOK,
			NewJSON("OK", "获取通讯录条目成功", addressList))
	}
}


func addressListsPutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var data models.DepartmentAddressList
		id, _ := strconv.Atoi(mux.Vars(req)["ID"])
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&data)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}

		_, has := models.DepartmentAddressListDAO.FindByID(id)
		if !has {
			formatter.JSON(w, http.StatusBadRequest, util.Error{
				Msg: "entity not found",
			})
		}
		
		data.ID = id
		models.DepartmentAddressListDAO.UpdateOne(&data)

		newAddressList, _ := models.DepartmentAddressListDAO.FindByID(id) //这样可以得到最新的吗

		formatter.JSON(w, http.StatusCreated,
			NewJSON("created", "用户修改自身通讯录信息成功", newAddressList))
	}
}