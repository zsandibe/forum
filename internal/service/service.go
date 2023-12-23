package service

import (
	r "forum/internal/repository"
	authService "forum/internal/service/auth"
	commentService "forum/internal/service/comment"
	postService "forum/internal/service/post"
	reactionService "forum/internal/service/reaction"
	reportService "forum/internal/service/report"
	requestService "forum/internal/service/request"
)

type Service struct {
	authService.Auth
	postService.Post
	commentService.Comment
	reactionService.Reaction
	reportService.Report
	requestService.Request
}

func NewService(repo *r.Repository) *Service {
	return &Service{
		Auth:     authService.NewAuthService(repo.Authorization),
		Post:     postService.NewPostService(repo.Post),
		Comment:  commentService.NewCommentService(repo.Comment),
		Reaction: reactionService.NewReactionService(repo.Reaction),
		Report:   reportService.NewReportService(repo.Report),
		Request:  requestService.NewRequestService(repo.Request),
	}
}
