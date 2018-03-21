var form = document.forms.namedItem("fileinfo");
var success = document.querySelector('#success');

// document.getElementById("push").addEventListener("click", function(ev) {

form.addEventListener('submit', function(ev) {

    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/push', true);

    var oOutput = document.querySelector("div"), oData = new FormData(form);

    xhr.onreadystatechange = function() {

        if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {

            var response = JSON.parse(xhr.responseText);
            console.log(response);

            success.innerHTML = response.EmergencySuccessMessage;
            //
            // // Redirect to Dashboad
            // if (response.Allowed == true) {
            //     window.location = "/dashboard";
            // }
        }
    };

    xhr.send(oData);
    ev.preventDefault();
}, false);


