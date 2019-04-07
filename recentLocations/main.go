package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/georace/recentLocations/models"
	"github.com/georace/recentLocations/service"

	"github.com/nats-io/nats"
	"github.com/patrickmn/go-cache"
	geojson "github.com/paulmach/go.geojson"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

var cacheService service.CacheStruct

// TODO Return GeoJson
// - https://github.com/paulmach/go.geojson
// - http://geojson.io/#map=2/34.7/28.1
// - https://github.com/perliedman/leaflet-realtime
func locations(w http.ResponseWriter, r *http.Request) {
	locations := cacheService.GetLocations("test")

	fc := geojson.NewFeatureCollection()

	for _, v := range locations {
		f := geojson.NewPointFeature([]float64{v.Location.Lat, v.Location.Lon})
		f.SetProperty("id", v.Player)
		f.ID = v.Player
		fc.AddFeature(f)
	}

	rawJSON, err := fc.MarshalJSON()
	if err != nil {
		fmt.Fprintln(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(rawJSON)
}

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

	sub, _ := nc.Subscribe("locations", func(m *nats.Msg) {
		fmt.Println("Got new location:", m.Subject, " : ", string(m.Data))
		loc, err := models.NewCurrentLocation(m.Data)
		if err != nil {
			panic(err)
		}
		cacheService.StoreLocation(loc)
	})
	fmt.Println("Worker subscribed to 'locations' for processing requests...")

	http.HandleFunc("/locations", locations)
	http.HandleFunc("/healthz", healthz)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server listening on port 80...")

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
