package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User it is a user for the social network
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Pass      string    `json:"pass,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// PrepareToCreate validate and format the new user to persist
func (user *User) PrepareToCreate() error {
	if error := user.validate("create"); error != nil {
		return error
	}

	if error := user.formaterToCreate(); error != nil {
		return error
	}

	return nil
}

// PrepareToUpdate validate and format the new user to update
func (user *User) PrepareToUpdate() error {
	if error := user.validate("update"); error != nil {
		return error
	}
	user.formaterToUpdate()

	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("The name is required field")
	}
	if user.Nick == "" {
		return errors.New("The nick is required field")
	}
	if user.Email == "" {
		return errors.New("The email is required field")
	}
	if checkmail.ValidateFormat(user.Email) != nil {
		return errors.New("The email has invalid format")
	}
	if step == "create" && user.Pass == "" {
		return errors.New("The pass is required field")
	}

	return nil
}

func (user *User) formaterToCreate() error {
	user.formaterToUpdate()

	hash, error := security.Hash(user.Pass)
	if error != nil {
		return error
	}
	user.Pass = string(hash)

	return nil
}

func (user *User) formaterToUpdate() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
