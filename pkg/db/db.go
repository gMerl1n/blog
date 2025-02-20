package db

import (
	"context"
	"fmt"
	"os"

	"github.com/gMerl1n/blog/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresDB(ctx context.Context, cfg *config.ConfigDB) (dbpool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.NameDB)
	fmt.Println("Database string")
	fmt.Println(dsn)
	dbpool, err = pgxpool.New(ctx, dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return dbpool, nil

}
