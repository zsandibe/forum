package post

import "errors"

var (
	ErrEmptyBody    = errors.New("Can`t create post with empty body")
	ErrPostNotFound = errors.New("Post not found")
)
