package repository

import (
	"fmt"
	"testing"
	"time"
)

func TestQueryFeedById(t *testing.T) {
	Init()
	feed, _ := NewFeedDaoInstance().QueryFeedById(1)
	fmt.Printf("%+v", feed)
}

func TestCreateFeed(t *testing.T) {
	Init()
	fmt.Println(db)
	feed := Feed{Author: 1, PlayUrl: "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 100,
		CommentCount:  100,
		IsFavorite:    true}
	fmt.Println(feed)
	_ = NewFeedDaoInstance().CreateFeed(&feed)
}

func TestFeedsByTime(t *testing.T) {
	Init()
	var feeds []Feed
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	feeds, _ = NewFeedDaoInstance().QueryFeedsByTime(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(len(feeds))
	for _, feed := range feeds {
		fmt.Printf("%+v\n", feed)
	}
}

func TestCreateUser(t *testing.T) {
	Init()
	fmt.Println(db)
	user := User{Name: "wangtianzi", FollowCount: 1000,
		FollowerCount: 1000, IsFollow: true}
	fmt.Println(user)
	_ = NewUserDaoInstance().CreateUser(&user)
}

func TestQueryUserById(t *testing.T) {
	Init()
	feed, _ := NewUserDaoInstance().QueryUserById(1)
	fmt.Printf("%+v", feed)
}
