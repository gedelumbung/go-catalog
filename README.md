# CATALOG API
CATALOG API is REST API written in Golang, can use multiple database, but now only available for MySQL

# Setup

- Create database with name `db_catalog` and import `db_catalog.sql`. Adjust database config with your own local environment (inside `.env` file).
- To run this project, you must be place this project under `$GOROOT` path, for example, my `$GOROOT` is inside `/Users/lumbung/Documents/Go`, so this project must be located in `/Users/lumbung/Documents/Go/src/github.com/gedelumbung/go-catalog`.
- Install all dependencies, run `dep ensure`. Run `go run main.go` from root application directory to start application

