package user

type User struct {
	ID              int
	Username        string
	Email           string
	Password        string
	ConfirmPassword string
	AuthMethod      string
	Role            string
	Requested       bool
}
