# LinkAja Simple REST API

Application for LinkAja assignment.

This application is a simple REST API for checking saldo and transfer using [Golang](https://golang.org/) as it's main programming language.


## Preparation

### What you have to install
* Golang
* MySQL
* Docker
* (Optional) Install `make` program for your system to run this application easier

### For Development

* Import the database provided in mysql/linkaja.sql
* Inside .env file, change STAG environtment value become 
```
STAG=dev
```
* Inside .env.dev file, there are configurations needed for the api to run, change them according to your system. The most important variable that you have to change is:
    - DB_USER
    - DB_PASS
    - DB_NAME
    - (Optional) MAIN_SERVER_HOST
    - (Optional) MAIN_SERVER_PORT

### For Production (Using Docker)

* Inside .env file, change STAG environtment value become 
```
STAG=prod
```

## How to run?
To run this application, use `make` command to simplify your workflow.

### Run with make command

To run unit testing and start application

```
make dev
```

To run application with docker

```
make build
```

### Run without make command

#### Executing with Go

To run coverge of unit testing with go command

```
go test ./... -cover
```

To run application with go command

```
go run main.go
```

#### Executing with Docker

To run this application 

```
docker-compose up --build
```

