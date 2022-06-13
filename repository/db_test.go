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

func TestFeedsByTime(t *testing.T) {
	Init()
	var feeds []Feed
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	feeds, _ = NewFeedDaoInstance().QueryVideosByTime(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(len(feeds))
	for _, feed := range feeds {
		fmt.Printf("%+v\n", feed)
	}
}

func TestQueryUserById(t *testing.T) {
	Init()
	feed, _ := NewUserDaoInstance().QueryUserById(1)
	fmt.Printf("%+v", feed)
}

func TestQueryFavorite(t *testing.T) {
	Init()
	var favorite Favorite
	err := db.Where("user = ? AND feed = ?", 3, 18).Not("deleted_at IS NULL").First(&favorite).Error
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println(favorite)
}
