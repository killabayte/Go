CREATE TABLE users (
    id integer not null primary key,
    name varchar(255) not null,
    email varchar(255) not null unique,
    age integer not null,
    password_hash varchar(255) not null,
)