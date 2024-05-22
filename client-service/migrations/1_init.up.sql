create table "user"
(
    uid        varchar not null
        constraint user_pk
            primary key,
    created_at timestamp default CURRENT_TIMESTAMP,
    email      varchar
);

alter table "user"
    owner to web_confleux_db_user;

create table todo_item
(
    description varchar,
    id          uuid not null
        constraint todo_item_pk
            primary key,
    created_at  timestamp default CURRENT_TIMESTAMP,
    uid         varchar
        constraint todo_item_user_uid_fk
            references "user"
);

alter table todo_item
    owner to web_confleux_db_user;
