create database if not exists network;
use network;

drop table if exists users;

create table users(
    id         int auto_increment primary key,
    name       varchar(50) not null,
    nick       varchar(50) not null unique,
    email      varchar(50) not null,
    pass       varchar(50) not null,
    created_at timestamp   default current_timestamp()
) ENGINE=INNODB;