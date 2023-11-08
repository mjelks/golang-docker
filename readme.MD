# Scaffolded GoLang Server

Clone this repo to get a start on a GoLang server with a CRUD-like api service with hot-reloading (CompileDaemon)
Includes a layout similar-ish to a Ruby-on-Rails type layout to get things started:

```

- initializers
- migrations
- models
- controllers

```

These were the steps used to generate the baseline image : 

## Init a local go module namespace : 
`go mod init github.com/<your-module-namespace-here>`

e.g. 

`go mod init github.com/gorm-starter`

## Install Dotenv
`go get github.com/joho/godotenv`

## Install GIN (web/routing framework)
`go get -u github.com/gin-gonic/gin`

## Install GORM
`go get -u gorm.io/gorm`

### install the approprate db adapter/driver
`go get -u gorm.io/driver/postgres`

The above commands give us a baseline `go.mod` && `go.sum` package list that will then be used by our `Dockerfile` and `docker-compose.yml` files.

## Build and Deploy

`docker-compose up --build -d` 