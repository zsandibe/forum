package post

import (
	modelsPost "forum/internal/models/post"
	modelsUser "forum/internal/models/user"
	modelsComment "forum/internal/models/comment"
)

type postProvider interface {
	CreatePost(post modelsPost.Post) error
	AllPostList(userID int) ([]modelsPost.Post, error)
	MyPosts(userID int) ([]modelsPost.Post, error)
	PostsByTag(userID int, tags []string) ([]modelsPost.Post, error)
	PostByID(postID, userID int) (modelsPost.Post, error)
	MyLikedPosts(userID int) ([]modelsPost.Post, error)
}

type userProvider interface {
	GetUserByID(userID int) (modelsUser.User, error)
}

type commentProvider interface{
	CreateComment(comment modelsComment.Comment) error
	GetCommentsByPostID(postID,userID int)	([]modelsComment.Comment, error)
}

type PostHandler struct {
	post    postProvider
	user    userProvider
	comment commentProvider
}

func NewPostHandler(post postProvider, user userProvider, comment commentProvider) *PostHandler {
	return &PostHandler{
		post:    post,
		user:    user,
		comment: comment,
	}
}
