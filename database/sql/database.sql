CREATE DATABASE quotes;

CREATE TABLE quotes(
    id SERIAL PRIMARY KEY,
    text VARCHAR(255),
    author VARCHAR(255)
);