
var geojsonMarkerOptions = {
    radius: 8,
    fillColor: "#ff7800",
    color: "#000",
    weight: 1,
    opacity: 1,
    fillOpacity: 0.8
};

var url = new URL(location.href);
var gameId = url.searchParams.get("game");

var map = L.map('map'),
    trails = [],
    trail = {
        type: 'Feature',
        properties: {
            id: 1
        },
        geometry: {
            type: 'LineString',
            coordinates: []
        }
    },
    realtime = L.realtime(function (success, error) {
        fetch('http://recentlocations.localtest.me/locations/' + gameId)
            .then(function (response) { return response.json(); })
            .then(function (data) {
                var ts = [];

                data.features.forEach(element => {
                    var t = trails[element.id];
                    if (t == undefined) {
                        t = { type: 'Feature', properties: { id: element.id + "-trail" }, geometry: { type: 'LineString', coordinates: [] } };
                        t.id = element.id + "-trail"
                    }

                    t.geometry.coordinates.push(element.geometry.coordinates);
                    t.geometry.coordinates.splice(0, Math.max(0, t.geometry.coordinates.length - 50));
                    trails[element.id] = t;
                    ts.push(t)
                });

                ts.forEach(e => { data.features.push(e) })

                success(data);
            })
            .catch(error);
    }, {
            pointToLayer: function (feature, latlng) {
                console.log(feature)
                if (feature.id == "Terry") {
                    return L.circleMarker(latlng, geojsonMarkerOptions);

                }
                return L.marker(latlng);
            },
            // updateFeature: function (feature, oldLayer, newLayer) {
            //     console.log(feature)
            //     feature.setIcon()
            //     oldLayer.setIcon(newLayer._icon);
            //     return L.Realtime.prototype.options.updateFeature(feature, oldLayer, newLayer)
            // },
            interval: 250
        }).addTo(map);

L.tileLayer('http://{s}.tile.osm.org/{z}/{x}/{y}.png', {
    attribution: '&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
}).addTo(map);


realtime.on('update', function () {
    // console.log(realtime.getBounds())
    map.fitBounds(realtime.getBounds(), { maxZoom: 20 });
});
