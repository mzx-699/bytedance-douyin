package repository

import (
	"douyin/util"
	"gorm.io/gorm"
	"sync"
)

type Favorite struct {
	gorm.Model
	User   uint  `gorm:"column:user"`
	Feed   uint  `gorm:"column:feed"`
	Cancel int64 `gorm:"column:cancel"`
}

func (Favorite) TableName() string {
	return "favorites"
}

type FavoriteDao struct {
}

var favoriteDao *FavoriteDao
var favoriteOnce sync.Once

func NewFavoriteDaoInstance() *FavoriteDao {
	feedOnce.Do(
		func() {
			favoriteDao = &FavoriteDao{}
		})
	return favoriteDao
}

func (*FavoriteDao) CreateFavorite(favoirte *Favorite) error {
	tx := db.Begin()
	// 查找是否有一个被取消的记录
	err := db.Where("user = ? AND feed = ?", favoirte.User, favoirte.Feed).First(&Favorite{}).Error
	// 当前没有这个记录
	if err == gorm.ErrRecordNotFound {
		// 为空则创建
		if err := db.Create(favoirte).Error; err != nil {
			util.Logger.Error("CreateFavorite err:" + err.Error())
			tx.Rollback()
			return err
		}
	} else if err != nil { //其他错误 回滚
		util.Logger.Error("CreateFavorite err:" + err.Error())
		tx.Rollback()
		return err
	} else {
		if err := db.Model(Favorite{}).Where("user = ? and feed = ? AND cancel = 1", favoirte.User, favoirte.Feed).
			Update("cancel", gorm.Expr("cancel & 0")).Error; err != nil {
			util.Logger.Error("CreateFavorite err:" + err.Error())
			tx.Rollback()
			return err
		}
	}
	// 此时已经存在一个记录了
	if err := db.Model(Feed{}).Where("id = ?", favoirte.Feed).
		Update("favorite_count", gorm.Expr("favorite_count + 1")).Error; err != nil {
		util.Logger.Error("AddVideoFavoriteCount err:" + err.Error())
		db.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (*FavoriteDao) DeleteFavorite(user uint, feed uint) error {
	tx := db.Begin()
	if err := db.Model(Favorite{}).Where("user = ? AND feed = ?", user, feed).
		Update("cancel", gorm.Expr("cancel | 1")).Error; err != nil {
		util.Logger.Error("DeleteFavorite err:" + err.Error())
		tx.Rollback()
		return err
	}
	if err := db.Model(Feed{}).Where("id = ? && favorite_count > 0", feed).
		Update("favorite_count", gorm.Expr("favorite_count - 1")).Error; err != nil {
		util.Logger.Error("DelVideoFavoriteCount err:" + err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (*FavoriteDao) CheckFavorite(user uint, feed uint) (bool, error) {
	if user <= 0 {
		return false, nil
	}
	var favorite Favorite
	err := db.Where("user = ? and feed = ?", user, feed).First(&favorite).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		util.Logger.Error("CheckFavorite err:" + err.Error())
		return false, err
	}
	return true, nil
}
