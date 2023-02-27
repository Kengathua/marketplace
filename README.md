# SET UP

## Set Up the DB

    CREATE DATABASE matawis_db;

    CREATE USER matawis_user WITH SUPERUSER CREATEDB CREATEROLE LOGIN PASSWORD 'matawis_pass';

    ALTER ROLE matawis_user SET client_encoding TO 'utf8';

    ALTER ROLE matawis_user SET default_transaction_isolation TO 'read committed';

    ALTER ROLE matawis_user SET timezone TO 'UTC';

    GRANT ALL PRIVILEGES ON DATABASE matawis_db TO matawis_user;

    CREATE DATABASE matawis_test;

    GRANT ALL PRIVILEGES ON DATABASE matawis_test TO matawis_user;

    \q

## Setting up the project

Clone the repository and if you have go set up on your machine run

        go mod download

        go mod tidy

## Migrate down

    migrate -database ${POSTGRESQL_URL} -path ./db/migrations down

## Migrate up

    source env.sh

    go run migrations/migrate.go

OR

    source env.sh

    migrate -database ${POSTGRESQL_URL} -path ./db/migrations up
