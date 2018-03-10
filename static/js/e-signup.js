var nameErr = document.querySelector('#username-err');
var passwordErr = document.querySelector('#password-err');
var policyErr = document.querySelector('#policy-err');
var success = document.querySelector('#success');

var form = document.forms.namedItem("fileinfo");

var progress = document.getElementById("progress");

form.addEventListener('submit', function(ev) {

    var xhr = new XMLHttpRequest();

    progress.style.display = 'center';

    progress.max = 100;

    xhr.open('POST', '/checksignup', true);

    var oOutput = document.querySelector("div"), oData = new FormData(form);

    xhr.onreadystatechange = function() {

        if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {

            var response = JSON.parse(xhr.responseText);

            if (response.Allowed) {
                progress.hidden = false;
                myLoop ();

                setTimeout(function () {
                    progress.hidden = true;
                    progress.value = 0;
                    success.innerHTML = response.LoginMessage;
                }, 2000);
            }

            nameErr.innerHTML = response.EmailMessage;
            passwordErr.innerHTML = response.PasswordMessage;
            policyErr.innerHTML = response.PolicyMessage;
        }
    };

    xhr.send(oData);
    ev.preventDefault();

}, false);

var i = 1;

function myLoop () {
    setTimeout(function () {
        progress.value += i;
        i++;
        if (i < 100) {
            myLoop();
        }
    }, 8);
}