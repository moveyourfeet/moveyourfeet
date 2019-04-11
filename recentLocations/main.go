package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/georace/recentLocations/docs"
	"github.com/georace/recentLocations/service"
	"github.com/gorilla/handlers"

	"github.com/patrickmn/go-cache"
	"github.com/subosito/gotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	gotenv.Load()

	// Setup NATS
	nc := connectToNat()
	sub := subscripeWorkers(nc)

	// Setup cache
	service.CacheService = service.CacheStruct{}
	service.CacheService.C = cache.New(1*time.Hour, 1*time.Hour)

	// Init router
	port := os.Getenv("PORT")
	router := NewRouter()
	router.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)

	// Create http server
	corsObj := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(corsObj)(router)))

	// Clean up
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println("Cleaning up connections")
		sub.Unsubscribe()
		nc.Close()
		os.Exit(0)
	}()
}
