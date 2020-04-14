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

function saveRepo(repo) {
    localStorage.setItem("repo", repo);
}

function getRepo() {
    return localStorage.getItem("repo");
}

function hasRepo() {
    return localStorage.getItem("repo") != null;
}

function savePath(path) {
    localStorage.setItem("path", path);
}

function getPath() {
    return localStorage.getItem("path");
}

function haspath() {
    return localStorage.getItem("path") != null;
}

function cleanAuthInfo() {
    localStorage.removeItem("username");
    localStorage.removeItem("token");
    localStorage.removeItem("repo");
    localStorage.removeItem("path");
}