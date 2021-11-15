DATABASE=pdoro
USER=postgres
MIGRATIONS_PATH=./db/migrations

dbup: 
	psql -U ${USER} -d ${DATABASE} -f ${MIGRATIONS_PATH}/up.sql

dbdown:
	psql -U ${USER} -d ${DATABASE} -f ${MIGRATIONS_PATH}/down.sql

dbdump:
	pg_dump -U ${USER} -d ${DATABASE} -s -f ${MIGRATIONS_PATH}/data/dump.sql