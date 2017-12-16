create table item_images(
	id uuid not null primary key,
	item_id uuid not null references items(id) on delete restrict,
	user_id uuid not null references users(id) on delete restrict,
	created_at timestamp not null,
	deleted_at timestamp
);

create index index_item_images_item_id on item_images(item_id);
create index index_item_images_user_id on item_images(user_id);
create index index_item_images_created_at on item_images(created_at);
create index index_item_images_deleted_at on item_images(deleted_at);
