package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	Response
	CommentList []service.Comment `json:"comment_list"`
}

func CommentList(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusOK, CommentListResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "succ",
			},
			CommentList: nil,
		})
		return
	}
	_, _, exsit := checkUser(token, "", c)
	if !exsit {
		c.JSON(http.StatusOK, CommentListResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "succ",
			},
			CommentList: nil,
		})
		return
	}
	video_id := c.Query("video_id")
	vid, err := strconv.ParseInt(video_id, 10, 64)
	if err != nil {
		panic(err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Video Id is false"})
	}
	if comments, success := service.QueryCommentsByVid(vid); !success {
		c.JSON(http.StatusOK, CommentListResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "succ",
			},
			CommentList: comments,
		})
		return
	}
	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "FavoriteList fail"})

}

func CommentAction(c *gin.Context) {
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	user_id := c.Query("user_id")
	user, _, exist := checkUser(token, user_id, c)
	if !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
	vid, err := strconv.ParseInt(video_id, 10, 64)
	if err != nil {
		panic(err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Video Id is false"})
	}

	at, err := strconv.ParseInt(action_type, 10, 64)
	if err != nil || at > 2 {
		panic(err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Action_type is false"})
	}
	if at == 1 {
		comment_text := c.Query("comment_text")
		comments := service.CreateComment(user, vid, comment_text)
		c.JSON(http.StatusOK, CommentListResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "succ",
			},
			CommentList: comments,
		})
	} else if at == 2 {
		comment_id := c.Query("comment_id")
		cid, err := strconv.ParseInt(comment_id, 10, 64)
		if err != nil {
			panic(err)
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Comment_id is false"})
		}
		res := service.DeleteComment(cid, vid)
		if res {
			c.JSON(http.StatusOK, Response{StatusCode: 0})
		} else {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Delete Comment fail"})
		}
	}
}
