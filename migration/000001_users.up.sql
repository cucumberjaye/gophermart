CREATE TABLE IF NOT EXISTS users (
    id            varchar(50)       not null unique,
    login          varchar(255) not null unique,
    password_hash varchar(255) not null,
    PRIMARY KEY(id)
);