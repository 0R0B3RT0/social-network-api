package repositories

import (
	"api/src/models"
	"database/sql"
)

type Publications struct {
	db *sql.DB
}

func NewPublicationRepositories(db *sql.DB) *Publications {
	return &Publications{db}
}

//Create persist a new publication
func (repository Publications) Create(pub models.Publication) (uint64, error) {
	statement, err := repository.db.Prepare("insert into publications(title, content, user_id) values(?,?,?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(pub.Title, pub.Content, pub.UserID)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil
}

//Find find a specific publication by id
func (repository Publications) Find(ID uint64) (publication models.Publication, err error) {
	rows, err := repository.db.Query("select p.id, p.title, p.content, p.likes, p.user_id, u.nick, p.created_at from publications p join users u on p.user_id = u.id where p.id = ?", ID)

	if err != nil {
		return
	}

	if rows.Next() {
		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.Likes,
			&publication.UserID,
			&publication.UserNick,
			&publication.CreatedAt,
		); err != nil {
			return
		}
	}

	return
}
