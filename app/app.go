package app

import "github.com/google/uuid"

type ID uuid.UUID

type User struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Biography string `json:"biography,omitempty"`
}

type AppStorage struct {
	Data map[ID]User
}

func NewAppStorage() *AppStorage {
	data := make(map[ID]User)
	return &AppStorage{Data: data}
}

func (d *AppStorage) Insert(user User) {
	newID := uuid.New()
	d.Data[ID(newID)] = user
}