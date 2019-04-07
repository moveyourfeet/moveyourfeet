From https://github.com/hellojebus/go-mux-jwt-boilerplate

# Game manager

The purpose of this service is to handle games, game rules, and join codes

## PI

See `/documentation/index.html` for API documentation.

To generate new swagger api run

```bash
$ swag init  
```

## Configuration

Make sure to copy `.env.sample` to `.env` and update all fields (DB fields are important!)

**Please note that this is using the Postgres driver, if you prefer to use another driver, update `db.go` accordingly**

Gorm is setup to automigrate the Games table, so it should be plug and play.

## Installation

Make sure to have all required external deps. Look at Godeps config file to view them all.

**Preferred Method, Live Reloading (optional):**

Install Gin `go get github.com/codegangsta/gin`

Then run: `gin`

**Otherwise:**

To run using Go: `go run *.go`

To view application in browser: `localhost:3000 (gin run) or locahost:YOUR_PORT_ENV (go run)`
  
## Todos
 
[] Detach gameHandlers DB functions from http for testing purposes<br>
