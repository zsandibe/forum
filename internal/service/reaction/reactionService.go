package reaction

import (
	repositoryReaction "forum/internal/repository/reaction"
)

type Reaction interface {
	CreatePostReaction(postID, userID int, vote string) error
	CreateCommentReaction(commentID, userID int, vote string) (int, error)
}

type ReactionService struct {
	repo repositoryReaction.Reaction
}

func NewReactionService(repo repositoryReaction.Reaction) *ReactionService {
	return &ReactionService{
		repo: repo,
	}
}
