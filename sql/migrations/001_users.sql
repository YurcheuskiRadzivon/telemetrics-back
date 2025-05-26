-- +goose Up
CREATE TABLE users (
    user_id VARCHAR(36) PRIMARY KEY,  
    username VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20) NOT NULL UNIQUE
);