package mockpostgresstore

//nolint:lll
//go:generate mockgen -destination=./store.go -package=mockpostgresstore crm-system/pkg/store UserRepository,AuthRepository
