package repository

import (
	"douyin/util"
	"gorm.io/gorm"
	"sync"
)

type Feed struct {
	gorm.Model
	Author        int64  `gorm:"column:author"`
	PlayUrl       string `gorm:"column:play_url"`
	CoverUrl      string `gorm:"column:cover_url"`
	Title         string `gorm:"column:title"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
	CommentCount  int64  `gorm:"column:commnet_count"`
	IsFavorite    bool   `gorm:"column:is_favorite"`
}

func (Feed) TableName() string {
	return "feeds"
}

type FeedDao struct {
}

var feedDao *FeedDao
var feedOnce sync.Once

func NewFeedDaoInstance() *FeedDao {
	feedOnce.Do(
		func() {
			feedDao = &FeedDao{}
		})
	return feedDao
}

func (*FeedDao) QueryFeedById(id int64) (*Feed, error) {
	var feed Feed
	err := db.Where("id = ?", id).First(&feed).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("QueryFeedById err:" + err.Error())
		return nil, err
	}
	return &feed, nil
}

func (*FeedDao) CreateFeed(feed *Feed) error {
	if err := db.Create(feed).Error; err != nil {
		util.Logger.Error("CreateFeed err:" + err.Error())
		return err
	}
	return nil
}

func (*FeedDao) QueryVideosByTime(t string) ([]Feed, error) {
	var feeds []Feed
	err := db.Where("created_at <= ?", t).Order("created_at desc, id").Limit(30).Find(&feeds).Error
	if err != nil {
		util.Logger.Error("QueryVideosByTime err:" + err.Error())
		return nil, err
	}
	return feeds, nil
}

func (*FeedDao) QueryVideosByToken(token string) ([]Feed, error) {
	var feeds []Feed
	sub := db.Table(User{}.TableName()).Select("id").Where("token = ?", token)
	err := db.Where("author = (?)", sub).Find(&feeds).Error
	if err != nil {
		util.Logger.Error("QueryVideosByToken err:" + err.Error())
		return nil, err
	}
	return feeds, nil
}

func (*FeedDao) QueryVideosByUid(uid int64) ([]Feed, error) {
	var feeds []Feed
	if err := db.Where("author = ?", uid).Find(&feeds).Error; err != nil {
		util.Logger.Error("QueryVideosByUid err:" + err.Error())
		return nil, err
	}
	return feeds, nil
}

func (*FeedDao) QueryFavoirteVideosByUid(uid int64) ([]Feed, error) {
	var feeds []Feed
	sub := db.Table(Favorite{}.TableName()).Select("feed").Where("user = ? AND cancel = 0", uid)
	if err := db.Where("id IN (?)", sub).Find(&feeds).Error; err != nil {
		util.Logger.Error("QueryFavoirteVideosByUid err:" + err.Error())
		return nil, err
	}
	return feeds, nil
}
