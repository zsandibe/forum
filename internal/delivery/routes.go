package delivery

import (
	"net/http"
)

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(http.Dir("./ui/"))))
	mux.HandleFunc("/", h.Middleware(h.IndexPage))
	mux.HandleFunc("/auth/sign-up", h.Auth.SignUp)
	mux.HandleFunc("/auth/sign-in", h.Auth.SignIn)
	mux.HandleFunc("/auth/log-out", h.Middleware(h.Auth.LogOut))
	mux.HandleFunc("/posts/create", h.Middleware(h.Post.CreatePost))
	// //----------------------------------------
	mux.HandleFunc("/posts", h.Middleware(h.Post.PostPage))
	// mux.HandleFunc("/post/tag", H.Middleware(h.Post.PostsByTag))
	mux.HandleFunc("/myposts", h.Middleware(h.Post.MyPosts))
	mux.HandleFunc("/mylikes", h.Middleware(h.Post.MyLikedPosts))
	// //----------------------------------------
	mux.HandleFunc("/posts/react/", h.Middleware(h.Reaction.CreatePostReaction))
	mux.HandleFunc("/comment/react/", h.Middleware(h.Reaction.CreateCommentReaction))
	return h.logRequest(mux)
}
