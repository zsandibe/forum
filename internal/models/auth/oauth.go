package auth

type OAuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type,omitempty"`
	Scope       string `json:"scope"`
}

type UserInfo struct {
	Email string `json:"email"`
}

// type GithubInfo struct {
// 	Email    string `json:"email"`
// 	Primary  bool   `json:"primary"`
// 	Verified bool   `json:"verified"`
// }
