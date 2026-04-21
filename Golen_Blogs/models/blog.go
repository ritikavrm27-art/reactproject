package models

type Blog struct {
	ID    int    `json:"id"`
	Title  string `json:"title"`
	Content string `json:"content"`
	Author string `json:"author"`
}
// JSON response structure
type Response struct {
	Message  string `json:"message"`
	FileName string `json:"fileName"`
}
type UserPage struct {
	Message string
	Blogs   []Blog
}