CREATE TABLE users

(
    id             serial       not null unique,
    name           varchar(255) not null,
    username       varchar(255) not null,
    surname        varchar(255) not null,
    patronymic     varchar(255) not null,
    password       varchar(255) not null,
    email          varchar(255) not null,
    email_verified boolean                  default false,
    req_date       timestamp with time zone default current_timestamp

)