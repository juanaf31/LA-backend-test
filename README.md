# LinkAja Simple REST API

Application for LinkAja assignment.

This application is a simple REST API for checking saldo and transfer using [Golang](https://golang.org/) as it's main programming language.


## Preparation

### What you have to install
* Golang
* MySQL
* Docker
* [GCC](https://jmeubank.github.io/tdm-gcc) for unit testing (for unit testing)
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
    - (Optional) MAIN_SERVER_HOST
    - (Optional) MAIN_SERVER_PORT

### For Production (Using Docker)

* Inside .env file, change STAG environtment value become 
```
STAG=prod
```

## How to run?
Use `make` command to simplify your workflow.

### Run with make command

To run unit testing and start application

```
make dev
```

To run application with docker

```
make build
```

For further command, you can see at file `makefile`

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

## Endpoints
### 1. Check Saldo
- Request 
    - URL : `/account/{account_number}`
    - Method : `GET`
- Response
    - Code : 200
    - Content : 
    ```json 
    {
        "account_number": "555001",
        "customer_number": "1001",
        "balance": "10000"
    } 
    ```
### 2. Transfer
- Request 
    - URL : `/account/{account_number}/transfer`
    - Method : `POST`
    - Body : 
    ```json 
    {
        "to_account_number" : "555002",
        "amount" : 100
    } 
    ```
- Response
    - Code : 201