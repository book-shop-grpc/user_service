package postgres

import (
	"context"

	//"github.com/book-shop-grpc/user_service/genproto/user_service"
	"github.com/book-shop-grpc/user_service/genproto/user_service"
)

type UserRepoI interface {
	CreateUser(ctx context.Context, user *user_service.User) (*user_service.User, error)
	ListUser(ctx context.Context, req *user_service.GetListReq) (*user_service.GetListResp, error)
	GetUser(ctx context.Context, userId string) (*user_service.User, error)
	UpdateUser(ctx context.Context, user *user_service.User) (*user_service.User, error)
	DeleteUser(ctx context.Context, userId string) error
}

type BasketServiceI interface {
	AddToBasket(ctx context.Context, basket *user_service.Basket) (*user_service.Basket, error)
	GetBasket(ctx context.Context, basket *user_service.GetBasketReq) (*user_service.Basket, error)
	RemoveFromBasket(ctx context.Context, basketId *user_service.GetBasketReq) (*user_service.Empty, error)
}
