package comment

import (
	"database/sql"
	"errors"
	modelsComment "forum/internal/models/comment"
)

func (r *CommentSql) CreateComment(comment modelsComment.Comment) error {
	query := `
		INSERT INTO comments(AuthorId, PostId, Content) VALUES ($1, $2, $3);
	`

	if _, err := r.db.Exec(query, comment.UserId, comment.PostId, comment.Body); err != nil {
		return err
	}
	return nil
}

func (r *CommentSql) GetCommentsByPostID(postId, userID int) ([]modelsComment.Comment, error) {
	var comments []modelsComment.Comment
	query := `
        SELECT comments.ID,comments.AuthorId,comments.PostId,comments.Content,users.Username
		FROM comments INNER JOIN users ON users.ID = comments.AuthorId
		WHERE comments.PostId = $1;
    `

	count := `
		SELECT COUNT(*), (
			SELECT COUNT(*) FROM REACTIONS WHERE VOTE=-1 AND CommentID = $1
		)
		FROM REACTIONS WHERE VOTE=1 AND CommentID = $1
	`

	rows, err := r.db.Query(query, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment modelsComment.Comment
		if err := rows.Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Body, &comment.Author); err != nil {
			return comments, err
		}

		if err := r.db.QueryRow(count, comment.Id).Scan(&comment.Like, &comment.Dislike); err != nil {
			return comments, err
		}
		vote, err := r.getReactionForComment(userID, comment.Id)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return comments, nil
			}
			vote = 0
		}
		comment.Vote = vote

		comments = append(comments, comment)
	}
	return comments, nil
}

func (s *CommentSql) getReactionForComment(userID, commentID int) (int, error) {
	var vote int
	query := `
        SELECT VOTE FROM reactions WHERE UserID = $1 AND CommentID = $2;
    `
	if err := s.db.QueryRow(query, userID, commentID).Scan(&vote); err != nil {
		return vote, err
	}
	return vote, nil
}
