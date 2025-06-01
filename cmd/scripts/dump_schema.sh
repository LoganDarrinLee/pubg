#!/bin/bash

# Exit on error 
set -e

# Set our development environment up
source .env 

# Run our migration script, log output to termianl & migrations log with tee
go run cmd/tools/dump_schema.go > >(tee -a ./logs/migrations/migrations.log) 2>&1

# If dump unsuccessful, exit to avoid errors.
if [ $? -eq 0 ]; then
    echo "Succesfully dumped database.."
else 
    echo "[Error] Failed to dump database schema."
    exit 1
fi 
