package post

import (
	"errors"
	"fmt"
	"forum/pkg"
	"html/template"
	"net/http"
	"strconv"

	modelsComment "forum/internal/models/comment"
	modelsData "forum/internal/models/data"
	modelsPost "forum/internal/models/post"
	modelsUser "forum/internal/models/user"
)

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("KIRDI: Creating")
	var tmplData modelsData.Data
	userValue := r.Context().Value(modelsData.ContextKeyUser)
	if userValue == nil {
		fmt.Println("ERROR : unauthorized user")
		// Обработка случая, когда пользователь не аутентифицирован
		pkg.Error(w, http.StatusUnauthorized, nil)
		return
	}

	user, ok := userValue.(modelsUser.User)
	if !ok {
		// Обработка случая, когда значение в контексте не является типом User
		fmt.Println("ERROR : invalid user data in context")
		pkg.Error(w, http.StatusInternalServerError,nil)
		return
	}

	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./ui/templates/create_post.html")
		if err != nil {
			pkg.Error(w, http.StatusInternalServerError,err)
			return
		}
		tmpl.Execute(w, tmplData)
	} else if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Println("ERROR : cannot parse form")
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		title,ok1 := r.Form["title"]
		content,ok2:= r.Form["content"]
		tags, ok3 := r.Form["tag"]
		fmt.Println(tags)
		// fmt.Println(title)
		// fmt.Println(content)
		// fmt.Println(tags)
		if !ok1 || !ok2 || !ok3 {
			pkg.Error(w, http.StatusBadRequest, nil)
			return
		}
		post := modelsPost.Post{
			Title:    title[0],
			AuthorID: user.ID,
			Body:     content[0],
			Tags:     tags,
			Author: user.Username,
		}
		if err := h.post.CreatePost(post); err != nil {
			if errors.Is(err, errors.New("ErrEmptyBody")) {
				pkg.Error(w, http.StatusBadRequest, err)
				return
			}
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		pkg.Error(w, http.StatusMethodNotAllowed, nil)
	}
}

func (h *PostHandler) PostPage(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("kirdi: post page")
	userValue := r.Context().Value(modelsData.ContextKeyUser)
	if userValue == nil {
		pkg.Error(w, http.StatusUnauthorized, nil)
		return
	}
	user, ok := userValue.(modelsUser.User)
	if !ok {
		pkg.Error(w, http.StatusInternalServerError,nil)
		return
	}
	// fmt.Println(r.URL.Query().Get("id"))
	postID, err := strconv.Atoi(r.URL.Query().Get("id"))
	// fmt.Println(postID)
	if err != nil {
		fmt.Println("owibka postID          ", err)
		pkg.Error(w, http.StatusNotFound, err)
		return
	}

	switch r.Method {
	case http.MethodGet:
		post, err := h.post.PostByID(postID, user.ID)
		if err != nil {
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		comments, err := h.comment.GetCommentsByPostID(postID, user.ID)
		// fmt.Println(comments, "comments")
		data := modelsData.Data{
			User:     user,
			Post:     post,
			Comments: comments,
		}
		tmpl, err := template.ParseFiles("./ui/templates/post.html")
		if err != nil {
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		tmpl.Execute(w, data)

	case http.MethodPost:
		if user == (modelsUser.User{}) {
			pkg.Error(w, http.StatusUnauthorized, nil)
			return
		}
		if err := r.ParseForm(); err != nil {
			// fmt.Println("ERROR : cannot parse form")
			pkg.Error(w, http.StatusInternalServerError,err)
			return
		}
		commentBody, ok1 := r.Form["comment"]
		if !ok1 {
			pkg.Error(w, http.StatusBadRequest, nil)
			return
		}
		comment := modelsComment.Comment{
			UserId: user.ID,
			PostId: postID,
			Body:   commentBody[0],
		}
		if err := h.comment.CreateComment(comment); err != nil {
			if errors.Is(err, errors.New("ErrEmptyBody")) {
				pkg.Error(w, http.StatusBadRequest, err)
				return
			}
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		path := fmt.Sprintf("/posts?id=%v", postID)
		http.Redirect(w, r, path, http.StatusSeeOther)
	default:
		pkg.Error(w, http.StatusMethodNotAllowed, nil)
	}
}

func (h *PostHandler) MyPosts(w http.ResponseWriter, r *http.Request) {
	userValue := r.Context().Value(modelsData.ContextKeyUser)
	if userValue == nil {
		pkg.Error(w, http.StatusUnauthorized, nil)
		return
	}
	user, ok := userValue.(modelsUser.User)
	if !ok {
		pkg.Error(w,http.StatusInternalServerError,nil)
		return
	}

	switch r.Method {
	case http.MethodGet:
		posts, err := h.post.MyPosts(user.ID)
		if err != nil {
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		data := modelsData.Data{
			User:  user,
			Posts: posts,
		}
		tmpl, err := template.ParseFiles("./ui/templates/myposts.html")
		if err != nil {
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		tmpl.Execute(w, data)
	case http.MethodPost:
		if user == (modelsUser.User{}) {
			pkg.Error(w, http.StatusUnauthorized, nil)
			return
		}
		if err := r.ParseForm(); err != nil {
			// fmt.Println("ERROR : cannot parse form")
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
	default:
		pkg.Error(w, http.StatusMethodNotAllowed, nil)
	}
}

func (h *PostHandler) MyLikedPosts(w http.ResponseWriter, r *http.Request) {
	userValue := r.Context().Value(modelsData.ContextKeyUser)
	if userValue == nil {
		fmt.Println("ERROR : unauthorized user")
		// Обработка случая, когда пользователь не аутентифицирован
		pkg.Error(w,http.StatusUnauthorized, nil)
		return
	}

	user, ok := userValue.(modelsUser.User)
	if !ok {
		// Обработка случая, когда значение в контексте не является типом User
		fmt.Println("ERROR : invalid user data in context")
		pkg.Error(w, http.StatusInternalServerError,nil)
		return
	}

	switch r.Method {
	case http.MethodGet:
		posts, err := h.post.MyLikedPosts(user.ID)
		if err != nil {
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		data := modelsData.Data{
			User:  user,
			Posts: posts,
		}
		tmpl, err := template.ParseFiles("./ui/templates/myliked_posts.html")
		fmt.Println(err)
		if err != nil {
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		tmpl.Execute(w, data)
	case http.MethodPost:
		if user == (modelsUser.User{}) {
			pkg.Error(w, http.StatusUnauthorized,nil)
			return
		}
		if err := r.ParseForm(); err != nil {
			// fmt.Println("ERROR : cannot parse form")
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
	default:
		pkg.Error(w, http.StatusMethodNotAllowed, nil)
	}
}
