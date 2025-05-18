package db

import (
	"errors"

	"github.com/bruno-holanda15/api-rest-challenge-rocketseat/app/entities"
	"github.com/google/uuid"
)

type ID uuid.UUID

var (
	ErrUserNotFound = errors.New("user not found")
)

func (id ID) String() string {
	return uuid.UUID(id).String()
}

type AppStorage struct {
	Data map[ID]entities.User
}

func NewAppStorage() *AppStorage {
	data := make(map[ID]entities.User)
	return &AppStorage{Data: data}
}

func (d *AppStorage) Insert(user entities.User) ID {
	newID := uuid.New()
	d.Data[ID(newID)] = user

	return ID(newID)
}

func (d *AppStorage) FindById(id ID) (entities.User, error) {
	user, ok := d.Data[id]
	if ok {
		return user, nil
	}

	return entities.User{}, ErrUserNotFound
}
