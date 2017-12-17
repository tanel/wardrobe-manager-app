# Wardrobe manager app
Web-based open source wardrobe manager with easy data entry and photos import.

# How to install

Start by cloning the repo. Then proceed to database setup.

## Database setup

Install PostgresSQL, for example with brew:

	brew install postgresql

Create database and user:

	createdb wardrobe
	createdb wardrobe

Install "migrate" package

	git clone github.com/wallester/migrate
	cd migrate
	make install

Run migrations

	make migrate

## Start server

	make run
