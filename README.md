# jaan-http

This is a simple HTTP server which attempts to make developers life easier by simulating
various real-world server responses.

## Features

- Supports all HTTP methods and status codes (`/:code`)
- Support specifing delay for response (`/:code?delay=1000`)

## Usage

```sh
# Copy env variables
cp .env.example .env
# Start project
go run api/api.go
```
