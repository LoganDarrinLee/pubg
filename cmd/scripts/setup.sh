#!/bin/bash

# Exit on error 
set -e

# Global environment variables. 
source .env

# Start the docker container
docker compose up -d db
