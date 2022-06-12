package repository

import (
	"douyin/util"
	"gorm.io/gorm"
	"sync"
)

type User struct {
	gorm.Model
	Name          string `gorm:"column:name"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
	Token         string `gorm:"column:token"`
}

func (User) TableName() string {
	return "users"
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	feedOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) QueryUserById(id uint) (*User, error) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("QueryUserById err:" + err.Error())
		return nil, err
	}
	return &user, nil
}

func (*UserDao) CreateUser(user *User) error {
	if err := db.Create(user).Error; err != nil {
		util.Logger.Error("CreateUser err:" + err.Error())
		return err
	}
	return nil
}

func (*UserDao) CheckUser(username string) (bool, error) {
	err := db.Where("name = ?", username).First(&User{}).Error
	if err == gorm.ErrRecordNotFound { //不存在
		return true, nil
	}
	if err != nil {
		util.Logger.Error("CreateUser err:" + err.Error())
		return false, err
	}
	return false, nil
}

func (*UserDao) QueryUserByTokenAndUid(token string, uid int64) (*User, error) {
	var user User
	err := db.Where("id = ? and token = ?", uid, token).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("QueryUserByTokenAndUid err:" + err.Error())
		return nil, err
	}
	return &user, nil
}

func (*UserDao) QueryUserByToken(token string) (*User, error) {
	var user User
	err := db.Where("token = ?", token).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("QueryUserByToken id err:" + err.Error())
		return nil, err
	}
	return &user, nil
}

// 关注
func (*UserDao) QueryFollowsByUid(follower uint) ([]User, error) {
	var users []User
	sub := db.Table(Relation{}.TableName()).Select("follow").Where("follower = ? AND cancel = 0", follower)
	if err := db.Where("id IN (?)", sub).Find(&users).Error; err != nil {
		util.Logger.Error("QueryFollowsByUid err:" + err.Error())
		return nil, err
	}
	return users, nil
}

// 粉丝
func (*UserDao) QueryFollowersByUid(follow uint) ([]User, error) {
	var users []User
	sub := db.Table(Relation{}.TableName()).Select("follower").Where("follow = ? AND cancel = 0", follow)
	if err := db.Where("id IN (?)", sub).Find(&users).Error; err != nil {
		util.Logger.Error("QueryFollowersByUid err:" + err.Error())
		return nil, err
	}
	return users, nil
}
