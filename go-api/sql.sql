create database goCrud;

\c goCrud;

create table if not exists users(
    uid bigserial PRIMARY KEY,
    nickname varchar(15) not null,
    email varchar(40) not null unique,
    password varchar(100) not null,
    status char(1) default '0',
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp
);

create table if not EXISTS wallets(
    public_key varchar(32) PRIMARY key unique not null,
    usr bigint not null,
    balance real default 0.0,
    updated_at timestamp DEFAULT current_timestamp,
    CONSTRAINT wallets_usr_fk foreign  key(usr)
    REFERENCES users(uid)
    on delete CASCADE
    on update CASCADE
);

create table if not EXISTS transactions(
    uid bigserial PRIMARY key,
    origin varchar(32) not null,
    target varchar(32) not null,
    cash real not null,
    message varchar(255),
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT current_timestamp,
    CONSTRAINT transactions_origin_fk foreign key(origin)
    REFERENCES wallets (public_key)
    on delete CASCADE
    on update CASCADE,

    CONSTRAINT transactions_target_fk foreign key(target)
    REFERENCES wallets (public_key)
    on delete CASCADE
    on update CASCADE

);