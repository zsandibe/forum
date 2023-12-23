package auth

import (
	"forum/internal/models/auth"
)

const (
	googleAuthURL       = "https://accounts.google.com/o/oauth2/auth"
	googleTokenURL      = "https://oauth2.googleapis.com/token"
	googleClientID      = "773702556277-p4vi75opdrbqf2rml4lot6ou5ldbmmm8.apps.googleusercontent.com"
	googleClientSecret  = "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"
	githubAuthURL       = "https://github.com/login/oauth/authorize"
	githubTokenURL      = "https://github.com/login/oauth/access_token"
	githubClientID1     = "8dd9c63d9466d1a3d6c2"
	githubClientSecret1 = "bc880768ac243983f29f566bf23df7ca8cd9dfc1"
	githubClientID2     = "29e8fe903e19ff75bd35"
	githubClientSecret2 = "96fa8311a1aa4a17b90788728390c3dca383c0cb"
)

var (
	googleSignInConfig = &auth.OAuthConfig{
		ClientID:     googleClientID,
		ClientSecret: googleClientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		RedirectURL: "https://localhost:8081/auth/sign-in/google/callback",
	}
	googleSignUpConfig = &auth.OAuthConfig{
		ClientID:     googleClientID,
		ClientSecret: googleClientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		RedirectURL: "https://localhost:8081/auth/sign-up/google/callback",
	}
	githubSignInConfig = &auth.OAuthConfig{
		ClientID:     githubClientID1,
		ClientSecret: githubClientSecret1,
		Scopes: []string{
			"user:email",
		},
		RedirectURL: "https://localhost:8081/auth/sign-in/github/callback",
	}
	githubSignUpConfig = &auth.OAuthConfig{
		ClientID:     githubClientID2,
		ClientSecret: githubClientSecret2,
		Scopes: []string{
			"user:email",
		},
		RedirectURL: "https://localhost:8081/auth/sign-up/github/callback",
	}
)
