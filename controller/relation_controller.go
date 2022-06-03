package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RelationListResponse struct {
	Response
	UserList []service.User `json:"user_list"`
}

func FollowList(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusOK, RelationListResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "succ",
			},
			UserList: nil,
		})
		return
	}
	// 当前登陆用户
	user, _, exsit := checkUser(token, "", c)
	if !exsit {
		c.JSON(http.StatusOK, RelationListResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "succ",
			},
			UserList: nil,
		})
		return
	}

	if users, success := service.QueryFollowsByUid(user.Id); !success {
		c.JSON(http.StatusOK, RelationListResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "succ",
			},
			UserList: users,
		})
		return
	}
	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "FollowList fail"})
}
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusOK, RelationListResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "succ",
			},
			UserList: nil,
		})
		return
	}
	user, _, exsit := checkUser(token, "", c)
	if !exsit {
		c.JSON(http.StatusOK, RelationListResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "succ",
			},
			UserList: nil,
		})
		return
	}

	if users, success := service.QueryFollowersByUid(user.Id); !success {
		c.JSON(http.StatusOK, RelationListResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "succ",
			},
			UserList: users,
		})
		return
	}
	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "FavoriteList fail"})
}
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	to_user_id := c.Query("to_user_id")
	action_type := c.Query("action_type")
	user_id := c.Query("user_id")
	user, _, exist := checkUser(token, user_id, c)
	if !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
	tuid, err := strconv.ParseInt(to_user_id, 10, 64)
	if err != nil {
		panic(err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "To User Id is false"})
	}

	at, err := strconv.ParseInt(action_type, 10, 64)
	if err != nil || at > 2 {
		panic(err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Action_type is false"})
	}
	if at == 1 {
		//关注
		service.CreateRelation(tuid, user.Id)
	} else if at == 2 {
		//取关
		service.DeleteRelation(tuid, user.Id)
	}
	c.JSON(http.StatusOK, Response{StatusCode: 0})
}
