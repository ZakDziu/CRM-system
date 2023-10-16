package store

import (
	"crm-system/pkg/config"
	"crm-system/pkg/store/postgresstore"
)

type Store struct {
	User UserRepository
	Auth AuthRepository
}

func NewStore(conf *config.Configs) (*Store, error) {
	postgres, err := postgresstore.NewPostgresStore(&conf.DBPostgresConfig)
	if err != nil {
		return nil, err
	}

	return &Store{
		User: postgres.User(),
		Auth: postgres.Auth(),
	}, nil
}
