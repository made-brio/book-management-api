
-- +migrate Up
-- +migrate StatementBegin

-- Tabel Kategori
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by VARCHAR(255),
    modified_at TIMESTAMP DEFAULT NOW(),
    modified_by VARCHAR(255)
);

-- Tabel Buku
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    image_url VARCHAR(255),
    release_year INT CHECK (release_year BETWEEN 1980 AND 2024),
    price INT,
    total_page INT,
    thickness VARCHAR(50),
    category_id INT REFERENCES categories(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by VARCHAR(255),
    modified_at TIMESTAMP DEFAULT NOW(),
    modified_by VARCHAR(255)
);

-- Tabel User
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    created_by VARCHAR(255),
    modified_at TIMESTAMP DEFAULT NOW(),
    modified_by VARCHAR(255)
);
-- +migrate StatementEnd
