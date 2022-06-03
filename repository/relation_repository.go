package repository

import (
	"douyin/util"
	"gorm.io/gorm"
	"sync"
)

type Relation struct {
	gorm.Model
	// 被关注
	Follow int64 `gorm:"column:follow"`
	// 关注人
	Follower int64 `gorm:"column:follower"`
	Cancel   int64 `gorm:"column:cancel"`
}

func (Relation) TableName() string {
	return "relations"
}

type RelationDao struct {
}

var relationDao *RelationDao
var relationOnce sync.Once

func NewRelationDaoInstance() *RelationDao {
	feedOnce.Do(
		func() {
			relationDao = &RelationDao{}
		})
	return relationDao
}

func (*RelationDao) CreateRelation(relation *Relation) error {
	tx := db.Begin()
	// 查找是否有一个被取消的记录
	err := db.Where("follow = ? AND follower = ?", relation.Follow, relation.Follower).First(&Favorite{}).Error
	// 当前没有这个记录
	if err == gorm.ErrRecordNotFound {
		// 为空则创建
		if err := db.Create(relation).Error; err != nil {
			util.Logger.Error("CreateRelation err:" + err.Error())
			tx.Rollback()
			return err
		}
	} else if err != nil { //其他错误 回滚
		util.Logger.Error("CreateRelation err:" + err.Error())
		tx.Rollback()
		return err
	} else {
		if err := db.Model(Relation{}).Where("follow = ? AND follower = ? AND cancel = 1", relation.Follow, relation.Follower).
			Update("cancel", gorm.Expr("cancel & 0")).Error; err != nil {
			util.Logger.Error("CreateRelation err:" + err.Error())
			tx.Rollback()
			return err
		}
	}
	// 此时已经存在一个记录了
	// 被关注的粉丝数+1
	if err := db.Model(User{}).Where("id = ?", relation.Follow).
		Update("follower_count", gorm.Expr("follower_count + 1")).Error; err != nil {
		util.Logger.Error("CreateRelation err:" + err.Error())
		db.Rollback()
		return err
	}
	// 粉丝的关注数+1
	if err := db.Model(User{}).Where("id = ?", relation.Follower).
		Update("follow_count", gorm.Expr("follow_count + 1")).Error; err != nil {
		util.Logger.Error("CreateRelation err:" + err.Error())
		db.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (*RelationDao) DeleteRelation(follow int64, follower int64) error {
	tx := db.Begin()
	if err := db.Model(Relation{}).Where("follow = ? AND follower = ?", follow, follower).
		Update("cancel", gorm.Expr("cancel | 1")).Error; err != nil {
		util.Logger.Error("DeleteRelation err:" + err.Error())
		tx.Rollback()
		return err
	}
	// 被关注的粉丝-1
	if err := db.Model(User{}).Where("id = ? && follower_count > 0", follow).
		Update("follower_count", gorm.Expr("follower_count - 1")).Error; err != nil {
		util.Logger.Error("DeleteRelation err:" + err.Error())
		tx.Rollback()
		return err
	}
	// 关注的关注数-1
	if err := db.Model(User{}).Where("id = ? && follow_count > 0", follower).
		Update("follow_count", gorm.Expr("follow_count - 1")).Error; err != nil {
		util.Logger.Error("DeleteRelation err:" + err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 关注 / 粉丝
func (*RelationDao) CheckRelation(follow int64, follower int64) (bool, error) {
	if follow <= 0 || follower <= 0 {
		return false, nil
	}
	var favorite Favorite
	err := db.Where("follow = ? and follower = ?", follow, follower).First(&favorite).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		util.Logger.Error("CheckRelation err:" + err.Error())
		return false, err
	}
	return true, nil
}
