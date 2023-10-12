package post

type Post struct {
	ID           int
	AuthorID     int
	Likecount    int
	Dislikecount int
	Commentcount int
	Vote         int
	
	// Created      time.Time
	// Time         string
	Author string
	Title  string
	Body   string
	Tags   []string
}
