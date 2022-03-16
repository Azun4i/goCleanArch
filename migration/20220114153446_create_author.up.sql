CREATE TABLE authors(
    uuid      varchar   not null,
    firstname varchar   not null,
    lastname  varchar   not null,
    email     varchar   not null unique,
    age       integer

);