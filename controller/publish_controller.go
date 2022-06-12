package controller

import (
	"douyin/service"
	"douyin/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type VideoListResponse struct {
	Response
	VideoList []service.Video `json:"video_list"`
}

func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	user, exist := service.QueryUserByToken(token)
	if !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s_%d", user.Id, filename, time.Now().Unix())
	videoUrl := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, videoUrl); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	var fileSuffix string
	fileSuffix = path.Ext(finalName)
	filenameOnly := strings.TrimSuffix(finalName, fileSuffix)
	filePathOnly := strings.TrimSuffix(videoUrl, fileSuffix)
	_, err = util.SaveImage(videoUrl, filePathOnly)
	if err != nil {
		panic(err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Cover save fail"})
		return
	}
	if b := service.CreateVideo(user.Id, finalName, filenameOnly+".png", title); !b {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "CreateVideo save fail"})
		return
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	user_id := c.Query("user_id")
	if _, uid, exist := checkUser(token, user_id, c); exist {
		videos := service.QueryVideosByUid(uint(uid), c.ClientIP(), Port)
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 0,
			},
			VideoList: videos,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}

}
