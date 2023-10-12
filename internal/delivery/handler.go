package delivery

import (
	authHandler "forum/internal/delivery/auth"
	postHandler "forum/internal/delivery/post"
	reactionHandler "forum/internal/delivery/reaction"
	s "forum/internal/service"
)

type Handler struct {
	Auth     *authHandler.AuthHandler
	Post     *postHandler.PostHandler
	Reaction *reactionHandler.ReactionHandler
	service  *s.Service
}

func NewHandler(service *s.Service) *Handler {
	return &Handler{
		Auth:     authHandler.NewAuthHandler(service.Auth),
		Post:     postHandler.NewPostHandler(service.Post, service.Auth, service.Comment),
		Reaction: reactionHandler.NewReactionHandler(service.Auth, service.Comment, service.Reaction),
		service:  service,
	}
}
