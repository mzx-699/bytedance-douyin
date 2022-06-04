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
	return new(UserService).convert(u), true
}

func QueryUserByToken(token string) (*User, bool) {
	u, err := repository.NewUserDaoInstance().QueryUserByToken(token)
	if err != nil || u == nil {
		return nil, false
	}
	return new(UserService).convert(u), true
}

func QueryUserById(uid uint) (*User, bool) {
	u, err := repository.NewUserDaoInstance().QueryUserById(uid)
	if err != nil || u == nil {
		return nil, false
	}

	return new(UserService).convert(u), true
}
func CheckUser(username string) bool {
	res, err := repository.NewUserDaoInstance().CheckUser(username)
	if err != nil {
		return false
	}
	return res
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

func (UserService) convert(ruser *repository.User) *User {
	user := User{
		Id:            ruser.ID,
		Name:          ruser.Name,
		FollowerCount: ruser.FollowerCount,
		FollowCount:   ruser.FollowCount,
		IsFollow:      ruser.IsFollow,
	}
	return &user
}
