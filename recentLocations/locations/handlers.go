package locations

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	customHTTP "github.com/moveyourfeet/moveyourfeet/recentLocations/http"
	"github.com/moveyourfeet/moveyourfeet/recentLocations/service"
	geojson "github.com/paulmach/go.geojson"
)

// ShowHandler godoc
// @Summary Get player locations for a game
// @Description Get player locations for a game
// @Accept  json
// @Produce  json
// @Success 200 {object} interface{}
// @Failure 404 {object} http.ErrorResponse
// @Param gameId path int true "Game ID"
// @Router /locations/{gameId} [get]
// @Tags Locations
func ShowHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	gameID, err := strconv.Atoi(params["gameId"])
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
	}
	locations := service.CacheService.GetLocations(gameID)

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

	customHTTP.NewResponseFromJson(w, rawJSON)
}
