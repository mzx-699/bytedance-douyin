package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "succ",
			},
			VideoList: nil,
		})
		return
	}
	user, _, _ := checkUser(token, "", c)
	if videos, success := service.QueryFavoirteVideosByUid(user.Id, c.ClientIP(), Port); success {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "succ",
			},
			VideoList: videos,
		})
		return
	}
	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "FavoriteList fail"})

}

func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	user_id := c.Query("user_id")
	user, _, exist := checkUser(token, user_id, c)
	if !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	vid, err := strconv.ParseInt(video_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Video Id is false"})
		return
	}

	at, err := strconv.ParseInt(action_type, 10, 64)
	if err != nil || at > 2 {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Action_type is false"})
		return
	}
	if at == 1 {
		service.CreateFavorite(user.Id, uint(vid))
	} else if at == 2 {
		service.DeleteFavorite(user.Id, uint(vid))
	}
	c.JSON(http.StatusOK, Response{StatusCode: 0})
}
