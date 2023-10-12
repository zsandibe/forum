package comment

import (
	modelsComment "forum/internal/models/comment"
	repositoryComment "forum/internal/repository/comment"
)

type Comment interface {
	CreateComment(comment modelsComment.Comment) error
	GetCommentsByPostID(postID, userID int) ([]modelsComment.Comment, error)
}

type CommentService struct {
	repo repositoryComment.Comment
}

func NewCommentService(repo repositoryComment.Comment) *CommentService {
	return &CommentService{repo: repo}
}
