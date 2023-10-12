package comment

import (
	"database/sql"
	modelsComment "forum/internal/models/comment"
)

type Comment interface {
	CreateComment(comment modelsComment.Comment) error
	GetCommentsByPostID(postID, userID int) ([]modelsComment.Comment, error)
}

type CommentSql struct {
	db *sql.DB
}

func NewCommentsSql(db *sql.DB) *CommentSql {
	return &CommentSql{
		db: db,
	}
}
