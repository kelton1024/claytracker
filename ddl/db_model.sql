DROP DATABASE IF EXISTS range_tracker;
CREATE DATABASE range_tracker;

DROP TABLE IF EXISTS states CASCADE;
CREATE TABLE states (
	state_id smallint primary key generated always as identity,
	name text
);

DROP TABLE IF EXISTS ranges CASCADE;
CREATE TABLE ranges(
	range_id smallint primary key generated always as identity,
	address1 text,
	address2 text,
	city text,
	name text,
	zipcode text,
	state_id smallint,
	lat float,
	lng float,
	CONSTRAINT fk_state_range FOREIGN KEY (state_id) REFERENCES states(state_id)
);

DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE users(
	user_id smallint primary key generated always as identity,
	username text,
	first_name text,
	last_name text,
	email text,
	address1 text,
	address2 text,
	city text,
	zipcode text,
	state_id smallint,
	password_hash text,
	CONSTRAINT fk_state_user FOREIGN KEY (state_id) REFERENCES states(state_id)
);

DROP TABLE IF EXISTS outings CASCADE;
CREATE TABLE outings(
	outing_id smallint primary key generated always as identity,
	range_id smallint,
	starttime TIMESTAMPTZ,
	owner smallint,
	CONSTRAINT fk_range_outing FOREIGN KEY (range_id) REFERENCES ranges(range_id)
);

DROP TABLE IF EXISTS stations CASCADE;
CREATE TABLE stations(
	station_id smallint primary key generated always as identity,
	range_id smallint,
	lat float,
	lng float,
	CONSTRAINT fk_range_station FOREIGN KEY (range_id) REFERENCES ranges(range_id)
);

DROP TABLE IF EXISTS groupings CASCADE;
CREATE TABLE groupings(
	groups_id smallint primary key generated always as identity,
	outing_id smallint,
	name text,
	CONSTRAINT fk_outing_group FOREIGN KEY (outing_id) REFERENCES outings(outing_id)
);
