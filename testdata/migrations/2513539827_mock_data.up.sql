insert into users(id, email, password_hash, created_at)
	values('18f25d1b-dd0a-4889-9610-d103164c2f2e', 'test@test.com', '$2a$10$lFM6vwGt5SxdIU6z/Ns/zOl6Yz9aBnyyyq8/XJxG/P2p5y9JYyQ6m', current_timestamp);

insert into items(id, user_id, name, category, created_at)
	values('9D4D02A2-2B1C-40B8-9C52-0B75904AE944', '18f25d1b-dd0a-4889-9610-d103164c2f2e', 'blazer', 'blazers', current_timestamp);

insert into item_images(id, item_id, user_id, created_at)
	values('D058078B-29AF-4375-ACAD-7B725A4807E9', '9D4D02A2-2B1C-40B8-9C52-0B75904AE944', '18f25d1b-dd0a-4889-9610-d103164c2f2e', current_timestamp);

insert into items(id, user_id, name, category, created_at)
	values('269881D6-80A9-4C68-BD5B-47010065271E', '18f25d1b-dd0a-4889-9610-d103164c2f2e', 'boots', 'footwear', current_timestamp);

insert into item_images(id, item_id, user_id, created_at)
	values('FC2D7ECB-DD0A-46CA-A132-46A204802F2A', '269881D6-80A9-4C68-BD5B-47010065271E', '18f25d1b-dd0a-4889-9610-d103164c2f2e', current_timestamp);

insert into items(id, user_id, name, category, created_at)
	values('155A829B-35B0-4813-A5D6-7879642F2205', '18f25d1b-dd0a-4889-9610-d103164c2f2e', 'another pair of boots', 'footwear', current_timestamp);

insert into item_images(id, item_id, user_id, created_at)
	values('FE4D6D93-7E7B-47FC-98A0-7E5C0AAA4C8B', '155A829B-35B0-4813-A5D6-7879642F2205', '18f25d1b-dd0a-4889-9610-d103164c2f2e', current_timestamp);

insert into items(id, user_id, name, category, created_at)
	values('623E27FD-A828-4CA3-BAE6-B67B54CC41E9', '18f25d1b-dd0a-4889-9610-d103164c2f2e', 'coat', 'outerwear', current_timestamp);

insert into item_images(id, item_id, user_id, created_at)
	values('523990F4-A5EB-4421-968C-25A51DDE1989', '623E27FD-A828-4CA3-BAE6-B67B54CC41E9', '18f25d1b-dd0a-4889-9610-d103164c2f2e', current_timestamp);

insert into items(id, user_id, name, category, created_at)
	values('9BCAE502-7A0F-4EDC-8EFD-0C58D79DDB72', '18f25d1b-dd0a-4889-9610-d103164c2f2e', 'bowler hat', 'hats', current_timestamp);

insert into item_images(id, item_id, user_id, created_at)
	values('F053EFC2-B053-47EC-A21F-FD5988F8E9F8', '9BCAE502-7A0F-4EDC-8EFD-0C58D79DDB72', '18f25d1b-dd0a-4889-9610-d103164c2f2e', current_timestamp);

insert into items(id, user_id, name, category, created_at)
	values('5A5DB675-70BB-488C-A239-90D1F7BEAEC7', '18f25d1b-dd0a-4889-9610-d103164c2f2e', 'baseball cap', 'hats', current_timestamp);

insert into item_images(id, item_id, user_id, created_at)
	values('F1A13E22-EFF1-4203-8F7F-822A0481D242', '5A5DB675-70BB-488C-A239-90D1F7BEAEC7', '18f25d1b-dd0a-4889-9610-d103164c2f2e', current_timestamp);

insert into items(id, user_id, name, category, created_at)
	values('B77EB2B3-96DD-4155-81E5-3045AA52E8C4', '18f25d1b-dd0a-4889-9610-d103164c2f2e', 'suit', 'suits', current_timestamp);

insert into item_images(id, item_id, user_id, created_at)
	values('E2E80E46-5E31-41DF-ACB3-50BA8516C1DE', 'B77EB2B3-96DD-4155-81E5-3045AA52E8C4', '18f25d1b-dd0a-4889-9610-d103164c2f2e', current_timestamp);

insert into items(id, user_id, name, category, created_at)
	values('7196268A-9710-4C50-AC45-0FE50981A8E4', '18f25d1b-dd0a-4889-9610-d103164c2f2e', 'trousers', 'trousers', current_timestamp);

insert into item_images(id, item_id, user_id, created_at)
	values('105D14AD-7FD4-4C0F-8DAF-F037E6FB93C6', '7196268A-9710-4C50-AC45-0FE50981A8E4', '18f25d1b-dd0a-4889-9610-d103164c2f2e', current_timestamp);
