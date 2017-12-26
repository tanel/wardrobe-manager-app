create table outfit_items(
	id uuid not null,
	outfit_id uuid not null references outfits(id) on delete cascade,
	item_id uuid not null references items(id) on delete cascade,
	created_at timestamp not null,
	deleted_at timestamp
);

create index outfit_items_outfit_id on outfit_items(outfit_id);
create index outfit_items_item_id on outfit_items(item_id);
create index outfit_items_created_at on outfit_items(created_at);
create index outfit_items_deleted_at on outfit_items(deleted_at);