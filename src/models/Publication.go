package models

import "time"

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
