/* Postgres Production DB */

create table if not exists users (
    id serial primary key,
    name text not null unique,
    email text not null unique,
    password_hash text,
    created_at timestamp default current_timestamp 
);

create table if not exists questions (
    id serial primary key,
    author_id integer,
    title text not null,
    question text not null,
    created_at timestamp default current_timestamp,
    foreign key (author_id) references users (id)
);

create table if not exists answers (
    id serial primary key,
    author_id integer,
    question_id integer,
    message text,
    created_at timestamp default current_timestamp,
    foreign key (author_id) references users (id),
    foreign key (question_id) references questions(id)
);

