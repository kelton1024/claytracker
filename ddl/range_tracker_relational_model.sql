--Create/recreate the database
drop database if exists cascade;
create database range_tracker;
-- \c range_trackers;

--Create tables
drop table if exists us_states cascade;
create table us_states
(
    state_id smallint primary key generated always as identity,
    state_full_name TEXT,
    state_oh char(2)
);

drop table if exists range_contestant cascade;
create table range_contestant
(
    contestant_username text primary key,
    password text,
    contestant_first_name text,
    contestant_last_name text
);

drop table if exists ranges cascade;
create table ranges
(
    range_id smallint primary key generated always as identity,
    range_name text,
    range_address1 text,
    range_address2 text,
    range_city text,
    range_state_id smallint references us_states(state_id),
    range_zip numeric(5)
);

drop table if exists score_tracking cascade;
create table score_tracking
(
    score_id smallint primary key generated always as identity,
    contestant_username text references range_contestant(contestant_username),
    range_id smallint references ranges(range_id),
    range_date timestamp,
    station_number smallint,
    score smallint
);
