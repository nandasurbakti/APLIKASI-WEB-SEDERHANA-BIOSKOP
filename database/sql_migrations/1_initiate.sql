-- +migrate Up
-- +migrate StatementBegin

create table bioskop (
    id SERIAL PRIMARY KEY,
 	nama VARCHAR(100) NOT NULL,
 	lokasi VARCHAR(100) NOT NULL,
 	rating FLOAT
)


-- +migrate StatementEnd