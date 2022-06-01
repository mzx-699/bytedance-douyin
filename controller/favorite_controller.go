package controller

//func FavoriteAction(c *gin.Context) {
//	token := c.Query("token")
//	videoId := c.Query("video_id")
//	actionType := c.Query("action_type")
//
//	if _, exist := usersLoginInfo[token]; exist {
//		c.JSON(http.StatusOK, Response{StatusCode: 0})
//	} else {
//		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
//	}
//}
//
//func FavoriteList(c *gin.Context) {
//	c.JSON(http.StatusOK, VideoListResponse{
//		Response: Response{
//			StatusCode: 0,
//		},
//		VideoList: DemoVideos,
//	})
//}
