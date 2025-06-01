// For running database migrations
package main

import (
	"context"
	"log"
	"os"

  "github.com/LoganDarrinLee/spider/internal/db"
	"github.com/LoganDarrinLee/spider/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func main() {
	// All migration logs will go to logs/migartions.log
	log.Println("Starting database migration.")

	// New environment configuration
	env := config.NewEnv()
	os.Setenv("PGPASSWORD", env.DBPassword)

	// New context
	ctx := context.Background()

	// Database connection
	dbpool := db.InitDB(env, ctx)
	defer dbpool.Close()

	// Migrate database
	migrateDB(dbpool)
}

// Migrate with goose. Takes connection pool and converts into *database/sql.DB
func migrateDB(dbpool *pgxpool.Pool) {
	// Set goose dialect to postgres
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	// Convert database connection pool to *sql.DB for goose
	db := stdlib.OpenDBFromPool(dbpool)
	if err := goose.Up(db, "./internal/db/migrations"); err != nil {
		log.Fatal(err)
	}

	// Explicitly close db connection, not the pool
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}

	// Successful database migration.
	log.Println("Database successfully migrated.")
}


