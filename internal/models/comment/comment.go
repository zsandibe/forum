package comment

type Comment struct {
	Id      int
	UserId  int
	PostId  int
	Like    int
	Dislike int
	Vote    int
	Body    string
	Author  string
}
