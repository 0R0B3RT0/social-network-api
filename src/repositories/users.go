package repositories

import (
	"api/src/models"
	"database/sql"
)

// Users it is a users's repository
type Users struct {
	db *sql.DB
}

// NewUserRepositories create new users's repository
func NewUserRepositories(db *sql.DB) *Users {
	return &Users{db}
}

// Create create a new user at the database
func (repository Users) Create(user models.User) (uint64, error) {
	statement, error := repository.db.Prepare("insert into users (name, nick, email, pass) values(?, ?, ?, ?)")
	if error != nil {
		return 0, nil
	}
	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Pass)
	if error != nil {
		return 0, error
	}

	lastID, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(lastID), nil
}
