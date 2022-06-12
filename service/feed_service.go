package service

import (
	"douyin/repository"
	"fmt"
	"strings"
	"time"
)

type FeeService struct {
}

func CreateVideo(userId uint, playUrl string, coverUrl string, title string) bool {
	feed := repository.Feed{UserId: userId, PlayUrl: playUrl,
		CoverUrl:      coverUrl,
		Title:         title,
		FavoriteCount: 0,
		CommentCount:  0,
	}
	err := repository.NewFeedDaoInstance().CreateFeed(&feed)
	if err != nil {
		panic(err)
		return false
	}
	return true
}

func QueryVideos(token string, latest int64, ip string, port string) (videos []Video, nextTime int64) {
	t := time.Unix(latest/1000, 0)
	var feeds []repository.Feed
	feeds, _ = repository.NewFeedDaoInstance().QueryVideosByTime(t.Format("2006-01-02 15:04:05"))
	if token == "" {
		videos, nextTime = new(FeeService).convert(0, feeds, ip, port)
	} else {
		user, _ := repository.NewUserDaoInstance().QueryUserByToken(token)
		if user != nil {
			videos, nextTime = new(FeeService).convert(user.ID, feeds, ip, port)
		} else {
			videos, nextTime = new(FeeService).convert(0, feeds, ip, port)
		}
	}
	return videos, nextTime
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

func QueryVideosByUid(uid uint, ip string, port string) []Video {
	var feeds []repository.Feed
	feeds, _ = repository.NewFeedDaoInstance().QueryVideosByUid(uid)
	videos, _ := new(FeeService).convert(uid, feeds, ip, port)
	return videos
}

func QueryFavoirteVideosByUid(uid uint, ip string, port string) ([]Video, bool) {
	var feeds []repository.Feed
	feeds, err := repository.NewFeedDaoInstance().QueryFavoirteVideosByUid(uid)
	if err != nil {
		return nil, false
	}
	videos, _ := new(FeeService).convert(uid, feeds, ip, port)
	return videos, true
}

func (FeeService) convert(uid uint, feeds []repository.Feed, ip string, port string) ([]Video, int64) {
	var videos []Video
	var ret_latest int64 = time.Now().Unix()
	for _, feed := range feeds {
		isFavorite, _ := repository.NewFavoriteDaoInstance().CheckFavorite(uid, feed.ID)
		isFollow, _ := repository.NewRelationDaoInstance().CheckRelation(feed.UserId, uid)
		feed := HandlerUrl(&feed, ip, port)
		videos = append(videos, Video{Id: int64(feed.ID),
			CommentCount:  feed.CommentCount,
			FavoriteCount: feed.FavoriteCount,
			Title:         feed.Title,
			PlayUrl:       feed.PlayUrl,
			CoverUrl:      feed.CoverUrl,
			IsFavorite:    isFavorite,
			Author: User{Id: feed.User.ID,
				Name:          feed.User.Name,
				IsFollow:      isFollow || feed.UserId == uid,
				FollowerCount: feed.User.FollowerCount,
				FollowCount:   feed.User.FollowCount}})
		ret_latest = feed.CreatedAt.Unix() * 1000
	}
	return videos, ret_latest
}
