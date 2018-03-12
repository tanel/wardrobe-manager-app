# Wardrobe Organizer

[![Build Status](https://travis-ci.org/tanel/wardrobe-organizer.svg?branch=master)](https://travis-ci.org/tanel/wardrobe-organizer) [![Go Report Card](https://goreportcard.com/badge/github.com/tanel/wardrobe-organizer)](https://goreportcard.com/report/github.com/tanel/wardrobe-organizer)

Web-based open source wardrobe organizer created while teaching Go programming.

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

# Deployment

## Configure server

Install postgresql and nginx

## Create user

Create user "deploy" on server

## Create deployment folder

Log into remote server and create a folder called "deploy"

## Configure nginx

For example, edit /etc/nginx/sites-enabled/default to look like this

```
server {
        listen 80 default_server;
        listen [::]:80 default_server;

        root /var/www/html;

        index index.html index.htm index.nginx-debian.html;

        server_name wardrobe-organizer.com;

        location / {
                proxy_set_header X-Real-IP $remote_addr;
                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                proxy_set_header Host $host;
                proxy_set_header X-NginX-Proxy true;
                proxy_pass http://localhost:9000/;
        }
}
```

and restart nginx

## Configure systemctl

For example, edit /etc/systemd/system/wardrobe.service to look like this

```
[Unit]
Description=wardrobe

[Service]
Environment=FOO=bar
WorkingDirectory=/home/deploy/wardrobe
ExecStart=/home/deploy/wardrobe/wardrobe-linux
Restart=always

[Install]
WantedBy=multi-user.target
```

## Allow user "deploy" to restart the service

For example, run visudo and append following:

```
Cmnd_Alias MYAPP_CMNDS = /bin/systemctl start wardrobe, /bin/systemctl stop wardrobe
deploy ALL=(ALL) NOPASSWD: MYAPP_CMNDS
```

## Deploy

Execute

	make deploy
