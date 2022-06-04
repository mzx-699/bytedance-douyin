package controller

import (
	"douyin/service"
	"douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserLoginResponse struct {
	Response
	UserId uint   `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User service.User `json:"user"`
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	user_id := c.Query("user_id")
	if user, _, exist := checkUser(token, user_id, c); exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     *user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func Register(c *gin.Context) {
	username := c.Query("username")
	if res := service.CheckUser(username); !res {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User is exist"},
		})
		return
	}
	password := c.Query("password")
	passWord, b := util.MD5PWD(password)
	if !b {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "MD5 fail"},
		})
		return
	}
	token := username + passWord

	if _, _, exist := checkUser(token, "", c); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})

	} else {
		// 创建用户
		if id, b := service.CreateUser(username, token); b {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   id,
				Token:    username + password,
			})

		} else {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "Create fail"},
			})
		}

	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	passWord, b := util.MD5PWD(password)
	if !b {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "MD5 fail"},
		})
		return
	}
	token := username + passWord

	if user, _, exist := checkUser(token, "", c); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
