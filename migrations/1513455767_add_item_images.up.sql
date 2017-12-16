create table item_images(
	id uuid not null primary key,
	item_id uuid not null references items(id) on delete restrict,
	created_at timestamp not null
);
