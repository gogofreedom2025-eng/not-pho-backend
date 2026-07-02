# not-pho-backend

Go + Gin API for the [not-pho](../not-pho-frontend) frontend. Uses GORM with SQLite — a local `dev.db` file acts as your database, seeded with sample beers on first run.

## Prerequisites

- [Go 1.25+](https://go.dev/dl/)

Check that Go is installed:

```bash
go version
```

If you see `command not found` in **Git Bash**, either restart your terminal after installing Go, or run:

```bash
export PATH="/c/Program Files/Go/bin:$PATH"
```

## Getting started

You do **not** need to install SQLite, run migrations, or create `dev.db` yourself. One command does everything.

### 1. Open a terminal in this folder

```bash
cd ~/Repos/not-pho-backend
```

### 2. Start the server (this creates the database)

```bash
go run .
```

On first run:

1. Go downloads dependencies (may take a minute)
2. **`dev.db` is created** in this folder (SQLite file — your whole database)
3. GORM creates the `beers` table automatically
4. Five sample beers are inserted if the table is empty

You should see:

```text
seeded 5 beers
[GIN-debug] Listening and serving HTTP on :8080
```

Confirm the database file exists (in a second terminal, or after stopping the server):

```bash
ls dev.db
```

Leave the server terminal running while you test the API. Press `Ctrl+C` to stop.

### 3. Try the API

**Browser** — open:

- http://localhost:8080/api/ping
- http://localhost:8080/api/beers

**Postman** — create a new request:

| Method | Value |
|--------|-------|
| Method | `GET` |
| URL | `http://localhost:8080/api/beers` |

Click **Send**. You should get a JSON array of 5 seeded beers.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/ping` | Health check — `{ "pong": true, "serverTime": "..." }` |
| GET | `/api/message` | Random encouragement message |
| GET | `/api/beers` | List all beers |
| GET | `/api/beers/:id` | Get one beer by id |
| POST | `/api/beers` | Create a beer |
| PUT | `/api/beers/:id` | Update a beer |
| DELETE | `/api/beers/:id` | Delete a beer |

CORS is enabled for `http://localhost:4200`.

## Postman examples

**List beers**

```
GET http://localhost:8080/api/beers
```

**Get one beer**

```
GET http://localhost:8080/api/beers/1
```

**Create a beer**

```
POST http://localhost:8080/api/beers
Content-Type: application/json

{
  "name": "Hazy Little Thing",
  "brewery": "Sierra Nevada",
  "style": "Hazy IPA",
  "abv": 6.7,
  "description": "Juicy and unfiltered"
}
```

**Update a beer**

```
PUT http://localhost:8080/api/beers/1
Content-Type: application/json

{
  "name": "Two Hearted Ale",
  "brewery": "Bell's",
  "style": "American IPA",
  "abv": 7.0,
  "description": "Updated tasting notes"
}
```

**Delete a beer**

```
DELETE http://localhost:8080/api/beers/1
```

## Database (`dev.db`)

| Question | Answer |
|----------|--------|
| Do I install SQLite? | No — it's bundled via Go (`github.com/glebarez/sqlite`) |
| Do I create `dev.db`? | No — `go run .` creates it on startup |
| Where does it live? | `not-pho-backend/dev.db` (same folder as `main.go`) |
| When does seed data load? | First run only, when the `beers` table is empty |
| Does data persist? | Yes — `dev.db` stays on disk between runs |

**Reset everything:** stop the server, delete the file, start again:

```bash
rm dev.db
go run .
```

You'll see `seeded 5 beers` again.

**Peek inside the DB (optional):** install [DB Browser for SQLite](https://sqlitebrowser.org/), open `dev.db`, and browse the `beers` table.

## Environment

| Variable | Default | Description |
|----------|---------|-------------|
| `DATABASE_PATH` | `dev.db` | SQLite file path (created automatically on startup) |
