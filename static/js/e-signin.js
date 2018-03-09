var nameErr = document.querySelector('#username-err');
var passwordErr = document.querySelector('#password-err');
var success = document.querySelector('#success');

var form = document.forms.namedItem("fileinfo");

form.addEventListener('submit', function(ev) {

    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/checksignin', true);

    var oOutput = document.querySelector("div"), oData = new FormData(form);

    xhr.onreadystatechange = function() {

        if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {

            var response = JSON.parse(xhr.responseText);
            console.log(response);

            nameErr.innerHTML = response.EmailMessage;
            passwordErr.innerHTML = response.PasswordMessage;
            success.innerHTML = response.LoginMessage;

            // Redirect to Dashboad
            if (response.Allowed == true) {
                window.location = "/dashboard";
            }
        }
    };

    xhr.send(oData);
    ev.preventDefault();
}, false);


