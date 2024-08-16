package db

import (
	"context"
	"fmt"
	"log"

	"github.com/book-shop-grpc/product_service/config"
	"github.com/jackc/pgx/v5"
)

func ConnectToDb(cfg config.PgConfig) (*pgx.Conn, error) {
	ctx := context.Background()

	url := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DatabaseName,
	)
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		log.Println("error on connecting db:", err)
		return nil, err
	}
	return conn, nil
}
