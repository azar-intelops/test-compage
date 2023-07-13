package models

type User struct {
	Id int64 `json:"id,omitempty"`

	Username `json:"username,omitempty"`
}
