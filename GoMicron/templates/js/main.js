document.addEventListener("DOMContentLoaded", function() {
    var user_email = document.getElementById('email-input');
    var user_password = document.getElementById('password-input');
    var submit_button = document.getElementById('submit-button');
    var baseUrl = 'http://' + window.location.hostname + ':' + window.location.port;

    function submitForm() {
        event.preventDefault(); 
        var data = {
            email: user_email.value,
            password: user_password.value
        };

        fetch(baseUrl + '/gomicron/backend/user/auth/', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify(data)
        }).then(response => {
            if (response.ok) {
                return response.json();
            } else {
                alert("Incorrect email or password. Please try again.");
                throw new Error("Authentication unsuccessful");
            }
        }).then(data => {
            console.log(data);
            document.cookie = `G0M1CR0N4UTHK3Y=${data.token}; expires=Thu, 01 Jan 2030 00:00:00 UTC; path=/`;
            window.location.href = '/dashboard/';
        }).catch((error) => {
            alert("Incorrect email or password. Please try again.");
            throw new Error("Authentication unsuccessful");
        });
    }

    submit_button.addEventListener("click", submitForm);

    user_email.addEventListener("keypress", function(event) {
        if (event.key === "Enter") {
            event.preventDefault();
            submitForm();
        }
    });

    user_password.addEventListener("keypress", function(event) {
        if (event.key === "Enter") {
            event.preventDefault();
            submitForm();
        }
    });
});
