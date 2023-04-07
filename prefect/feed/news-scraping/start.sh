#!/bin/sh
set -e
source app.env

echo "start the flow"
exec "$@"