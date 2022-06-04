package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var Port = ":8080"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func checkUser(token string, user_id string, c *gin.Context) (*service.User, int64, bool) {
	var user *service.User
	var exist bool
	if user_id == "" {
		user, exist = service.QueryUserByToken(token)
		return user, -1, exist
	}
	uid, err := strconv.ParseInt(user_id, 10, 64)
	if err != nil {
		panic(err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User Id fail"})
	}
	if token == "" {
		user, exist = service.QueryUserById(uint(uid))
	} else {
		user, exist = service.QueryUserByTokenAndUid(token, uid)
	}
	return user, uid, exist
}
