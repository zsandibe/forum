package reaction

import (
	"database/sql"

	modelsReaction "forum/internal/models/reaction"
)

func (r *ReactionSql) CreatePostReaction(reaction modelsReaction.Reaction) error {
	selectQuery := `
		SELECT VOTE FROM reactions WHERE UserID = $1 AND PostID = $2
	`
	insertQuery := `
		INSERT INTO reactions (UserID, PostID, VOTE) VALUES ($1, $2, $3)
	`
	deleteQuery := `
		DELETE FROM reactions WHERE UserID = $1 AND PostID = $2 AND VOTE = $3
	`
	// for scan reactions count
	var vote int
	if err := r.db.QueryRow(selectQuery, reaction.UserID, reaction.PostID).Scan(&vote); err != nil {
		if err != sql.ErrNoRows {
			return err
		}
		if _, err := r.db.Exec(insertQuery, reaction.UserID, reaction.PostID, reaction.Vote); err != nil {
			return err
		}
	} else {
		if _, err := r.db.Exec(deleteQuery, reaction.UserID, reaction.PostID, vote); err != nil {
			return err
		}
		if vote != reaction.Vote {
			if _, err := r.db.Exec(insertQuery, reaction.UserID, reaction.PostID, reaction.Vote); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *ReactionSql) CreateCommentReaction(reaction modelsReaction.Reaction) (int, error) {
	var postID int
	var vote int
	selectQuery := `
        SELECT VOTE FROM reactions WHERE UserID = $1 AND CommentID = $2
    `
	insertQuery := `
        INSERT INTO reactions (UserID, CommentID, VOTE) VALUES ($1, $2, $3)
    `
	deleteQuery := `
        DELETE FROM reactions WHERE UserID = $1 AND CommentID = $2 AND VOTE = $3
    `
	// for get postID
	tempQuery := `
		SELECT PostID FROM comments WHERE ID = $1
	`

	if err := r.db.QueryRow(tempQuery, reaction.CommentID).Scan(&postID); err != nil {
		return postID, err
	}
	if err := r.db.QueryRow(selectQuery, reaction.UserID, reaction.CommentID).Scan(&vote); err != nil {
		if err != sql.ErrNoRows {
			return postID, err
		}
		if _, err := r.db.Exec(insertQuery, reaction.UserID, reaction.CommentID, reaction.Vote); err != nil {
			return postID, err
		}
	} else {
		if _, err := r.db.Exec(deleteQuery, reaction.UserID, reaction.CommentID, vote); err != nil {
			return postID, err
		}
		if vote != reaction.Vote {
			if _, err := r.db.Exec(insertQuery, reaction.UserID, reaction.CommentID, reaction.Vote); err != nil {
				return postID, err
			}
		}
	}
	return postID, nil
}
