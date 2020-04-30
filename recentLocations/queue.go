package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/moveyourfeet/moveyourfeet/recentLocations/models"
	"github.com/moveyourfeet/moveyourfeet/recentLocations/service"
	"github.com/nats-io/nats.go"
)

const subject string = "locations"

func connectToNat() *nats.Conn {
	uri := os.Getenv("NATS_URI")
	var err error
	var nc *nats.Conn

	for i := 0; i < 5; i++ {
		nc, err = nats.Connect(uri)
		if err == nil {
			break
		}
		fmt.Printf("%v\n", err)

		fmt.Println("Waiting before connecting to NATS at: ", uri)
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		log.Fatal("Error establishing connection to NATS:", err)
	}
	fmt.Println("Connected to NATS at: ", nc.ConnectedUrl())
	return nc
}

func subscripeWorkers(nc *nats.Conn) *nats.Subscription {
	sub, err := nc.Subscribe(subject, func(m *nats.Msg) {

		fmt.Printf("Got new location: %s : %s\n", m.Subject, string(m.Data))
		loc, err := models.NewCurrentLocation(m.Data)
		if err != nil {
			panic(err)
		}
		service.CacheService.StoreLocation(loc)
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Worker subscribed to '%s' for processing requests...", subject)
	return sub
}
