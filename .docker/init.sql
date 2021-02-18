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
    following_id int not null,
    created_at  timestamp default current_timestamp(),
    primary key (following_id, follower_id),

    foreign key (following_id) references users(id) on delete cascade,
    foreign key (follower_id) references users(id) on delete cascade
) ENGINE=INNODB;

create table publications(
    id int auto_increment primary key,
    title varchar(50) not null,
    content varchar(300) not null,
    user_id int not null,
    likes int not null default 0,
    created_at timestamp default current_timestamp(),

    foreign key (user_id) references users(id) on delete cascade
)