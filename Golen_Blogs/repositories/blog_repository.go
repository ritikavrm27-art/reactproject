package repositories

import (
	"blogapi/config"
	"blogapi/models"
)

func GetBlogs() ([]models.Blog, error) {

	rows, err := config.DB.Query("SELECT id, title, content, author FROM public.tblogs")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	blogs := []models.Blog{}

	for rows.Next() {
		var blog models.Blog

		err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author)
		if err != nil {
			return nil, err
		}

		blogs = append(blogs, blog)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}