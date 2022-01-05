CREATE table IF NOT EXISTS policy (
    id serial PRIMARY KEY,
    name varchar(64) UNIQUE NOT NULL,
    description varchar(128) NOT NULL
);