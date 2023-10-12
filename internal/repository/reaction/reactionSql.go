package reaction

import (
	"database/sql"
	modelsReaction "forum/internal/models/reaction"
)

type Reaction interface {
	CreatePostReaction(reaction modelsReaction.Reaction) error
	CreateCommentReaction(reaction modelsReaction.Reaction) (int, error)
}

type ReactionSql struct {
	db *sql.DB
}

func NewReactionSql(db *sql.DB) *ReactionSql {
	return &ReactionSql{
		db: db,
	}
}
