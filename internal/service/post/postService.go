package post

import (
	modelsPost "forum/internal/models/post"
	repositoryPost "forum/internal/repository/post"
)

type Post interface {
	CreatePost(post modelsPost.Post) error
	AllPostList(userID int) ([]modelsPost.Post, error)
	MyPosts(userID int) ([]modelsPost.Post, error)
	PostsByTag(userID int, tags []string) ([]modelsPost.Post, error)
	PostByID(postID, userID int) (modelsPost.Post, error)
	MyLikedPosts(userID int) ([]modelsPost.Post, error)
	DeletePostById(postId int) error
	DeleteTagsToPost(postId int) error
}

type PostService struct {
	repo repositoryPost.Post
}

func NewPostService(repo repositoryPost.Post) *PostService {
	return &PostService{repo: repo}
}
