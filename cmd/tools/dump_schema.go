// For running database migrations
package main

import (
	"context"
	"log"
	"os"
	"os/exec"

  "github.com/LoganDarrinLee/pubg/internal/db"
	"github.com/LoganDarrinLee/pubg/internal/config"
)

func main() {
	// All logs will go to logs/migartions.log
	log.Println("Starting schema dump.")

	// New environment configuration
	env := config.NewEnv()
	os.Setenv("PGPASSWORD", env.DBPassword)

	// New context
	ctx := context.Background()

	// Database connection
	dbpool := db.InitDB(env, ctx)
	defer dbpool.Close()

	// Export schema for sqlc
	schemafile := "./internal/db/schema/schema.sql"
	if err := exportSchema(env, schemafile); err != nil {
		log.Fatal(err)
	}

	// Successful database migration.
	log.Println("Sucessfully dumped database schema.")
}

// Export our new schema file for sqlc to run.
func exportSchema(env *config.EnvConfig, schemafile string) error {
	// Define pg_dump command with args
	cmd := exec.Command(
		"pg_dump",
		"-U", env.DBUser,
		"-h", env.DBHost,
		"-p", env.DBPort,
		"-s", env.DBName,
	)

	// Setup the file to dump the schema into
	outputFile, err := os.Create(schemafile)
	if err != nil {
		log.Println("Erorr with schema file: ", err)
		return err
	}
	defer outputFile.Close()
	cmd.Stdout = outputFile

	// Execute command
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
