package post

type Post struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID int    `json:"author_id"`
}
