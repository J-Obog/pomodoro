include .env

MIGRATION_VERSION=$(v)
MIGRATION_PATH=./migrations
DB_CONN_URL=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable

migratedown:
	migrate -path ${MIGRATION_PATH} -database "${DB_CONN_URL}" down ${MIGRATION_VERSION}

migrateup:
	migrate -path ${MIGRATION_PATH} -database "${DB_CONN_URL}" up ${MIGRATION_VERSION}