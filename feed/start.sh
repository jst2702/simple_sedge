#!/bin/sh

set -e

echo "run db migration"
# source /app/app.env
export $(cat /app/app.env | xargs) && rails c
echo $DB_SOURCE
/app/migrate -path /app/migrations -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"