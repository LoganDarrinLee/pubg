#!/bin/bash

# Exit on error 
set -e

# There is definitely a better solution for getting the goose to actually work, but this 
# was a quick and easy solution that will help absolutely no one but me. 
USER=$(whoami)

if [[ "$USER" == "organicplant" ]]; then
  GOOSECMD="/home/$USER/go/bin/goose"
elif [[ "$USER" == "ldl" ]]; then 
  GOOSECMD="/home/$USER/.goose/bin/goose"
else
  echo "Looks like you need to fix goose yourself. Checkout cmd/scripts/new_migration.sh"
  exit 1
fi


# Exit if no file variable provided.
if [ -z "$1" ]; then 
    echo "[Error] No name provided for migration file."
    echo "[Solution] make new_migration file=<name>"
    exit 1
fi

# Migrations directory
MIGRATIONS_DIR="./internal/db/migrations"

# Create new migration file prefixed with the current timestamp
$GOOSECMD -dir "$MIGRATIONS_DIR" create "_$1" sql


exit 0

