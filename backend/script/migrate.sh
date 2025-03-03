#!/bin/bash
set -e
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
else
    echo ".env file not found!"
fi

if ! command -v migrate &> /dev/null; then
  apk add --no-cache curl
  curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
  mv migrate /usr/local/bin/
fi

echo "Waiting for PostgreSQL..."
until PGPASSWORD=$POSTGRES_PASSWORD psql -h $POSTGRES_HOST -p $POSTGRES_PORT -U $POSTGRES_USER -d $POSTGRES_DATABASE -c '\q' 2>/dev/null; do
  echo "Postgres is unavailable - sleeping"
  sleep 2
done
echo "PostgreSQL is up - running migrations."
/usr/local/bin/migrate -path ./internal/migration/postgres -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=${POSTGRES_SSL_MODE}" -verbose up
echo "Migrations completed!"
