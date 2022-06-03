package service

import "douyin/repository"

func CreateFavorite(user int64, feed int64) bool {
	favorite := repository.Favorite{User: user, Feed: feed, Cancel: 0}
	err := repository.NewFavoriteDaoInstance().CreateFavorite(&favorite)
	if err != nil {
		return false
	}
	return true
}

func DeleteFavorite(user int64, feed int64) bool {
	err := repository.NewFavoriteDaoInstance().DeleteFavorite(user, feed)
	if err != nil {
		return false
	}
	return true
}
