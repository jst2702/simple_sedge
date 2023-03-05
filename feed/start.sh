#!/bin/sh

set -e

echo "run db migration"

echo "before env file"
echo $DB_SOURCE
echo "contents of the env file\n"

search_dir=/app
for entry in "$search_dir"/*
do
  echo "$entry"
done

echo "loading env file"
source /app/app.env
echo "after env file"
echo $DB_SOURCE
/app/migrate -path /app/migrations -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"