package post

import (
	"database/sql"
	"errors"
	"fmt"
	p "forum/internal/models/post"
	"log"
	"strings"
)

const count = `
SELECT COUNT(*), (
	SELECT COUNT(*) FROM reactions WHERE VOTE=-1 AND PostID = $1
), (
	SELECT COUNT(*) FROM comments WHERE PostID = $1
)
FROM reactions WHERE VOTE=1 AND PostID = $1
`

func (r *PostSql) CreatePost(post p.Post) error {
	tag := strings.Join(post.Tags, ", ")
	query := `INSERT INTO posts (AuthorID,Title,Tag,Body,Author,Image_hash) VALUES ($1, $2, $3, $4,$5,$6)`
	res, err := r.db.Exec(query, &post.AuthorID, &post.Title, &tag, &post.Body, &post.Author, &post.Image)
	if err != nil {
		fmt.Println("OWIBKA ERROR")
		return err
	}
	id, _ := res.LastInsertId()
	fmt.Println(id)

	for _, tag := range post.Tags {
		query := `INSERT INTO tags (PostID,Tag) VALUES ($1, $2)`
		if _, err := r.db.Exec(query, id, tag); err != nil {
			log.Println(err)
			return err
		}
	}
	queryImage := `
			INSERT INTO images (PostID, Image_hash,File_type) VALUES ($1, $2,$3)
		`
	if _, err := r.db.Exec(queryImage, id, post.Image, post.FileType); err != nil {
		return err
	}
	return nil
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (r *PostSql) GetAllPosts(UserID int) ([]p.Post, error) {
	var posts []p.Post
	query := `
		SELECT posts.ID, posts.AuthorID, posts.Title,posts.Tag,posts.Body,users.Username
		FROM posts	INNER JOIN users ON users.ID = posts.AuthorID
		ORDER BY posts.ID DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var post p.Post
		var tag string
		if err := rows.Scan(&post.ID, &post.AuthorID, &post.Title, &tag, &post.Body, &post.Author); err != nil {
			fmt.Println("OK")
			return posts, err
		}

		if err := r.db.QueryRow(count, &post.ID).Scan(&post.Likecount, &post.Dislikecount, &post.Commentcount); err != nil {
			return posts, err
		}

		post.Tags = append(post.Tags, tag)
		if err != nil {
			return nil, err
		}

		vote, err := r.getReactionPost(UserID, post.ID)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return posts, err
			}
			vote = 0
		}
		post.Vote = vote

		posts = append(posts, post)
		fmt.Println(posts)
	}
	return posts, nil
}

func (r *PostSql) GetPostsByTag(userID int, tags []string) ([]p.Post, error) {
	var posts []p.Post

	getPostID, err := r.getPostTags(tags)
	if err != nil {
		return nil, err
	}
	// query := `
	// 	SELECT posts.ID, posts.AuthorID, posts.Title,posts.Tag,posts.Body,users.Username
	// 	FROM posts	INNER JOIN users ON users.ID = posts.AuthorID
	// 	ORDER BY posts.ID DESC
	// `

	query := `SELECT posts.ID,posts.AuthorID,posts.Title,posts.Tag,posts.Body,posts.Author FROM posts WHERE ID = ?`
	// queryUserID := `
	// 	SELECT posts.ID,users.Username
	// 	FROM posts INNER JOIN users ON users.ID = posts.AuthorID
	// `
	for _, val := range getPostID {
		row, err := r.db.Query(query, val)
		if err != nil {
			return nil, err
		}
		for row.Next() {
			var post p.Post
			var tag string
			if err := row.Scan(&post.ID, &post.AuthorID, &post.Title, &tag, &post.Body, &post.Author); err != nil {
				return posts, err
			}

			if err := r.db.QueryRow(count, &post.ID).Scan(&post.Likecount, &post.Dislikecount, &post.Commentcount); err != nil {
				return posts, err
			}

			vote, err := r.getReactionPost(userID, post.ID)
			if err != nil {
				if !errors.Is(err, sql.ErrNoRows) {
					return posts, err
				}
				vote = 0
			}
			// fmt.Println(posts,"posts")
			post.Vote = vote
			if validID(post, posts) {
				post.Tags = strings.Split(tag, " ")
				posts = append(posts, post)
			}

		}
	}

	return posts, nil
}

func (r *PostSql) GetMyPosts(userID int) ([]p.Post, error) {
	var posts []p.Post
	query := `
		SELECT posts.ID,posts.AuthorID,posts.Title,posts.Tag,posts.Body,posts.Author FROM posts WHERE AuthorID = ?
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var tag string
		var post p.Post
		if err := rows.Scan(&post.ID, &post.AuthorID, &post.Title, &tag, &post.Body, &post.Author); err != nil {
			return posts, err
		}
		if err := r.db.QueryRow(count, &post.ID).Scan(&post.Likecount, &post.Dislikecount, &post.Commentcount); err != nil {
			return posts, err
		}
		post.Tags = append(post.Tags, tag)
		if err != nil {
			return nil, err
		}

		vote, err := r.getReactionPost(userID, post.ID)
		if err != nil {
			return nil, err
		}
		post.Vote = vote
		posts = append(posts, post)

	}
	return posts, nil
}

func (r *PostSql) GetLikedPosts(postID int) ([]p.Post, error) {
	var posts []p.Post
	query := `
		SELECT posts.ID,posts.AuthorID,posts.Title,posts.Tag,posts.Body,users.Username 
		FROM posts INNER JOIN users ON users.ID = posts.AuthorID,
		reactions WHERE reactions.PostID = posts.ID AND reactions.VOTE = 1 AND reactions.UserID = $1
		ORDER BY posts.ID DESC
	`
	rows, err := r.db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tag string
		var post p.Post
		if err := rows.Scan(&post.ID, &post.AuthorID, &post.Title, &tag, &post.Body, &post.Author); err != nil {
			return posts, err
		}

		if err := r.db.QueryRow(count, &post.ID).Scan(&post.Likecount, &post.Dislikecount, &post.Commentcount); err != nil {
			return posts, err
		}

		post.Tags = append(post.Tags, tag)
		post.Vote = 1
		posts = append(posts, post)

	}
	return posts, nil
}

func (r *PostSql) GetPostById(postID, UserID int) (p.Post, error) {
	var tag string
	var post p.Post
	query := `
		SELECT posts.ID, posts.AuthorID, posts.Title, posts.Tag,posts.Body, users.Username
		FROM posts INNER JOIN users ON users.ID = Posts.AuthorID
		WHERE posts.ID = $1
	`
	if err := r.db.QueryRow(query, postID).Scan(&post.ID, &post.AuthorID, &post.Title, &tag, &post.Body, &post.Author); err != nil {
		return post, err
	}

	if err := r.db.QueryRow(count, &postID).Scan(&post.Likecount, &post.Dislikecount, &post.Commentcount); err != nil {
		return post, err
	}

	post.Tags = strings.Split(tag, " ")
	post.Tags = append(post.Tags, tag)

	vote, err := r.getReactionPost(UserID, post.ID)
	fmt.Println(err)
	if err != nil {
		return post, err
	}
	image, fileType, err := r.getPostImages(post.ID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return post, err
		}
	}
	post.Image = image
	post.FileType = fileType
	post.Vote = vote
	// fmt.Println("OK")
	return post, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (r *PostSql) getPostTags(tags []string) ([]int, error) {
	var res []int
	query := `
		SELECT PostID FROM tags WHERE Tag = ?
	`
	for _, val := range tags {
		row, err := r.db.Query(query, val)
		if err != nil {
			return nil, err
		}
		for row.Next() {
			var id int
			if err := row.Scan(&id); err != nil {
				return nil, err
			}
			res = append(res, id)

		}
	}
	// fmt.Println(res)
	return res, nil
}

func (r *PostSql) getReactionPost(userID, postID int) (int, error) {
	query := `
		SELECT VOTE FROM reactions WHERE UserID = $1 AND PostID = $2	
	`
	var res int
	if err := r.db.QueryRow(query, userID, postID).Scan(&res); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return res, err
	}

	return res, nil
}

func validID(post p.Post, posts []p.Post) bool {
	// fmt.Println(post)
	// fmt.Println(posts)
	for _, v := range posts {
		if v.ID == post.ID {
			return false
		}
	}
	return true
}

func (r *PostSql) getPostImages(postID int) (string, string, error) {
	const query = `
		SELECT Image_hash,File_type FROM images WHERE PostID = $1
	`
	rows, err := r.db.Query(query, postID)
	if err != nil {
		return "", "", err
	}

	defer rows.Close()

	var images string
	var fileTypes string
	for rows.Next() {
		var image string
		var fileType string
		if err := rows.Scan(&image, &fileType); err != nil {
			return images, fileTypes, err
		}
		images += image
		fileTypes += fileType
	}

	return images, fileTypes, err
}

func (r *PostSql) DeletePostById(postId int) error {
	query := `DELETE FROM posts where ID=?`
	if _, err := r.db.Exec(query, postId); err != nil {
		return fmt.Errorf("can't set post: %w", err)
	}
	return nil
}

func (r *PostSql) DeleteTagsToPost(postId int) error {
	query := `DELETE FROM tags WHERE PostID=?`
	if _, err := r.db.Exec(query, postId); err != nil {
		return fmt.Errorf("can't set post: %w", err)
	}
	return nil
}
