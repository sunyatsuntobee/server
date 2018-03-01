package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/sunyatsuntobee/server/logger"
	"github.com/sunyatsuntobee/server/models"
)

func initCollectionMomentsRouter(router *mux.Router) {
	url := "/api/moments"

	// GET /moments{?user_id}
	router.HandleFunc(url,
		mommentsGetHandler()).Methods(http.MethodGet)
	// GET /users_like_moments{?moment_id}
	router.HandleFunc("/api/users_like_moments",
		usersLikeMommentsGetHandler()).Methods(http.MethodGet)
	// POST /moments
	router.HandleFunc(url,
		mommentsCreateHandler()).Methods(http.MethodPost)
	// POST /users_like_moments
	router.HandleFunc("/api/users_like_moments",
		usersLikeMomentsCreateHandler()).Methods(http.MethodPost)
	// DELETE /users_like_moments/{ID}
	router.HandleFunc("/api/users_like_moments",
		usersLikeMomentsDeleteHandler()).Methods(http.MethodDelete)
	// GET /moment_comments{?moment_id}
	router.HandleFunc("/api/moment_comments",
		momentCommentsGetHandler()).Methods(http.MethodGet)
	// POST /moment_comments
	router.HandleFunc("/api/momment_comments",
		mommentsCommentsCreateHandler()).Methods(http.MethodPost)
	// DELETE /moment_comments/{ID}
	router.HandleFunc("/api/users_like_moments",
		momentCommentsDeleteHandler()).Methods(http.MethodDelete)
}
func usersLikeMomentsDeleteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		likeIdInt, _ := strconv.Atoi(mux.Vars(req)["ID"])

		models.MomentCommentDAO.DeleteByID(likeIdInt)
		formatter.JSON(w, http.StatusNoContent, nil)
	}
}

func momentCommentsDeleteHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		commentIdInt, _ := strconv.Atoi(mux.Vars(req)["ID"])

		models.MomentCommentDAO.DeleteByID(commentIdInt)
		formatter.JSON(w, http.StatusNoContent, nil)
	}
}
func momentCommentsGetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		momentCommentId := req.FormValue("momment_id")
		momentCommentIdInt, _ := strconv.Atoi(momentCommentId)

		commentList :=
			models.MomentCommentDAO.FindFullByMomentID(
				momentCommentIdInt)
		formatter.JSON(w, http.StatusOK,
			NewJSON("OK", "获取评论成功", commentList))
	}
}
func usersLikeMommentsGetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		likeMommentId := req.FormValue("momment_id")
		likeMommentIdInt, _ := strconv.Atoi(likeMommentId)

		resultList :=
			models.UsersLikeMomentsDAO.FindFullByMomentID(
				likeMommentIdInt)
		formatter.JSON(w, http.StatusOK,
			NewJSON("OK", "获取点赞列表成功", resultList))
	}
}

func mommentsGetHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		userId := req.FormValue("user_id")
		userIdInt, _ := strconv.Atoi(userId)
		mommentsList := models.MomentDao.FindByUserID()
		formatter.JSON(w, http.StatusOK,
			NewJSON("OK", "成功获取用户动态", mommentsList))
	}
}

func mommentsCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		decoder := json.NewDecoder(req.Body)
		var momment models.Moment
		err := decoder.Decode(&momment)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}
		momment.Timestamp = time.Now()
		models.MomentDAO.InsertOne(&momment)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("Created", "成功创建用户动态", momment))
	}
}

func usersLikeMomentsCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		decoder := json.NewDecoder(req.Body)
		var userLikeMoments models.UsersLikeMoments
		err := decoder.Decode(&userLikeMoments)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}

		models.userLikeMomentsDao.InsertOne(&userLikeMoments)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("Created", "用户点赞成功", userLikeMoments))
	}
}

func mommentsCommentsCreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		decoder := json.NewDecoder(req.Body)
		var mommentComment models.MomentComment
		err := decoder.Decode(&mommentComment)
		if err != nil {
			logger.E.Println(err)
			formatter.JSON(w, http.StatusBadRequest,
				NewJSON("bad request", "数据格式错误", nil))
			return
		}
		mommentComment.Timestamp = time.Now()
		models.MomentCommentDAO.InsertOne(&mommentComment)
		formatter.JSON(w, http.StatusCreated,
			NewJSON("Created", "评论成功", mommentComment))
	}
}
