function hasToken() {
    return localStorage.getItem("token") != null;
}

function getToken() {
    return localStorage.getItem("token")
}


function saveToken(token) {
    localStorage.setItem("token", token);
}


function saveAuthInfo(token) {
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
    localStorage.removeItem("token");
    localStorage.removeItem("repo");
    localStorage.removeItem("path");
}