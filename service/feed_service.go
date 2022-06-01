package service

import (
	"douyin/repository"
	"fmt"
	"strings"
	"time"
)

func CreateVideo(author int32, playUrl string, coverUrl string, title string) bool {
	feed := repository.Feed{Author: author, PlayUrl: playUrl,
		CoverUrl:      coverUrl,
		Title:         title,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false}
	err := repository.NewFeedDaoInstance().CreateFeed(&feed)
	if err != nil {
		panic(err)
		return false
	}
	return true
}

func QueryVideos(latest int64, ip string, port string) ([]Video, int64) {
	var videos []Video
	t := time.Unix(latest/1000, 0)
	var feeds []repository.Feed
	feeds, _ = repository.NewFeedDaoInstance().QueryFeedsByTime(t.Format("2006-01-02 15:04:05"))
	var ret_latest int64 = time.Now().Unix()
	for _, feed := range feeds {
		user, _ := repository.NewUserDaoInstance().QueryUserById(feed.Author)
		feed := HandlerUrl(&feed, ip, port)
		videos = append(videos, Video{Id: int64(feed.ID),
			CommentCount:  feed.CommentCount,
			FavoriteCount: feed.FavoriteCount,
			Title:         feed.Title,
			PlayUrl:       feed.PlayUrl,
			CoverUrl:      feed.CoverUrl,
			Author: User{Id: int64(user.ID),
				Name:          user.Name,
				IsFollow:      user.IsFollow,
				FollowerCount: user.FollowerCount,
				FollowCount:   user.FollowCount}})
		ret_latest = feed.CreatedAt.Unix() * 1000
	}
	return videos, ret_latest
}

var StaticPath = "http://%s%s/static/%s"

func HandlerUrl(feed *repository.Feed, ip string, port string) *repository.Feed {
	if !strings.HasPrefix(feed.PlayUrl, "http") {
		feed.PlayUrl = fmt.Sprintf(StaticPath, ip, port, feed.PlayUrl)
	}
	if !strings.HasPrefix(feed.CoverUrl, "http") {
		feed.CoverUrl = fmt.Sprintf(StaticPath, ip, port, feed.CoverUrl)
	}
	return feed
}

func QueryVideosByToken(token string) []Video {
	var videos []Video
	var feeds []repository.Feed
	feeds, _ = repository.NewFeedDaoInstance().QueryVideosByToken(token)
	for _, feed := range feeds {
		user, _ := repository.NewUserDaoInstance().QueryUserById(feed.Author)
		videos = append(videos, Video{Id: int64(feed.ID),
			CommentCount:  feed.CommentCount,
			FavoriteCount: feed.FavoriteCount,
			PlayUrl:       feed.PlayUrl,
			CoverUrl:      feed.CoverUrl,
			Author:        User{Id: int64(user.ID), Name: user.Name}})
	}
	return videos
}
