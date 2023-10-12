package repository

import (
	"database/sql"
	authRepository "forum/internal/repository/auth"
	commentRepository "forum/internal/repository/comment"
	postRepository "forum/internal/repository/post"
	reactionRepository "forum/internal/repository/reaction"
)

type Repository struct {
	authRepository.Authorization
	postRepository.Post
	commentRepository.Comment
	reactionRepository.Reaction
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: authRepository.NewAuthSql(db),
		Post:          postRepository.NewPostSql(db),
		Comment:       commentRepository.NewCommentsSql(db),
		Reaction:      reactionRepository.NewReactionSql(db),
	}
}
