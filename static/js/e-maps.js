

function initMap() {
    var lat = position.coords.latitude;
    var lng = position.coords.longitude;

    var uluru = {lat: lat, lng: lng};

    var map = new google.maps.Map(document.getElementById('map'), {
        zoom: 16,
        center: uluru
    });
    var marker = new google.maps.Marker({
        position: uluru,
        map: map
    });
}