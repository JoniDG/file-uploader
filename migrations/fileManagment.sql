create schema if not exists core

create table if not exists core.status
(
    filename varchar not null,
    status   varchar not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

