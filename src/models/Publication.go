package models

import (
	"errors"
	"strings"
	"time"
)

//Publication it is a user's publish
type Publication struct {
	ID        uint64     `json:"id,omitempty"`
	Title     string     `json:"title,omitempty"`
	Content   string     `json:"content,omitempty"`
	UserID    uint64     `json:"userId,omitempty"`
	UserNick  string     `json:"userNick,omitempty"`
	Likes     uint64     `json:"likes"`
	CreatedAt *time.Time `json:"createdAd,omitempty"`
}

func (publication *Publication) Prepare() error {
	if err := publication.validate(); err != nil {
		return err
	}

	publication.format()

	return nil
}

func (publication *Publication) validate() error {
	if strings.ReplaceAll(publication.Title, " ", "") == "" {
		return errors.New("the tittle is required field")
	}
	if strings.ReplaceAll(publication.Content, " ", "") == "" {
		return errors.New("the content is required field")
	}

	return nil
}

func (publication *Publication) format() {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)
}
