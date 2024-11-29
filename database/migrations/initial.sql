
-- +migrate Up
-- +migrate StatementBegin

-- Tabel Kategori
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
);

-- Tabel Buku
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    release_year SMALLINT CHECK (release_year BETWEEN 1980 AND 2024),
    price INT,
    total_page SMALLINT,
    thickness VARCHAR(50),
    category_id INT REFERENCES categories(id) ON DELETE CASCADE,
);

-- Tabel User
CREATE TABLE user_account (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
);
-- +migrate StatementEnd
