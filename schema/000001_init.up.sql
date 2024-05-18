create table if not exists "user" (
    id          serial primary key unique,
    email       VARCHAR(100) not null unique,
    password    VARCHAR(255) not null,
    name        VARCHAR(100),
    surname     VARCHAR(100),
    is_admin    BOOLEAN,
    is_approved BOOLEAN default false
);

create table if not exists "folder" (
    id      serial primary key unique,
    name    varchar(100) not null,
    user_id integer references "user"(id) on DELETE CASCADE not null,
    is_root BOOLEAN default false,
    is_bin  BOOLEAN default false
);

create table if not exists "file" (
    id         serial primary key unique,
    name       varchar(100) not null,
    path       varchar(255) not null,
    size_bytes integer,
    type       VARCHAR(100),
    user_id    integer references "user"(id) on DELETE CASCADE not null
);

create table if not exists "upload_url" (
    id        serial primary key unique,
    hex       varchar(10) not null,
    hour_live integer default 0 not null
);

create table if not exists "file_url" (
    id      serial primary key unique,
    file_id integer references "file"(id) not null,
    url_id  integer references "upload_url"(id) not null
);