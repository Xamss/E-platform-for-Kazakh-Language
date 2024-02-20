create table users
(
    id              bigserial primary key,
    username        varchar(255) not null,
    email           varchar(255) not null,
    hashed_password varchar(255) not null
)
