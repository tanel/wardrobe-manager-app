create table weight_entries(
	id uuid not null,
	user_id uuid not null references users(id) on delete restrict,
	created_at timestamp,
	value float
);