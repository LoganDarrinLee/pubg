SHELL := /bin/bash

setup: 
	cmd/scripts/setup.sh 	# Start the Postgres container

stop:
	cmd/scripts/breakdown.sh	# Stop and remove the container

run: 
	cmd/scripts/run.sh # Run app in local docker container.

dump_schema:
	cmd/scripts/dump_schema.sh 
	# Grab and dump schema in the ./internal/db/schema/schema.sql file to view current schema

new_migration:
	cmd/scripts/new_migration.sh $(file)

migrate:
	cmd/scripts/migrate.sh && cmd/scripts/dump_schema.sh

