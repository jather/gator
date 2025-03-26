# Gator
## Description
Gator is an RSS-feed-aggre(gator) built with go and postgresql. Fetches RSS feeds at a configurable interval, storing them in the database. Supports multiple profiles, allowing you to browse the RSS posts per-user. 
Mostly a project to learn how to use databases and queries in a program. Also became familiarized with goose and sqlc. 

## Instructions for Installation
### Install postgres and Go
Requires postgresql and Go installed on your machine. 

###Install Goose
```go install github.com/pressly/goose/v3/cmd/goose@latest```
You'll also need to install goose for the database migration when setting up. 

###Install Gator
```go install```
Run "go install" to install gator and be able to run the commands. 

###Configure .gatorconfig.json
Requires a ".gatorconfig.json" in your home directory, with the following contents:
```
{"dburl":"<username>://<postgress_password>:@localhost:5432/gator?sslmode=disable","current_user_name":""}
```
replace <username> with postgres user
replace <postgress_password> with your postgress password. 5432 is the default port for postgres.

### Create the database in psql
```
CREATE DATABASE gator;
```

### Run the goose migration in the gator/sql/schema directory
```
cd sql/schema
goose postgres "postgres://postgres:<your postgres password>@localhost:5432/gator" up
```
## Use
### Commands:
- register <user> - create user and login as user
- login <user> - change user
- reset - reset database
- users- list users
- addfeed <name> <url> - add a feed and automatically follow for current user
- feeds - list all feeds
- follow <url> - follow an added feed for this user
- unfollow <url> - unfollow feed for this user
- following - list followed feeds for this user
- agg <interval> - fetch new posts every set interval. EG. "agg 20s" "agg 10m". Intended to have this running in the background in a different terminal
- browse <limit optional> - view posts for current user, the number of posts shown is up to the limit. default 2

##Possible improvements
shell script for installlation
extending browse functionality- pagination, filtering, search
TUI?
service manager to keep agg command running
