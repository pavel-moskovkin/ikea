package manager

import (
	"localshop/models"
	"localshop/storage"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type UserManager struct {
	db storage.UserStore
}

func NewUserManager(db storage.UserStore) *UserManager {
	return &UserManager{
		db: db,
	}
}

func (m *UserManager) UserCreate(user models.User) (uuid.UUID, error) {
	uid, err := m.db.UserCreate(user)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "Error creating user")
	}
	return uid, nil
}

func (m *UserManager) UserGet(uid uuid.UUID) (models.User, error) {
	user, err := m.db.UserGet(uid)
	if err != nil {
		return models.User{}, errors.Wrap(err, "Error get user")
	}
	return user, nil
}
