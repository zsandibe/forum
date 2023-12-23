package post

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	modelsPost "forum/internal/models/post"
)

func (s *PostService) CreatePost(post modelsPost.Post) error {
	fmt.Println("service post create post")
	if strings.TrimSpace(post.Body) == "" {
		return ErrEmptyBody
	}
	// fmt.Println("OK")
	return s.repo.CreatePost(post)
}

func (s *PostService) PostByID(postID, userID int) (modelsPost.Post, error) {
	post, err := s.repo.GetPostById(postID, userID)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("service post")
		return modelsPost.Post{}, ErrPostNotFound
	} else if err != nil {
		return post, err
	}
	// post.ImagePath = fmt.Sprintf("/images/%s/%s.%s",
	// 	post.FileType, post.FileName, post.FileType)
	// post.Image = fmt.Sprintf("images/%s/%s.%s", post.FileType, post.Image, post.FileType)
	fmt.Println(post.Image)
	return post, nil
}

func (s *PostService) MyPosts(userID int) ([]modelsPost.Post, error) {
	return s.repo.GetMyPosts(userID)
}

func (s *PostService) AllPostList(userID int) ([]modelsPost.Post, error) {
	return s.repo.GetAllPosts(userID)
}

func (s *PostService) PostsByTag(userID int, tags []string) ([]modelsPost.Post, error) {
	if err := tagCheck(tags); err != nil {
		return nil, err
	}
	return s.repo.GetPostsByTag(userID, tags)
}

func (s *PostService) DeletePostById(postId int) error {
	if err := s.repo.DeleteTagsToPost(postId); err != nil {
		return fmt.Errorf("can`t delete tags: %w", err)
	}
	return s.repo.DeletePostById(postId)
}

func (s *PostService) MyLikedPosts(userID int) ([]modelsPost.Post, error) {
	return s.repo.GetLikedPosts(userID)
}

func tagCheck(tags []string) error {
	if len(tags) == 0 {
		return fmt.Errorf("empty tag")
	}

	tag := []string{"Action", "Fantasy", "Adventure", "Horror", "Thriller"}
	for _, v := range tag {
		for _, j := range tags {
			if v == j {
				return nil
			}
		}
	}
	return fmt.Errorf("you cannot select another category")
}

func (s *PostService) DeleteTagsToPost(postId int) error {
	return s.repo.DeleteTagsToPost(postId)
}
