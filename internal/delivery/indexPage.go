package delivery

import (
	"fmt"
	modelsData "forum/internal/models/data"
	modelsPost "forum/internal/models/post"
	modelsUser "forum/internal/models/user"
	"forum/pkg"
	"net/http"
	"strconv"
	"text/template"
)

func (h *Handler) IndexPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		pkg.Error(w, http.StatusNotFound, nil)
		return
	}

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
		pkg.Error(w,http.StatusInternalServerError,nil)
		return
	}
	switch r.Method {
	case http.MethodGet:
		var posts []modelsPost.Post
		if len(r.URL.Query()) == 0 {
			var err error
			posts, err = h.service.Post.AllPostList(user.ID)
			if err != nil {
				pkg.Error(w, http.StatusInternalServerError, err)
				return
			}		
		} else {
			var err error
			values := r.URL.Query()
			// fmt.Println(values)
			tags := values["tag"]
			// fmt.Println(tags)
			posts, err = h.service.Post.PostsByTag(user.ID, tags)
			if err != nil {
				fmt.Println("OK")
				pkg.Error(w, http.StatusInternalServerError, err)
				return
			}
		}

		data := modelsData.Data{
			User:  user,
			Posts: posts,
		}
		tmpl, err := template.ParseFiles("./ui/templates/index.html")
		if err != nil {
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}
		tmpl.Execute(w, data)
	case http.MethodPost:
		if user == (modelsUser.User{}) {
			pkg.Error(w,http.StatusUnauthorized, nil)
			return
		}

		if err := r.ParseForm(); err != nil {
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}

		postID, ok1 := r.Form["postID"]
		react, ok2 := r.Form["react"]

		if !ok1 || !ok2 {
			pkg.Error(w, http.StatusBadRequest, nil)
			return
		}

		id, err := strconv.Atoi(postID[0])
		if err != nil {
			pkg.Error(w, http.StatusBadRequest, err)
			return
		}

		if err := h.service.Reaction.CreatePostReaction(id, user.ID, react[0]); err != nil {
			pkg.Error(w, http.StatusInternalServerError, err)
			return
		}

		http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
	default:
		pkg.Error(w, http.StatusMethodNotAllowed, nil)
	}
}
