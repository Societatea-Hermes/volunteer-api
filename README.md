# Volunteer-api

For compiling the projects server you can take two approaches:

1. Manually create a postgres database (I used PgAdmin4) alongside a .env file in which you can specify these environment variables.
- DB_USER=postgres   (the user of the database)
- DB_PASSWORD=postgres (the password of the db)
- DB_HOST=localhost (the host of the db)
- DB_PORT=5432 (the default port for connecting to the db)
- DB_NAME=some_name  (the name of the db)
- APP_PORT=80 (the port on which the API server will run)

After creating the database and a .env file you can run the following commands to start the server (assuming you are in the root of the project):
```
cd .\server
go mod download
go run .\cmd\api\main.go
```
2. The simpler and more stable approach will need you to have Docker installed.

It's as simple as running the docker compose command.
```
cd .\server
docker compose up -d
``` 