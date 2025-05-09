#!/bin/sh
set -e

echo "Ожидание базы данных $DB_HOST:$DB_PORT..."

# ожидание доступности базы данных
until pg_isready -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER"; do
  echo "Ожидание PostgreSQL $DB_HOST:$DB_PORT..."
  sleep 2
done

echo "База данных готова к подключению."
exec "$@"