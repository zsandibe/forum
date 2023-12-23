package post

import "errors"

var (
	ErrEmptyBody    = errors.New("Can`t create post with empty body")
	ErrPostNotFound = errors.New("Post not found")
)

var (
	ErrImgSize   = errors.New("image file size is too big")
	ErrImgFormat = errors.New("your image format is not provided")
)
