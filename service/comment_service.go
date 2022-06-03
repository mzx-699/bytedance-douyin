package service

import (
	"douyin/repository"
)

type ComService struct {
}

func CreateComment(user *User, vid int64, comment_text string) (comments []Comment) {
	rcomment := repository.Comment{User: user.Id, Feed: vid,
		Content: comment_text}
	if err := repository.NewCommentDaoInstance().CreateComment(&rcomment); err != nil {
		return nil
	}
	comment := Comment{Id: int64(rcomment.ID),
		User:       *user,
		Content:    rcomment.Content,
		CreateDate: rcomment.CreatedAt.Format("01-02")}
	comments = append(comments, comment)
	return comments
}

func DeleteComment(cid int64, vid int64) bool {
	if err := repository.NewCommentDaoInstance().DeleteComment(cid, vid); err != nil {
		return false
	}
	return true
}

func QueryCommentsByVid(vid int64) (comments []Comment, res bool) {
	rcomments, err := repository.NewCommentDaoInstance().QueryCommentsByVid(vid)
	if err != nil {
		return nil, false
	}
	return new(ComService).convert(rcomments), true
}
func (ComService) convert(rcomments []repository.Comment) (comments []Comment) {
	for _, rcomment := range rcomments {
		user, _ := QueryUserById(rcomment.User)
		comment := Comment{Id: int64(rcomment.ID),
			User:       *user,
			Content:    rcomment.Content,
			CreateDate: rcomment.CreatedAt.Format("01-02")}
		comments = append(comments, comment)
	}
	return comments
}
