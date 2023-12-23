package repository

import (
	"database/sql"
	authRepository "forum/internal/repository/auth"
	commentRepository "forum/internal/repository/comment"
	postRepository "forum/internal/repository/post"
	reactionRepository "forum/internal/repository/reaction"
	reportRepository "forum/internal/repository/report"
	requestRepository "forum/internal/repository/request"
)

type Repository struct {
	authRepository.Authorization
	postRepository.Post
	commentRepository.Comment
	reactionRepository.Reaction
	reportRepository.Report
	requestRepository.Request
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: authRepository.NewAuthSql(db),
		Post:          postRepository.NewPostSql(db),
		Comment:       commentRepository.NewCommentsSql(db),
		Reaction:      reactionRepository.NewReactionSql(db),
		Report:        reportRepository.NewReportSql(db),
		Request:       requestRepository.NewRequestSql(db),
	}
}
