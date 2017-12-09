create table users(
	id uuid not null primary key,
	name text,
	email text unique not null,
	password_hash text not null
);
