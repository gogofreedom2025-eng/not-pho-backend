# not-pho-backend

Go + Gin API for the [not-pho](../not-pho-frontend) frontend.

## Prerequisites

- Go 1.25+

## Run

```bash
go run .
```

Server listens on `http://localhost:8080`.

## Endpoints

| Method | Path | Response |
|--------|------|----------|
| GET | `/api/ping` | `{ "pong": true, "serverTime": "..." }` |
| GET | `/api/message` | `{ "message": "..." }` |

CORS is enabled for `http://localhost:4200`.
