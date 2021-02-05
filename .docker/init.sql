create database if not exists network;
use network;

drop table if exists users;
drop table if exists followers;

create table users(
    id         int auto_increment primary key,
    name       varchar(50) not null,
    nick       varchar(50) not null unique,
    email      varchar(50) not null unique,
    pass       varchar(100) not null,
    created_at timestamp   default current_timestamp()
) ENGINE=INNODB;

create table followers(
    follower_id int not null,
    followed_id int not null,
    created_at  timestamp default current_timestamp(),
    primary key (followed_id, follower_id),
    foreign key (followed_id) references users(id) on delete cascade,
    foreign key (follower_id) references users(id) on delete cascade
) ENGINE=INNODB;
