package reaction

import (
	modelsReaction "forum/internal/models/reaction"
	"strconv"
)

func (s *ReactionService) CreatePostReaction(postID, userID int, vote string) error {
	react, err := strconv.Atoi(vote)
	// fmt.Println(react, "react")
	if err != nil {
		return err
	}

	reaction := modelsReaction.Reaction{
		PostID: postID,
		UserID: userID,
		Vote:   react,
	}
	return s.repo.CreatePostReaction(reaction)
}

func (s *ReactionService) CreateCommentReaction(commentID, userID int, vote string) (int, error) {
	react, err := strconv.Atoi(vote)
	if err != nil {
		return 0, err
	}

	reaction := modelsReaction.Reaction{
		CommentID: commentID,
		UserID:    userID,
		Vote:      react,
	}
	return s.repo.CreateCommentReaction(reaction)
}
