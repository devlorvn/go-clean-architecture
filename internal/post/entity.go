package post

type Post struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID int64  `json:"author_id"`
}
