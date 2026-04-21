package services

import (
	"blogapi/models"
	"blogapi/repositories"
	"blogapi/config"
)

func GetAllBlogs() ([]models.Blog, error) {
	return repositories.GetBlogs()
}
func InsertBlog(blog models.Blog) error {

	query := `INSERT INTO public.tblogs (title, content, author) VALUES ($1, $2, $3)`

	_, err := config.DB.Exec(query, blog.Title, blog.Content, blog.Author)
	return err
}
func DeleteBlog(blog models.Blog) error {

	query := `DELETE FROM public.tblogs WHERE ID = $1`

	_, err := config.DB.Exec(query, blog.ID)
	return err
}