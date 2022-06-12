package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []service.Video `json:"video_list,omitempty"`
	NextTime  int64           `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {
	parm := c.Query("latest_time")
	token := c.Query("token")
	var latest int64
	if parm == "" {
		latest = time.Now().Unix() * 1000
	} else {
		latest, _ = strconv.ParseInt(parm, 10, 64)
		if latest < 100000000000 {
			latest *= 1000
		}
	}
	videos, ret_latest := service.QueryVideos(token, latest, c.ClientIP(), Port)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  ret_latest,
	})
}
