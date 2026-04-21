package repositories

import (
	"userapi/config"
	"userapi/models"
)

func GetUsers() ([]models.User, error) {

	rows, err := config.DB.Query("SELECT id, uname, uemail, upassword FROM public.tusers")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}

	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}