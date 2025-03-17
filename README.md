# Gin-API-Template

A simple RESTful API built with Go and the Gin framework. This project serves as a boilerplate for creating web services with Gin.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [License](#license)

## Features

- Lightweight and fast
- RESTful API design
- Easy routing with Gin
- Middleware support
- Simple error handling
- JSON response formatting
- JSON Logging support

## Requirements

- Go 1.20+
- Git

## Installation

1. Clone the repository:

    ```bash
    git clone git@github.com:chetuachar/gin-api-template.git
    cd gin-api-template
    ```

2. Install dependencies:

    ```bash
    go mod init gin-api-template
    go mod tidy
    ```

3. Add environment variables which are need to put below mentioned example environment

    ```bash
    export VERSION=v1.0.0
    export LOG_LEVEL=debug
    export LOG_STDOUT=false
    export LOG_FILE=gin-api.log
    export LOG_PATH=./log
    export LOG_MAX_SIZE=10
    export LOG_MAX_BACKUPS=3
    export LOG_MAX_AGE=7
    export LOG_COMPRESS=false
    export SERVER_PORT=8086
    export SELF_API_TOKEN=f144a0da7c41f4725d83c50ea4fdf1bd
    export REQ_RATE_LIMIT=10
    export REQ_RATE_LIMIT_TIME=2
    export LOG_SHOW_PASSWD=XyAdjllll

    ```

## Usage

1. Run the API server:

    ```bash
    go run main.go
    ```

2. Access the API at <http://localhost:8086>

## API Endpoints

| Method | Endpoint                             | Description                               |
|--------|--------------------------------------|-------------------------------------------|
| GET    | `/v1/hi` | Hi back |
| GET    | `/health` | Give you the api running status. |
| GET    | `/logs?passwd={password}` | Load the html template, (logging informations.) |

## Example request

1. Getting information.

    ```bash
    curl -X GET "https://<dns_name>/v1/hi" \
        -H "Auth-Token: {token}" \
        -H "Content-Type: application/json"
    ```

## License

Complete project belonging to chethan nv @chetu_achar