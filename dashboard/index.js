var map = L.map('map'),
    realtime = L.realtime('http://localhost:8181/locations', {
        interval: 500
    }).addTo(map);

L.tileLayer('http://{s}.tile.osm.org/{z}/{x}/{y}.png', {
    attribution: '&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
}).addTo(map);

realtime.on('update', function() {
    console.log(realtime.getBounds())
    map.fitBounds(realtime.getBounds(), {maxZoom: 20});
});
