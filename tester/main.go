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

func (s server) createTask(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	players := []string{
		"Terry",
		"Graham",
		"John",
		"Eric",
	}
	n := rand.Int() % len(players)

	loc := models.Location{Lat: rand.Float64(), Lon: rand.Float64()}
	userLoc := models.CurrentLocation{Game: "test", Player: players[n], Location: loc}

	json, err := json.Marshal(userLoc)
	if err != nil {
		log.Panic(err)
	}

	err = s.nc.Publish("locations", json)
	if err != nil {
		log.Println("Error making NATS request:", err)
	}

	fmt.Fprintf(w, "Location published")
}

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
	http.HandleFunc("/createTask", s.createTask)
	http.HandleFunc("/healthz", s.healthz)

	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
