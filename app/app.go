package app

import (
	"errors"

	"github.com/google/uuid"
)

type ID uuid.UUID

var (
	ErrUserNotFound = errors.New("user not found")
)

func (id ID) String() string {
	return uuid.UUID(id).String()
}

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

func (d *AppStorage) Insert(user User) ID {
	newID := uuid.New()
	d.Data[ID(newID)] = user

	return ID(newID)
}

func (d *AppStorage) FindById(id ID) (User, error) {
	user, ok := d.Data[id]
	if ok {
		return user, nil
	}
	
	return User{}, ErrUserNotFound
}