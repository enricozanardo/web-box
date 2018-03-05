var nameErr = document.querySelector('#username-err');
var passwordErr = document.querySelector('#password-err');
var policyErr = document.querySelector('#policy-err');

var form = document.forms.namedItem("fileinfo");

form.addEventListener('submit', function(ev) {

    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/checksignup', true);

    var oOutput = document.querySelector("div"), oData = new FormData(form);

    xhr.onreadystatechange = function() {

        if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {

            var response = JSON.parse(xhr.responseText);
            console.log(response);

            nameErr.innerHTML = 'Username already taken - <a href="/signin">singin</a>?' + response.Email;
            passwordErr.innerHTML = response.Password;
            policyErr.innerHTML = response.Policy;
        }
    };

    xhr.send(oData);
    ev.preventDefault();
}, false);


