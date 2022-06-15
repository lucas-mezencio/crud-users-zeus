create schema if not exists zeus;


create table if not exists users(
    id serial primary key,
    nome varchar not null,
    email varchar unique not null,
    senha varchar not null,
    telefone varchar unique not null
)