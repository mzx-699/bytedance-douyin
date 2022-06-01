package service

import (
	"douyin/repository"
)

func QueryUserByToken(token string) (User, bool) {
	u, err := repository.NewUserDaoInstance().QueryUserByToken(token)
	if err != nil {
		panic(err)
		return User{}, false
	}
	if u.ID == 0 {
		return User{}, false
	}
	user := User{
		Id:            int64(u.ID),
		Name:          u.Name,
		FollowerCount: u.FollowerCount,
		FollowCount:   u.FollowCount,
		IsFollow:      u.IsFollow,
	}
	return user, true
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
		panic(err)
		return 0, false
	}
	return user.ID, true
}
