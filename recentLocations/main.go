package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/georace/recentLocations/models"
	"github.com/georace/recentLocations/service"

	"github.com/nats-io/nats"
	"github.com/patrickmn/go-cache"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func locations(w http.ResponseWriter, r *http.Request) {
	locations := cacheService.GetLocations("test")
	fmt.Fprintln(w, locations)
}

var cacheService service.CacheStruct

func main() {

	uri := os.Getenv("NATS_URI")
	var err error
	var nc *nats.Conn

	for i := 0; i < 5; i++ {
		nc, err = nats.Connect(uri)
		if err == nil {
			break
		}

		fmt.Println("Waiting before connecting to NATS at:", uri)
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		log.Fatal("Error establishing connection to NATS:", err)
	}
	fmt.Println("Connected to NATS at:", nc.ConnectedUrl())

	cacheService = service.CacheStruct{}
	cacheService.C = cache.New(1*time.Hour, 1*time.Hour)

	nc.Subscribe("locations", func(m *nats.Msg) {
		fmt.Println("Got new location:", m.Subject, " : ", string(m.Data))
		loc, err := models.NewCurrentLocation(m.Data)
		if err != nil {
			panic(err)
		}
		cacheService.StoreLocation(loc)
	})

	fmt.Println("Worker subscribed to 'tasks' for processing requests...")
	fmt.Println("Server listening on port 8181...")

	http.HandleFunc("/locations", locations)

	http.HandleFunc("/healthz", healthz)
	if err := http.ListenAndServe(":8181", nil); err != nil {
		log.Fatal(err)
	}
}
