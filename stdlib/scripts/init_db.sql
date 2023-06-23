-- DROP CREATE DB
DROP DATABASE IF EXISTS book_service_db;
CREATE DATABASE book_service_db;

-- DROP CREATE TABLE
DROP TABLE IF EXISTS books;
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    publication_year INT
);