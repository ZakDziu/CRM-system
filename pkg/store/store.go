package store

import (
	"crm-system/pkg/config"
	"crm-system/pkg/store/postgresstore"
)

type PostgresStore struct {
	User UserRepository
	Auth AuthRepository
}

func NewPostgres(conf *config.DBPostgresConfig) (*PostgresStore, error) {
	postgres, err := postgresstore.NewPostgresStore(conf)
	if err != nil {
		return nil, err
	}

	return &PostgresStore{
		User: postgres.User(),
		Auth: postgres.Auth(),
	}, nil
}
