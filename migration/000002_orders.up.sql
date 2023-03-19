CREATE TABLE IF NOT EXISTS orders (
    id            varchar(50)       not null unique,
    user_id          varchar(255) not null,
    status int not null,
    accrual integer,
    uploaded_at timestamp not null,
    FOREIGN KEY(user_id) REFERENCES users(id)
);