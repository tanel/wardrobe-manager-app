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
	created_at timestamp not null
);

create index index_items_user_id on items(user_id);
create index index_items_category on items(category);