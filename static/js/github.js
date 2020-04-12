
const CHECK_TOKEN = "https://api.github.com";
const UPLOAD = 'https://api.github.com/repos';

function submitLogin(token, succHandler, failedHandler) {
    $.ajax({
        url: CHECK_TOKEN,
        type: "GET",
        contentType: "application/json",
        beforeSend: function(xhr) {
            xhr.setRequestHeader("Authorization",token);
            xhr.setRequestHeader("Accept","application/vnd.github.v3+json");
        },
        error: function (err) {
            console.log(JSON.stringify(err));
            if(failedHandler) {
                failedHandler(err);
            }
        },
        success: function (resp) {
            if (succHandler) {
                succHandler(resp);
            }
        }
    });
}

function submitGetRepos(token, succHandler) {
    var repoUrl = CHECK_TOKEN + "/user/repos";
    $.ajax({
        url: repoUrl,
        type: "GET",
        contentType: "application/json",
        beforeSend: function(xhr) {
            xhr.setRequestHeader("Authorization",token);
            xhr.setRequestHeader("Accept","application/vnd.github.v3+json");
        },
        data: {
            "type": "all",
            "sort": "updated",
        },
        error: function (err) {
            console.log(JSON.stringify(err));
        },
        success: function (resp) {
            if (succHandler) {
                succHandler(resp);
            }
        }
    });
}

function submitUploadFile(targetPath, content, token, resultHandler) {
    

    $.ajax({
        url: targetPath,
        type: "PUT",
        contentType: "application/json",
        beforeSend: function(xhr) {
            xhr.setRequestHeader("Authorization", token);
            xhr.setRequestHeader("Accept","application/vnd.github.v3+json");
        },
        data: {
            message: 'upload file',
            content: content +'',
        },
        error: function (err) {
            console.log(JSON.stringify(err));
            if(resultHandler) {
                resultHandler(false, JSON.stringify(err));
            }
        },
        success: function (resp) {
            if (succHandler) {
                resultHandler(true);
            }
        }
    });
}