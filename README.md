# GO CATALOG

GO CATALOG is REST API written in Golang, can use multiple database, but now only available for MySQL. You can see frontend application which consumed this API in https://github.com/gedelumbung/react-catalog.

# Setup

- Create database with name `db_catalog` and import `db_catalog.sql`. Adjust database config with your own local environment (inside `.env` file).

- To run this project, you must be place this project under `$GOROOT` path, for example, my `$GOROOT` is inside `/Users/lumbung/Documents/Go`, so this project must be located in `/Users/lumbung/Documents/Go/src/github.com/gedelumbung/go-catalog`.

- Install all dependencies, run `dep ensure`. Run `go run main.go` from root application directory to start application

# Route List

    `GET`/v1/ping
    `GET`/v1/products
    `GET`/v1/products/:id
    `GET`/v1/products/:id/images/:image_id
    `POST`/v1/products
    `DELETE`/v1/products/:id
    
