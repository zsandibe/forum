package reaction

import (
	modelsComment "forum/internal/models/comment"
	modelsUser "forum/internal/models/user"
)

type reactionProvider interface {
	CreatePostReaction(postID, userID int, vote string) error
	CreateCommentReaction(commentID, userID int, vote string) (int, error)
}

type userProvider interface {
	GetUserByID(userID int) (modelsUser.User, error)
}

type commentProvider interface {
	GetCommentsByPostID(postID, userID int) ([]modelsComment.Comment, error)
}

type ReactionHandler struct {
	user     userProvider
	comment  commentProvider
	reaction reactionProvider
}

func NewReactionHandler(user userProvider, comment commentProvider, reaction reactionProvider) *ReactionHandler {
	return &ReactionHandler{
		user:     user,
		comment:  comment,
		reaction: reaction,
	}
}
