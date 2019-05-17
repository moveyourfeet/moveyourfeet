package service

import (
	"log"
	"strconv"
	"sync"

	"github.com/moveyourfeet/moveyourfeet/recentLocations/models"
	cache "github.com/patrickmn/go-cache"
)

type CacheStruct struct {
	C     *cache.Cache
	mutex sync.Mutex
}

var CacheService CacheStruct

// StoreLocation saves the current location of a player in a cache
func (cs *CacheStruct) StoreLocation(currentLoc models.CurrentLocation) {
	k := strconv.Itoa(currentLoc.Game)
	cs.mutex.Lock()
	locations := make(map[string]models.CurrentLocation)

	if list, found := cs.C.Get(k); found {
		locations := list.(map[string]models.CurrentLocation)
		log.Print("Found cache: ")
		log.Print(locations)
		locations[currentLoc.Player] = currentLoc

		cs.C.Set(k, locations, 0)
	} else {
		locations[currentLoc.Player] = currentLoc
		cs.C.Set(k, locations, 0)
	}
	cs.mutex.Unlock()
}

// GetLocations gets the current location of all players in a game
func (cs *CacheStruct) GetLocations(game int) map[string]models.CurrentLocation {

	k := strconv.Itoa(game)

	if list, found := cs.C.Get(k); found {
		locations := list.(map[string]models.CurrentLocation)
		return locations
	}
	log.Print("Did not find any locations for Game: ", game)
	return make(map[string]models.CurrentLocation)
}
