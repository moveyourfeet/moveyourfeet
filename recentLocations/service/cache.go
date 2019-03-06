package service

import (
	"log"
	"sync"

	"github.com/georace/recentLocations/models"

	cache "github.com/patrickmn/go-cache"
)

type CacheStruct struct {
	C     *cache.Cache
	mutex sync.Mutex
}

// StoreLocation saves the current location of a player in a cache
func (cs *CacheStruct) StoreLocation(currentLoc models.CurrentLocation) {

	cs.mutex.Lock()
	locations := make(map[string]models.CurrentLocation)

	if list, found := cs.C.Get(currentLoc.Game); found {
		locations := list.(map[string]models.CurrentLocation)
		log.Print("Found cache: ")
		log.Print(locations)
		locations[currentLoc.Player] = currentLoc

		cs.C.Set(currentLoc.Game, locations, 0)
	} else {
		locations[currentLoc.Player] = currentLoc
		cs.C.Set(currentLoc.Game, locations, 0)
	}
	cs.mutex.Unlock()
}

// GetLocations gets the current location of all players in a game
func (cs *CacheStruct) GetLocations(game string) map[string]models.CurrentLocation {

	if list, found := cs.C.Get(game); found {
		locations := list.(map[string]models.CurrentLocation)
		return locations
	}
	log.Print("Did not find any locations for Game: " + game)
	return make(map[string]models.CurrentLocation)
}
