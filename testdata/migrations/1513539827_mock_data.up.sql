insert into users(id, email, password_hash, created_at)
	values('18f25d1b-dd0a-4889-9610-d103164c2f2e', 'test@test.com', '$2a$10$lFM6vwGt5SxdIU6z/Ns/zOl6Yz9aBnyyyq8/XJxG/P2p5y9JYyQ6m', current_timestamp);

insert into items(id, user_id, name, category, created_at)
	values('9D4D02A2-2B1C-40B8-9C52-0B75904AE944', '18f25d1b-dd0a-4889-9610-d103164c2f2e', 'blazer', 'blazers', current_timestamp);

insert into items(id, user_id, name, category, created_at)
	values('269881D6-80A9-4C68-BD5B-47010065271E', '18f25d1b-dd0a-4889-9610-d103164c2f2e', 'boots', 'footwear', current_timestamp);

insert into items(id, user_id, name, category, created_at)
	values('155A829B-35B0-4813-A5D6-7879642F2205', '18f25d1b-dd0a-4889-9610-d103164c2f2e', 'another pair of boots', 'footwear', current_timestamp);

insert into items(id, user_id, name, category, created_at)
	values('623E27FD-A828-4CA3-BAE6-B67B54CC41E9', '18f25d1b-dd0a-4889-9610-d103164c2f2e', 'coat', 'outerwear', current_timestamp);

insert into items(id, user_id, name, category, created_at)
	values('9BCAE502-7A0F-4EDC-8EFD-0C58D79DDB72', '18f25d1b-dd0a-4889-9610-d103164c2f2e', 'hat 1', 'hats', current_timestamp);

insert into items(id, user_id, name, category, created_at)
	values('5A5DB675-70BB-488C-A239-90D1F7BEAEC7', '18f25d1b-dd0a-4889-9610-d103164c2f2e', 'hat 2', 'hats', current_timestamp);

insert into items(id, user_id, name, category, created_at)
	values('B77EB2B3-96DD-4155-81E5-3045AA52E8C4', '18f25d1b-dd0a-4889-9610-d103164c2f2e', 'suit', 'suits', current_timestamp);

insert into items(id, user_id, name, category, created_at)
	values('7196268A-9710-4C50-AC45-0FE50981A8E4', '18f25d1b-dd0a-4889-9610-d103164c2f2e', 'trousers', 'trousers', current_timestamp);
