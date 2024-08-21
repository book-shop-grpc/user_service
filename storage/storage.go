package storage

import (
	"github.com/book-shop-grpc/user_service/storage/postgres"
	"github.com/jackc/pgx/v5"
)

type StorageI interface {
	GetStorageRepo() postgres.UserRepoI
}

type storage struct {
	userRepo postgres.UserRepoI
}

func NewStorage(db *pgx.Conn) StorageI {
	return &storage{userRepo: postgres.NewUserRepo(db)}
}

func (s *storage) GetStorageRepo() postgres.UserRepoI {
	return s.userRepo
}
