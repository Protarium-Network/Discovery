# Protarium Discovery Service

A lightweight HTTP server written in Go that serves a discovery XML feed (`Endpoint-index`), fully configurable via environment variables.

## ⚙️ Environment Variables

The server uses the following environment variables (with default values if they are not set):

| Variable     | Description                    |
| ------------ | ------------------------------ |
| `API_HOST`   | Host for the API               |
| `PORTAL_HOST`| Host for the portal            |
| `N3DS_HOST`  | Host for the N3DS service      |
| `HTTP_PORT`  | HTTP server listening port     |

---

## How to run it ?

```bash
go mod tidy
go run main.go
```
