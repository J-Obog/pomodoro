HOST=localhost
DATABASE=pdoro
USER=postgres
MIGRATIONS_PATH=./db/migrations

dbup: 
	psql -h ${HOST} -U ${USER} -d ${DATABASE} -f ${MIGRATIONS_PATH}/up.sql

dbdown:
	psql -h ${HOST} -U ${USER} -d ${DATABASE} -f ${MIGRATIONS_PATH}/down.sql