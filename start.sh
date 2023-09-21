#!/bin/sh

set -e

echo "run db migration"

echo "db_source"
echo "$DB_SOURCE"

# execute migration up
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"
