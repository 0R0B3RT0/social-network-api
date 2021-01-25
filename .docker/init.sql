create database if not exists network;
use network;

drop table if exists usuarios;

create table usuarios(
    id int auto_increment  primary key,
    name       varchar(50) not null,
    nick       varchar(50) not null unique,
    email      varchar(50) not null,
    pass       varchar(50) not null,
    created_at timestamp   default current_timestamp()
) ENGINE=INNODB;