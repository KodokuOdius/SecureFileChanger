create table if not exists "user"
(
    id         serial primary key,
    email      VARCHAR(255) not null,
    password   TEXT         not null,
    created_at TIMESTAMP    not null default now(),
    is_admin   BOOLEAN
);

create table if not exists "file" (
    id serial primary key,
    name varchar(255) not null,
    path varchar(255) not null,
    created_at TIMESTAMP not null default now(),
    owner integer references "user"(id) ON DELETE CASCADE
);