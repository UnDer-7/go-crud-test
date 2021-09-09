create table users
(
    id         SERIAL PRIMARY KEY,
    email      VARCHAR                 NOT NULL UNIQUE,
    password   VARCHAR                 NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
