#!/bin/sh
# This script waits for MySQL to be available before starting the application

set -e

host="$1"
shift
cmd="$@"

until mysql --ssl=0 -h "$host" -u root -ppassword -e 'select 1' > /dev/null 2>&1; do
  echo "Waiting for MySQL..."
  sleep 2
done

echo "MySQL is up - executing command"
exec $cmd
