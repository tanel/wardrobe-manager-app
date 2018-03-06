# Wardrobe manager app
Web-based open source wardrobe manager with easy data entry and photos import.

# How to install

Start by cloning the repo. Then proceed to database setup.

## Database setup

Install PostgresSQL, for example with brew:

	brew install postgresql

Create database and user:

	createdb wardrobe
	createuser wardrobe

Install "migrate" package

	git clone github.com/wallester/migrate
	cd migrate
	make install

Run migrations

	make migrate

## Test data (optional)

To add test user to database with mock data, run

	make testuser

It creates user "test@test.com" with password "123".

## Start server

	make run

## Migrations

### Create a new migration

	name=my_new_migration_name make migration

### Apply all migrations

	make migrate

### Apply only one migration

	make migrate-up

### Dis-apply one migration

	make migrate-down

