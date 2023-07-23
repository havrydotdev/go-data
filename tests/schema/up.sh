#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER docker;
    CREATE DATABASE docker;
    GRANT ALL PRIVILEGES ON DATABASE docker TO docker;
    CREATE TABLE test (id serial PRIMARY KEY, name varchar not null, count int not null);

    INSERT INTO test (name, count) VALUES ('test1', 1);
    INSERT INTO test (name, count) VALUES ('test2', 2);
    INSERT INTO test (name, count) VALUES ('test3', 3);
    INSERT INTO test (name, count) VALUES ('test4', 4);
    INSERT INTO test (name, count) VALUES ('test5', 5);
EOSQL
