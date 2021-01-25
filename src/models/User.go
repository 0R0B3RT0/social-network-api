package modelos

import "time"

type Usuario struct {
	ID        uint64    `json:id,omitempty`
	Name      string    `json:name,omitempty`
	Nick      string    `json:nick,omitempty`
	Email     string    `json:email,omitempty`
	Pass      string    `json:pass,omitempty`
	CreatedAt time.Time `json:createdAt,omitempty`
}
