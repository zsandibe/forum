package comment

import (
	"errors"
	modelsComment "forum/internal/models/comment"
	"strings"
)

func (s *CommentService) CreateComment(comment modelsComment.Comment) error {
	if strings.TrimSpace(comment.Body) == "" {
		return errors.New("comment body cannot be empty")
	}
	return s.repo.CreateComment(comment)
}

func (s *CommentService) GetCommentsByPostID(postID, userID int) ([]modelsComment.Comment, error) {
	// fmt.Println(postID, userID)
	return s.repo.GetCommentsByPostID(postID, userID)
}
