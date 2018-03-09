var nameErr = document.querySelector('#username-err');
var passwordErr = document.querySelector('#password-err');
var policyErr = document.querySelector('#policy-err');
var success = document.querySelector('#success');

var form = document.forms.namedItem("fileinfo");

var progress = document.getElementById("progress");

form.addEventListener('submit', function(ev) {

    var xhr = new XMLHttpRequest();

    progress.style.display = 'center';

    progress.hidden = false;
    xhr.onprogress = function (e) {
        progress.max = 100;

        for (i = 10; i < progress.max; i++) {
            setTimeout(function () {
                progress.value += i;
            }, 500);
        }
    }

    xhr.open('POST', '/checksignup', true);

    var oOutput = document.querySelector("div"), oData = new FormData(form);

    xhr.onreadystatechange = function() {

        if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {

            var response = JSON.parse(xhr.responseText);

            nameErr.innerHTML = response.EmailMessage;
            success.innerHTML = response.LoginMessage;
            passwordErr.innerHTML = response.PasswordMessage;
            policyErr.innerHTML = response.PolicyMessage;

            setTimeout(function () {
                progress.hidden = true;
                progress.value = 0;
            }, 1100);
        }
    };

    xhr.send(oData);
    ev.preventDefault();
}, false);

