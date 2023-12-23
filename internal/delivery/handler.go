package delivery

import (
	authHandler "forum/internal/delivery/auth"
	postHandler "forum/internal/delivery/post"
	reactionHandler "forum/internal/delivery/reaction"
	reportHandler "forum/internal/delivery/report"
	requestHandler "forum/internal/delivery/request"
	s "forum/internal/service"
)

type Handler struct {
	Auth     *authHandler.AuthHandler
	Post     *postHandler.PostHandler
	Reaction *reactionHandler.ReactionHandler
	Request  *requestHandler.RequestHandler
	Report   *reportHandler.ReportHandler
	service  *s.Service
}

func NewHandler(service *s.Service) *Handler {
	return &Handler{
		Auth:     authHandler.NewAuthHandler(service.Auth),
		Post:     postHandler.NewPostHandler(service.Post, service.Auth, service.Comment),
		Reaction: reactionHandler.NewReactionHandler(service.Auth, service.Comment, service.Reaction),
		Request:  requestHandler.NewRequestHandler(service.Request),
		Report:   reportHandler.NewReportHandler(service.Report),
		service:  service,
	}
}
