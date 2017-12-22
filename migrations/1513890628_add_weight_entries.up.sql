create table weight_entries(
	id uuid not null,
	user_id uuid not null references users(id) on delete restrict,
	value float not null,
	created_at timestamp not null,
	deleted_at timestamp
);