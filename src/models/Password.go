package models

type Password struct {
	LastPassword string `json:"last_password"`
	NewPassword  string `json:"new_password"`
}
