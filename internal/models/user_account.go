package models

type UserAccount struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
