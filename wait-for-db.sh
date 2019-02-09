#!/bin/sh
# wait-for-db.sh

#command: ["./wait-for-db.sh", "mysql-server", "go-goa", "-f", "/go/src/github.com/hiromaily/go-goa/resources/tomls/docker.toml"]

set -e

host="$1"
shift
cmd="$@"

while ! mysqladmin ping -h"$host" --silent; do
    >&2 echo "Database is unavailable - sleeping"
    sleep 1
done

>&2 echo "Database is up - executing command"
exec $cmd