package services

import (
	"userapi/models"
	"userapi/repositories"
	"userapi/config"
)

func GetAllUsers() ([]models.User, error) {
	return repositories.GetUsers()
}
func InsertUser(user models.User) error {

	query := `INSERT INTO public.tusers (uname, uemail, upassword) VALUES ($1, $2, $3)`

	_, err := config.DB.Exec(query, user.Name, user.Email, user.Password)
	return err
}
func DeleteUser(user models.User) error {

	query := `DELETE FROM public.tusers WHERE ID = $1`

	_, err := config.DB.Exec(query, user.ID)
	return err
}