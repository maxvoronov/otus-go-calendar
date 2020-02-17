-- +goose Up
-- +goose StatementBegin
create extension if not exists "uuid-ossp";

create table events (
    id uuid not null,
    title varchar not null,
    status varchar not null,
    date_from timestamp not null,
    date_to timestamp not null
);

create unique index events_id_uindex on events (id);
alter table events add constraint events_pk primary key (id);
create index events_dates_index on events (date_from, date_to);
create index events_dates_status_index on events (status, date_from, date_to);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists events;
-- +goose StatementEnd
