package post

import (
	"database/sql"
	p "forum/internal/models/post"
)

type PostSql struct {
	db *sql.DB
}

func NewPostSql(db *sql.DB) *PostSql {
	return &PostSql{
		db: db,
	}
}

type Post interface {
	CreatePost(post p.Post) error
	GetPostById(postID, userID int) (p.Post, error)
	GetAllPosts(userID int) ([]p.Post, error)
	GetMyPosts(userID int) ([]p.Post, error)
	GetPostsByTag(userID int, tags []string) ([]p.Post, error)
	GetLikedPosts(userID int) ([]p.Post, error)
}
