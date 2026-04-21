package models

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"uname"`
	Email string `json:"uemail"`
	Password string `json:"upassword"`
}
// JSON response structure
type Response struct {
	Message  string `json:"message"`
	FileName string `json:"fileName"`
}
type UserPage struct {
	Message string
	Users   []User
}