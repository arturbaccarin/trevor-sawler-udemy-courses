<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>JWT Test</title>
    <link rel="icon" href="data:;base64,iVBORw0KGgo=">
    <link href="//cdn.jsdelivr.net/npm/bootstrap@5.2.1/dist/css/bootstrap.min.css" rel="stylesheet"
          integrity="sha384-iYQeCzEYFbKjA/T2uDLTpkwGzCiq6soy8tYaI1GyVh/UjpbCx/TYkiZhlZB6+fzT" crossorigin="anonymous">
    <style>
        pre {
            font-size: 9pt;
        }

        label {
            font-weight: bold;
        }
    </style>
</head>

<body>

<div class="container">
    <div class="row">
        <div class="col">
            <form id="login-form" autocomplete="off">
                <h1 class="mt-3">Login</h1>
                <hr>
                <div class="mb-3">
                    <label for="email" class="form-label">Email address</label>
                    <input type="email" class="form-control" required name="email" id="email"
                           autocomplete="email-new">
                </div>
                <div class="mb-3">
                    <label for="password" class="form-label">Password</label>
                    <input type="password" class="form-control" required name="password" id="password"
                           autocomplete="password-new">
                </div>
                <a class="btn btn-primary" id="login">Login</a>
            </form>
            <hr>
            <div id="tokens" class="d-none">
                <h4>JWT Token</h4>
                <pre id="token"></pre>
                <hr>
                <h4>Refresh Token</h4>
                <pre id="refresh"></pre>
            </div>
            <hr>
            <a href="javascript:void(0);" id="getUserBtn" class="btn btn-outline-secondary">Get User ID 1</a>
            <br>
            <div class="mt-2" style="outline: 1px solid silver; padding: 1em;">
                <pre id="user-output">Nothing from server yet...</pre>
            </div>
            <hr>
            <a id="logout" class="btn btn-danger" href="javascript:void(0)">Logout</a>
        </div>
    </div>
</div>

<script>
    // we store our access token in memory - the only safe place
    let access_token = "";
    let refresh_token = "";

    // get references to UI elements
    let loginForm = document.getElementById("login-form");
    let loginBtn = document.getElementById("login");
    let userBtn = document.getElementById("getUserBtn");
    let userOutput = document.getElementById("user-output");
    let tokensDiv = document.getElementById("tokens");
    let tokenDisplay = document.getElementById("token");
    let refreshTokenDisplay = document.getElementById("refresh");
    let logoutBtn = document.getElementById("logout");

    document.addEventListener("DOMContentLoaded", function() {
        // call refeshTokens; this will, by default, log the user in
        // if he or she has a valid, non-expired __Host-refresh_token cookie.
        refreshTokens();
    })

    let refreshRunning = false;
    let refreshTime = new Date();
    let secondsRemaining = (600 - refreshTime.getSeconds()) * 1000; // every 10 minutes
    // let secondsRemaining = (5 - refreshTime.getSeconds()) * 1000; // every 5 seconds

    function autoRefresh() {
        if (!refreshRunning) {
            setTimeout(function() {
                if (access_token !== "") {
                    setInterval(refreshTokens, 10 * 60 * 1000); // every 10 minutes
                    // setInterval(refreshTokens, 5 * 1000); // every 5 seconds
                }
            }, secondsRemaining);
        }
        refreshRunning = true;
    }

    // login button
    loginBtn.addEventListener("click", function() {
        const payload = {
            email: document.getElementById("email").value,
            password: document.getElementById("password").value,
        }

        const requestOptions = {
            method: "POST",
            credentials: "include",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(payload),
        }

        fetch(`/web/auth`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                if (data.access_token) {
                    access_token = data.access_token;
                    refresh_token = data.refresh_token;
                    setUI(true);
                    autoRefresh();
                }
            })
            .catch(error => {
                alert(error);
            })
    })

    userBtn.addEventListener("click", function() {
        const myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json")
        myHeaders.append("Authorization", "Bearer " + access_token);

        const requestOptions = {
            method: "GET",
            headers: myHeaders,
        }

        fetch("/users/1", requestOptions)
            .then((response) => response.json())
            .then((data) => {
                if (data) {
                    userOutput.innerHTML = JSON.stringify(data, undefined, 4);
                }
            })
            .catch(err => {
                userOutput.innerHTML = "Log in first!";
            })
    })

    function refreshTokens() {
        // we'll send a get request which includes the __Host-refresh-tokien cookie if it exists
        const requestOptions = {
            method: "GET",
            credentials: "include",
        }

        fetch(`/web/refresh-token`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                if (data.access_token) {
                    access_token = data.access_token;
                    refresh_token = data.refresh_token;
                    setUI(true);
                    autoRefresh();
                } else {
                    setUI(false);
                }
            })
            .catch(error => {
                console.log("user is not logged in");
            })
    }

    function setUI(loggedIn) {
        if (loggedIn) {
            tokensDiv.classList.remove("d-none");
            loginForm.classList.add("d-none");
            logoutBtn.classList.remove("d-none");
            tokenDisplay.innerHTML = access_token;
            refreshTokenDisplay.innerHTML = refresh_token;
        } else {
            tokensDiv.classList.add("d-none");
            loginForm.classList.remove("d-none");
            logoutBtn.classList.add("d-none");
            document.getElementById("password").value = "";
            userOutput.innerHTML = "Nothing from server yet...";
            tokenDisplay.innerHTML = "No token!";
            refreshTokenDisplay.innerHTML = "No refresh token!";
        }
    }

    logoutBtn.addEventListener("click", function() {
        access_token = "";
        refresh_token = "";

        fetch("/web/logout", {method: "GET"})
        .then(response => {
            setUI(false);
        })
        .catch(error => {
            userOutput.innerHTML = error;
        })
    })

</script>

</body>

</html>