#!/bin/bash
set -e

if [ -f .env ]; then
  source .env
else
  echo ".env file not found!"
  # exit 1
fi

# Check connection using sslmode=require
PGPASSWORD=$POSTGRES_PASSWORD psql -v ON_ERROR_STOP=1 \
  -d lnk \
  -U lnk \
  -h "${POSTGRES_HOST}" \
  --set=sslmode=require \
  -c "SELECT 1" >/dev/null

# Create user if not exists (if your plan permits user creation)
PGPASSWORD=$POSTGRES_PASSWORD psql -v ON_ERROR_STOP=1 \
  -d lnk \
  -U lnk \
  -h "${POSTGRES_HOST}" \
  --set=sslmode=require \
  -tc "SELECT 1 FROM pg_roles WHERE rolname = '$POSTGRES_USER'" | grep -q 1 || \
  PGPASSWORD="$POSTGRES_PASSWORD" psql -v ON_ERROR_STOP=1 \
  -d lnk \
  -U lnk \
  -h "${POSTGRES_HOST}" \
  --set=sslmode=require \
  -c "CREATE USER $POSTGRES_USER WITH PASSWORD '$POSTGRES_PASSWORD';"

# Create database if not exists (if allowed)
PGPASSWORD=$POSTGRES_PASSWORD psql -v ON_ERROR_STOP=1 \
  -d lnk \
  -U lnk \
  -h "${POSTGRES_HOST}" \
  --set=sslmode=require \
  -tc "SELECT 1 FROM pg_database WHERE datname = '$POSTGRES_DATABASE'" | grep -q 1 || \
  PGPASSWORD="$POSTGRES_PASSWORD" psql -v ON_ERROR_STOP=1 \
  -d lnk \
  -U lnk \
  -h "${POSTGRES_HOST}" \
  --set=sslmode=require \
  -c "CREATE DATABASE $POSTGRES_DATABASE OWNER $POSTGRES_USER;"

# Grant privileges on the database
PGPASSWORD=$POSTGRES_PASSWORD psql -v ON_ERROR_STOP=1 \
  -d lnk \
  -U lnk \
  -h "${POSTGRES_HOST}" \
  --set=sslmode=require <<EOSQL
  GRANT ALL PRIVILEGES ON DATABASE $POSTGRES_DATABASE TO $POSTGRES_USER;
  GRANT ALL PRIVILEGES ON SCHEMA public TO $POSTGRES_USER;
  ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO $POSTGRES_USER;
  ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON SEQUENCES TO $POSTGRES_USER;
EOSQL

echo "Database $POSTGRES_DATABASE created (if allowed) and privileges granted to user $POSTGRES_USER."