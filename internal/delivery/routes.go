package delivery

import (
	"net/http"
)

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(http.Dir("./ui/"))))
	mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images/"))))
	mux.HandleFunc("/", h.Middleware(h.IndexPage))
	//-------------------------------------------
	mux.HandleFunc("/auth/sign-up", h.Auth.SignUp)
	mux.HandleFunc("/auth/sign-up/google", h.Auth.GoogleSignUp)
	mux.HandleFunc("/auth/sign-up/google/callback", h.Auth.SignUpCallbackFromGoogle)
	mux.HandleFunc("/auth/sign-up/github", h.Auth.GithubSignUp)
	mux.HandleFunc("/auth/sign-up/github/callback", h.Auth.SignUpCallbackFromGithub)
	// ------------------------------------------
	mux.HandleFunc("/auth/sign-in", h.Auth.SignIn)
	mux.HandleFunc("/auth/sign-in/google", h.Auth.GoogleSignIn)
	mux.HandleFunc("/auth/sign-in/google/callback", h.Auth.SignInCallbackFromGoogle)
	mux.HandleFunc("/auth/sign-in/github", h.Auth.GithubSignIn)
	mux.HandleFunc("/auth/sign-in/github/callback", h.Auth.SignInCallbackFromGithub)
	//------------------------------------------
	mux.HandleFunc("/auth/log-out", h.Middleware(h.Auth.LogOut))
	// -----------------------------------------
	mux.HandleFunc("/posts/create", h.Middleware(h.Post.CreatePost))
	mux.HandleFunc("/posts", h.Middleware(h.Post.PostPage))
	mux.HandleFunc("/posts/delete", h.Middleware(h.Post.DeletePost))
	// mux.HandleFunc("/post/tag", H.Middleware(h.Post.PostsByTag))
	mux.HandleFunc("/myposts", h.Middleware(h.Post.MyPosts))
	mux.HandleFunc("/mylikes", h.Middleware(h.Post.MyLikedPosts))
	// -----------------------------------------
	mux.HandleFunc("/request/create", h.Middleware(h.Request.CreateRequest))
	mux.HandleFunc("/my-request", h.Middleware(h.Request.MyRequest))
	// -----------------------------------------
	mux.HandleFunc("/posts/react/", h.Middleware(h.Reaction.CreatePostReaction))
	mux.HandleFunc("/comment/react/", h.Middleware(h.Reaction.CreateCommentReaction))
	// ---------------ADMIN---------------------
	mux.HandleFunc("/users", h.Middleware(h.Auth.GetAllUsersList))
	mux.HandleFunc("/admin/requests", h.Middleware(h.Request.GetAllRequests))
	mux.HandleFunc("/admin/requests/change", h.Middleware(h.Request.UpdateRequestStatus))
	mux.HandleFunc("/users/type/change", h.Middleware(h.Request.UpdateUserRole))
	mux.HandleFunc("/admin/reports", h.Middleware(h.Report.GetAllReports))
	mux.HandleFunc("/admin/reports/change", h.Middleware(h.Report.UpdateReportStatus))
	mux.HandleFunc("/my-reports", h.Middleware(h.Report.GetReportByUserID))
	// ---------------Moderator--------------------
	mux.HandleFunc("/moderator/report/create", h.Middleware(h.Report.CreateReport))

	// mux.HandleFunc("/request-moderator", h.Middleware())
	return h.logRequest(mux)
}
