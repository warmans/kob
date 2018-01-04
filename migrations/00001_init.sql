-- +goose Up

CREATE TABLE author (
    id SERIAL PRIMARY KEY, 
    sub TEXT,
    name TEXT,
    given_name TEXT,
    family_name TEXT,
    profile TEXT,
    picture TEXT,
    email TEXT UNIQUE,
    email_verified TEXT,
    gender TEXT
);

CREATE TABLE entry (
    id SERIAL PRIMARY KEY,
    title TEXT,
    content TEXT,
    author_id INT REFERENCES author(id),
    created_date TIMESTAMP,
    updated_date TIMESTAMP 
);

CREATE TABLE tag (
    id SERIAL PRIMARY KEY,
    label TEXT
);

CREATE TABLE entry_tag (
    entry_id INT REFERENCES entry(id),
    tag_id INT REFERENCES tag(id),
    PRIMARY KEY (entry_id, tag_id)
);

CREATE TABLE comment (
    id SERIAL PRIMARY KEY,
    content TEXT,
    entry_id INT REFERENCES entry(id),
    author_id INT REFERENCES author(id),
    created_date TIMESTAMP,
    updated_date TIMESTAMP 
);

-- +goose Down


DROP TABLE IF EXISTS comment;
DROP TABLE IF EXISTS entry_tag;
DROP TABLE IF EXISTS tag;
DROP TABLE IF EXISTS entry;
DROP TABLE IF EXISTS author;
