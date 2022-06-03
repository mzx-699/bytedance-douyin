package service

import (
	"douyin/repository"
)

type UserService struct {
}

func QueryUserByTokenAndUid(token string, uid int64) (*User, bool) {
	u, err := repository.NewUserDaoInstance().QueryUserByTokenAndUid(token, uid)
	if err != nil || u == nil {
		return nil, false
	}
	user := User{
		Id:            int64(u.ID),
		Name:          u.Name,
		FollowerCount: u.FollowerCount,
		FollowCount:   u.FollowCount,
		IsFollow:      u.IsFollow,
	}
	return &user, true
}

func QueryUserByToken(token string) (*User, bool) {
	u, err := repository.NewUserDaoInstance().QueryUserByToken(token)
	if err != nil || u == nil {
		return nil, false
	}
	user := User{
		Id:            int64(u.ID),
		Name:          u.Name,
		FollowerCount: u.FollowerCount,
		FollowCount:   u.FollowCount,
		IsFollow:      u.IsFollow,
	}
	return &user, true
}

func QueryUserById(uid int64) (*User, bool) {
	u, err := repository.NewUserDaoInstance().QueryUserById(uid)
	if err != nil || u == nil {
		return nil, false
	}
	user := User{
		Id:            int64(u.ID),
		Name:          u.Name,
		FollowerCount: u.FollowerCount,
		FollowCount:   u.FollowCount,
		IsFollow:      u.IsFollow,
	}
	return &user, true
}

func CreateUser(name string, token string) (uint, bool) {
	user := repository.User{
		Name:          name,
		Token:         token,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
	err := repository.NewUserDaoInstance().CreateUser(&user)
	if err != nil {
		return 0, false
	}
	return user.ID, true
}
