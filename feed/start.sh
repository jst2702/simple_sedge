#!/bin/sh

set -e

echo "run db migration"
cat /app/app.env
source /app/app.env
echo $DB_SOURCE
/app/migrate -path /app/migrations -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"