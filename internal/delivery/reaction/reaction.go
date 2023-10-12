package reaction

import (
	"fmt"
	modelsData "forum/internal/models/data"
	modelsUser "forum/internal/models/user"
	"forum/pkg"
	"net/http"
	"strconv"
	"strings"
)

func (h *ReactionHandler) CreatePostReaction(w http.ResponseWriter, r *http.Request) {
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

	if r.Method == http.MethodGet {
		pkg.Error(w, http.StatusNotFound, nil)
		return
	}
	if r.Method != http.MethodPost {
		pkg.Error(w, http.StatusMethodNotAllowed, nil)
		return
	}
	if err := r.ParseForm(); err != nil {
		pkg.Error(w, http.StatusInternalServerError, err)
		return
	}

	postID, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/posts/react/"))
	// fmt.Println(postID)
	if err != nil {
		pkg.Error(w, http.StatusBadRequest, err)
		return
	}

	reactions, ok := r.Form["react"]
	fmt.Println(reactions, "reactions")

	if !ok {
		pkg.Error(w, http.StatusBadRequest, nil)
		return
	}

	if err := h.reaction.CreatePostReaction(postID, user.ID, reactions[0]); err != nil {
		fmt.Println("OK")
		pkg.Error(w, http.StatusInternalServerError, err)
		return
	}

	path := fmt.Sprintf("/posts?id=%v", postID)

	http.Redirect(w, r, path, http.StatusSeeOther)
}

func (h *ReactionHandler) CreateCommentReaction(w http.ResponseWriter, r *http.Request) {
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

	if r.Method == http.MethodGet {
		pkg.Error(w, http.StatusNotFound, nil)
		return
	}
	if r.Method != http.MethodPost {
		pkg.Error(w, http.StatusMethodNotAllowed, nil)
		return
	}
	if err := r.ParseForm(); err != nil {
		pkg.Error(w, http.StatusInternalServerError, err)
		return
	}

	reaction, ok1 := r.Form["react"]
	comment, ok2 := r.Form["commentID"]
	fmt.Println(reaction)
	fmt.Println(comment)

	if !ok1 || !ok2 {
		pkg.Error(w, http.StatusBadRequest, nil)
		return
	}

	commentID, err := strconv.Atoi(comment[0])
	fmt.Println("strconv", commentID)
	if err != nil {
		pkg.Error(w, http.StatusBadRequest, err)
		return
	}

	postID, err := h.reaction.CreateCommentReaction(commentID, user.ID, reaction[0])
	if err != nil {
		fmt.Println("creating")
		pkg.Error(w, http.StatusInternalServerError, err)
		return
	}

	path := fmt.Sprintf("/posts?id=%v", postID)

	http.Redirect(w, r, path, http.StatusSeeOther)
}
