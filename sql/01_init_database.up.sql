BEGIN;
create table articles
(
    id         bigserial
        constraint articles_pk
            primary key,
    writer     varchar                                               not null,
    body       text                                                  not null,
    time       timestamp without time zone default CURRENT_TIMESTAMP not null,
    view_count int                         default 0
);
COMMIT;