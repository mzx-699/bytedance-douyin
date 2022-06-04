package service

import "douyin/repository"

func CreateFavorite(user uint, feed uint) bool {
	favorite := repository.Favorite{User: user, Feed: feed, Cancel: 0}
	err := repository.NewFavoriteDaoInstance().CreateFavorite(&favorite)
	if err != nil {
		return false
	}
	return true
}

func DeleteFavorite(user uint, feed uint) bool {
	err := repository.NewFavoriteDaoInstance().DeleteFavorite(user, feed)
	if err != nil {
		return false
	}
	return true
}
