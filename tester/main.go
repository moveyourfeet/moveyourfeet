package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/georace/tester/models"

	"github.com/nats-io/nats"
)

type server struct {
	nc *nats.Conn
}

func (s server) baseRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Basic NATS based microservice example v0.0.1")
}

func (s server) randomWalker(n string) {
	lat := 10.17333
	lng := 56.21281
	deltaLat := 0.00001
	deltaLng := 0.00001

	for {
		r1 := rand.Intn(100)
		r2 := rand.Intn(100)
		if r1 < 10 {
			deltaLat = deltaLat * -1
		}

		if r2 < 10 {
			deltaLng = deltaLng * -1
		}
		lat += deltaLat
		lng += deltaLng

		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%s  %f %f \n", n, lat, lng)

		loc := models.Location{Lat: lat, Lon: lng}
		userLoc := models.CurrentLocation{Game: "test", Player: n, Location: loc}

		json, err := json.Marshal(userLoc)
		if err != nil {
			log.Panic(err)
		}

		err = s.nc.Publish("locations", json)

		if err != nil {
			log.Println("Error making NATS request:", err)
		} else {
			log.Println("Published: ", json)
		}
	}
}

// https://github.com/kelseyhightower/app-healthz
func (s server) healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	var s server
	var err error
	uri := os.Getenv("NATS_URI")

	for i := 0; i < 5; i++ {
		nc, err := nats.Connect(uri)
		if err == nil {
			s.nc = nc
			break
		}

		fmt.Println("Waiting before connecting to NATS at:", uri)
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		log.Fatal("Error establishing connection to NATS:", err)
	}

	fmt.Println("Connected to NATS at:", s.nc.ConnectedUrl())
	http.HandleFunc("/", s.baseRoot)
	http.HandleFunc("/healthz", s.healthz)

	go s.randomWalker("Terry")
	go s.randomWalker("Graham")
	go s.randomWalker("John")
	go s.randomWalker("Eric")

	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
