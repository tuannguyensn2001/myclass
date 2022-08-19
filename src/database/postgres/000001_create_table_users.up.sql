create table users (
    id serial primary key ,
    username varchar,
    password varchar,
    email varchar,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
)