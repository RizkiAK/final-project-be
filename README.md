# Service Golang API

## Tech Stack
- Golang v1.21.1
- Mysql

## How to run

1. Clone it
1. Copy paste `.env.example` and rename it into `.env`
1. Adjust the database in your `.env`
1. Run `go run main.go`. It will download all dependencies and running your application

## Build the binary

Just run `go build main.go` it build the binary named `build-app`

## Deploy to Nginx

Here for [`guide nginx`](deploy-nginx.txt) to deploy this service into nginx

## Postman Documentation

Here for [Postman Documentation](final-project-be.postman_collection.json)