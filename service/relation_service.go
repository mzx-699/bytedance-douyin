package service

import "douyin/repository"

type RelService struct {
}

func CreateRelation(follow uint, follower uint) bool {
	relation := repository.Relation{Follow: follow, Follower: follower, Cancel: 0}
	err := repository.NewRelationDaoInstance().CreateRelation(&relation)
	if err != nil {
		return false
	}
	return true
}

// 被关注 / 粉丝
func DeleteRelation(follow uint, follower uint) bool {
	err := repository.NewRelationDaoInstance().DeleteRelation(follow, follower)
	if err != nil {
		return false
	}
	return true
}

// 找关注
func QueryFollowsByUid(follower uint) ([]User, bool) {
	users, err := repository.NewUserDaoInstance().QueryFollowsByUid(follower)
	if err != nil {
		return nil, false
	}
	return new(RelService).convert(0, users, true), true
}

//找粉丝
func QueryFollowersByUid(follow uint) ([]User, bool) {
	users, err := repository.NewUserDaoInstance().QueryFollowersByUid(follow)
	if err != nil {
		return nil, false
	}
	// 关注 要变成 粉丝
	return new(RelService).convert(follow, users, false), true
}

func (RelService) convert(follower uint, rusers []repository.User, isFollow bool) (users []User) {
	for _, ruser := range rusers {
		if !isFollow { //如果是找粉丝
			// 检查是否关注了粉丝
			isFollow, _ = repository.NewRelationDaoInstance().CheckRelation(ruser.ID, follower)
		}
		user := User{Id: ruser.ID,
			Name:          ruser.Name,
			FollowCount:   ruser.FollowCount,
			FollowerCount: ruser.FollowerCount,
			IsFollow:      isFollow}
		users = append(users, user)
	}
	return users
}
