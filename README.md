## HomeFileStorage

### Note!

This is part of my little system for my personal usage. I DO NOT propuse super ultimate solution for home storage. It's just one of my infinite pet-projects. This sistem write my electrysity availability log, store my books and other files, check whether and many more. For all this stuff I use RaspberryPi 4 with static IP address...  

Developer documentation and tasks will be soon.


### Start

```bash

# start infrastructure
$ docker compose up -d --build


# Create database scheme
$ docker exec homefilestorage-db-1 psql -h localhost -d postgres -U postgres -p 5432 -a -w -f ./app/database/scheme/scheme.sql


# start project
$ go run . -config=config.example.toml
```
