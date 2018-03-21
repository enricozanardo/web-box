

function initMap() {
    var uluru = {lat: 46.4894107, lng: 11.3208888};
    var map = new google.maps.Map(document.getElementById('map'), {
        zoom: 16,
        center: uluru
    });
    var marker = new google.maps.Marker({
        position: uluru,
        map: map
    });
}