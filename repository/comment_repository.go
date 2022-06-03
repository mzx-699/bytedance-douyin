package repository

import (
	"douyin/util"
	"gorm.io/gorm"
	"sync"
)

type Comment struct {
	gorm.Model
	User    int64  `gorm:"column:user"`
	Feed    int64  `gorm:"column:feed"`
	Content string `gorm:"column:content"`
}

func (Comment) TableName() string {
	return "comments"
}

type CommentDao struct {
}

var commentDao *CommentDao
var commentOnce sync.Once

func NewCommentDaoInstance() *CommentDao {
	feedOnce.Do(
		func() {
			commentDao = &CommentDao{}
		})
	return commentDao
}

func (*CommentDao) CreateComment(comment *Comment) error {
	tx := db.Begin()
	if err := db.Create(&comment).Error; err != nil {
		util.Logger.Error("CreateComment err:" + err.Error())
		tx.Rollback()
		return err
	}
	// 评论数 +1
	if err := db.Table(Feed{}.TableName()).Where("id = ?", comment.Feed).
		Delete("comment_count", gorm.Expr("comment_count + 1")).Error; err != nil {
		util.Logger.Error("CreateComment err:" + err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (*CommentDao) DeleteComment(cid int64, vid int64) error {
	tx := db.Begin()
	if err := db.Where("id = ? AND feed = ?", cid, vid).Delete(&Comment{}).Error; err != nil {
		util.Logger.Error("DeleteComment err:" + err.Error())
		tx.Rollback()
		return err
	}
	// 评论数 -1
	if err := db.Table(Feed{}.TableName()).Where("id = ?", vid).
		Delete("comment_count", gorm.Expr("comment_count - 1")).Error; err != nil {
		util.Logger.Error("DeleteComment err:" + err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (*CommentDao) QueryCommentsByVid(vid int64) ([]Comment, error) {
	var comments []Comment
	if err := db.Where("feed = ?", vid).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
