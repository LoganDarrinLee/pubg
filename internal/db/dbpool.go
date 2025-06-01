package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

  "github.com/LoganDarrinLee/pubg/internal/config"
)

// Does not close database pool
func InitDB(env *config.EnvConfig, ctx context.Context) (*pgxpool.Pool) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.DBHost, env.DBPort, env.DBUser, env.DBPassword, env.DBName,
	)

	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
    log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Database connected successfully.")
	return dbpool
}
