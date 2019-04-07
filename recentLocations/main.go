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
	"github.com/georace/recentLocations/models"
	"github.com/georace/recentLocations/service"
	"github.com/gorilla/handlers"

	"github.com/nats-io/nats"
	"github.com/patrickmn/go-cache"
	"github.com/subosito/gotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	gotenv.Load()

	nc := connectToNat()

	service.CacheService = service.CacheStruct{}
	service.CacheService.C = cache.New(1*time.Hour, 1*time.Hour)

	sub := subscripeWorkers(nc)

	//init router
	port := os.Getenv("PORT")
	router := NewRouter()
	router.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)

	corsObj := handlers.AllowedOrigins([]string{"*"})

	//create http server
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

func connectToNat() *nats.Conn {
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
	return nc
}

func subscripeWorkers(nc *nats.Conn) *nats.Subscription {
	sub, err := nc.Subscribe("locations", func(m *nats.Msg) {
		fmt.Println("Got new location:", m.Subject, " : ", string(m.Data))
		loc, err := models.NewCurrentLocation(m.Data)
		if err != nil {
			panic(err)
		}
		service.CacheService.StoreLocation(loc)
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Worker subscribed to 'locations' for processing requests...")
	return sub
}
