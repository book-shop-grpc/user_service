package postgres

import (
	"context"

	"github.com/book-shop-grpc/user_service/genproto/product_service"
	"github.com/book-shop-grpc/user_service/genproto/user_service"
)

type userRepoI interface {
	CreateUser(ctx context.Context, user *user_service.User) (*user_service.User, error)
	GetUser(ctx context.Context, userId string) (*user_service.User, error)
	UpdateUser(ctx context.Context, user *user_service.User) (*product_service.Author, error)
	DeleteUser(ctx context.Context, userId string) error
	ListUser(ctx context.Context,  ) (*product_service.AuthorListResponse, error)
}
