create table outfits(
	id uuid not null primary key,
	user_id uuid not null references users(id),
	name text not null,
	created_at timestamp not null,
 	deleted_at timestamp
);

create index outfits_user_id on outfits(user_id);
create index outfits_created_at on outfits(created_at);
create index outfits_deleted_at on outfits(deleted_at);