-- +goose Up

CREATE TABLE author (
    id SERIAL PRIMARY KEY, 
    sub TEXT,
    name TEXT,
    given_name TEXT,
    family_name TEXT,
    profile TEXT,
    picture TEXT,
    email TEXT,
    email_verified TEXT,
    gender TEXT
);

CREATE TABLE entry (
    id SERIAL PRIMARY KEY,
    title TEXT,
    content TEXT,
    author_id INT REFERENCES author(id)
);

-- +goose Down

DROP TABLE IF EXISTS entry;
DROP TABLE IF EXISTS author;