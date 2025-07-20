#!/usr/bin/env bash

# This script runs Atlas migrate diff to generate a new migration file.

set -a
source .env
set +a

if [ -z "$1" ]; then
  echo "Usage: $0 <migration_name>"
  exit 1
fi

atlas migrate diff \
  --env gorm \
  --config file://./database/atlas.hcl \
  "$1"