#!/bin/bash

# Exit on error 
set -e

# Global environment variables. 
source .env

docker compose down 
