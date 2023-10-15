package postgresstore

import (
	"crm-system/pkg/model"

	uuid "github.com/satori/go.uuid"
)

type UserRepository struct {
	store *PostgresStore
}

func NewUserRepository(store *PostgresStore) *UserRepository {
	return &UserRepository{store: store}
}

func (r *UserRepository) Get(id uuid.UUID) (*model.User, error) {
	var user *model.User
	err := r.store.DB.Where("user_id=?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Update(user *model.User) error {
	return r.store.DB.Where("user_id=?", user.UserID).Updates(&user).Error
}
