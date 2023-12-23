package data

import (
	modelsComment "forum/internal/models/comment"
	modelsPost "forum/internal/models/post"
	modelsReport "forum/internal/models/report"
	modelsRequest "forum/internal/models/request"
	modelsUser "forum/internal/models/user"
)

type contextKey string

const ContextKeyUser contextKey = "user"

type Data struct {
	UserID   int
	User     modelsUser.User
	Users    []modelsUser.User
	Post     modelsPost.Post
	Posts    []modelsPost.Post
	Comments []modelsComment.Comment
	Request  modelsRequest.Request
	Requests []modelsRequest.Request
	Report   modelsReport.Report
	Reports  []modelsReport.Report
	Error    ErrorMsg
}

type ErrorMsg struct {
	Status  int
	Message string
}
