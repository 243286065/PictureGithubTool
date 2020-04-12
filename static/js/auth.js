function hasToken() {
    return localStorage.getItem("token") != null;
}

function getToken() {
    return localStorage.getItem("token")
}

function hasUsername() {
    return localStorage.getItem("username") != null;
}

function getUsername() {
    return localStorage.getItem("username");
}

function generateToken(username, password) {
    return "Basic " + btoa(username+":"+password);
}

function saveToken(token) {
    localStorage.setItem("token", token);
}

function saveUsername(username) {
    localStorage.setItem("username", username);
}

function saveAuthInfo(username, password, token) {
    saveUsername(username);
    saveToken(token);
}

function cleanAuthInfo() {
    localStorage.removeItem("username");
    localStorage.removeItem("token");
}