-- +goose Up
create table items
(
    id          bigserial primary key,
    campaign_id int,
    name        text      not null,
    description text,
    priority    int,
    removed     bool,
    created_at  timestamp not null default now(),
    updated_at  timestamp
);

create table campaigns
(
    id   bigserial primary key,
    name text
);

-- +goose Down
drop table items;
drop table campaigns;
