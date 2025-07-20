#!/usr/bin/env bash

# This script runs Atlas migrate apply to apply the migration.

set -a
source .env
set +a
atlas migrate apply \
  --env gorm \
  --config file://./database/atlas.hcl