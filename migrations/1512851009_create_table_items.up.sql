create table items(
	id uuid not null primary key,
	user_id uuid not null references users(id) on delete restrict,
	name text not null,
	description text,
	color text,
	size text,
	brand text,
	price float,
	currency text,
	category text,
	season text not null default 'all-ywar',
	formal bool not null default false,
	created_at timestamp not null,
	deleted_at timestamp
);

create index index_items_user_id on items(user_id);
create index index_items_category on items(category);
create index index_items_created_at on items(created_at);
create index index_items_deleted_at on items(deleted_at);