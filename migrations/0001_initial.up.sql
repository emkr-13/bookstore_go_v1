CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    refresh_token VARCHAR(255),
    refresh_token_exp TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    bio VARCHAR(1000),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE publishers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(500),
    description VARCHAR(500),
    phone VARCHAR(20),
    email VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TYPE genre_book AS ENUM (
    'fiction', 'non-fiction', 'mystery', 'fantasy', 'science fiction',
    'biography', 'history', 'romance', 'thriller', 'self-help',
    'children', 'young adult', 'horror', 'poetry', 'cookbook',
    'graphic novel', 'travel', 'health', 'business', 'religion',
    'philosophy', 'art', 'music', 'sports', 'technology',
    'education', 'parenting', 'home and garden', 'crafts and hobbies',
    'computers', 'internet', 'science', 'mathematics', 'engineering',
    'law', 'politics', 'social sciences'
);

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author_id INT REFERENCES authors(id) NOT NULL,
    publisher_id INT REFERENCES publishers(id) NOT NULL,
    isbn VARCHAR(20) NOT NULL,
    price VARCHAR(20) NOT NULL,
    stock VARCHAR(10) NOT NULL,
    description TEXT,
    year INT NOT NULL,
    genre genre_book NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);