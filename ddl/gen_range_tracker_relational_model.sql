DROP DATABASE IF EXISTS range_tracker CASCADE;
CREATE DATABASE range_tracker

---- Table: range_contestant
DROP TABLE IF EXISTS range_contestant CASCADE;
CREATE TABLE range_contestant
(
	contestant_username primary key
	password text
	contestant_first_name text
	contestant_last_name text
);

---- Table: ranges
DROP TABLE IF EXISTS ranges CASCADE;
CREATE TABLE ranges
(
	range_id smallint primary key generated always as identity
	range_address2 text
	range_city text
	range_zip number(5)
	range_name text
	range_address1 text
	range_state_id smallint references us_states(state_id)
);

---- Table: scores_tracking
DROP TABLE IF EXISTS scores_tracking CASCADE;
CREATE TABLE scores_tracking
(
	score_id smallint primary key generated always as identity
	range_date text
	station_number smallint
	score smallint
	contestant_username references range_contestant(contestant_username)
	range_id smallint references ranges(range_id)
);

---- Table: us_states
DROP TABLE IF EXISTS us_states CASCADE;
CREATE TABLE us_states
(
	state_id smallint primary key generated always as identity
	state_full_name text
	state_abbreviation char(2)
);

