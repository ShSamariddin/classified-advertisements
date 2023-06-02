package db

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func GetDSN() string {
	dsn := os.Getenv("PG_DSN")
	if dsn == "" {
		log.Fatal("DSN for PostgreSQL database is required")
	}

	return dsn
}

func NewPool() (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), GetDSN())
	if err != nil {
		return nil, fmt.Errorf("unable to create new db pool: %v", err)
	}
	dbPool = pool

	return pool, nil
}

func GetPool() *pgxpool.Pool {
	if dbPool != nil {
		return dbPool
	}

	pool, err := NewPool()
	if err != nil {
		log.Fatalf("unable to create new db pool: %v", err)
	}
	dbPool = pool

	return pool
}

func Db() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}
