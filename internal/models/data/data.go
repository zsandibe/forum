package data

import (
	modelsComment "forum/internal/models/comment"
	modelsPost "forum/internal/models/post"
	modelsUser "forum/internal/models/user"
)

type contextKey string

const ContextKeyUser contextKey = "user"

type Data struct {
	UserID   int
	User     modelsUser.User
	Post     modelsPost.Post
	Posts    []modelsPost.Post
	Comments []modelsComment.Comment
	Error ErrorMsg
}


type ErrorMsg struct {
	Status int
	Message string
}