package postgres

import (
	"context"
	"log"

	"github.com/book-shop-grpc/user_service/genproto/user_service"
	"github.com/jackc/pgx/v5"
)

type basket struct {
	db *pgx.Conn
}

func NewBascet(db *pgx.Conn) BasketServiceI {
	return &basket{db: db}
}

func (b *basket) AddToBasket(ctx context.Context, basket *user_service.Basket) (*user_service.Basket, error) {

	query := `
		INSERT INTO 
			baskets(
				basket_id,
				user_id,
				quantity 
			)VALUES(
				$1, $2, $3
			)`
	_, err := b.db.Exec(
		ctx, query,
		basket.BasketId,
		basket.UserId,
		basket.Quantity,
	)
	if err != nil {
		log.Println("error on creating basket")
		return nil, err
	}
	return basket, nil
}
func (b *basket) GetBasket(ctx context.Context, basketId *user_service.GetBasketReq) (*user_service.Basket, error) {
	query := `
		SELECT 
			basket_id,
			user_id,
			quantity
		FROM
			baskets
		WHERE
			basket_id = $1
	`
	var basket user_service.Basket
	err := b.db.QueryRow(ctx, query, basketId).Scan(
		&basket.BasketId,
		&basket.UserId,
		&basket.Quantity,
	)
	if err != nil {
		log.Println("basket not found:", err)
		return nil, err
	}

	return &basket, nil
}
func (b *basket) RemoveFromBasket(ctx context.Context, basketId *user_service.GetBasketReq) (*user_service.Empty, error) {
	query := `
		DELETE FROM
			baskets
		WHERE
			basket_id = $1	
	`
	_, err := b.db.Exec(ctx,query,basketId)
	if err != nil {
		log.Println("error Deleting basket")
		return nil, err
	}
	return &user_service.Empty{}, nil
}
