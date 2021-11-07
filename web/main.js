const SERVER_PATH = "http://localhost:8080"

var TimerId 

var c = document.getElementById("myCanvas");
var ctx = c.getContext("2d");

function postPosition(positionName = '') {
    postData = {
        "position": positionName
    }
    fetch(SERVER_PATH + "/currentWay", {
        method: 'post',
        body: JSON.stringify(postData)
    });
}

function getAddTail() {
    fetch(SERVER_PATH + "/addTail", {
        method: 'get',
    })
}

function changeGameSettings(postData = {}) {
    fetch(SERVER_PATH + "/changeGameSettings", {
        method: 'post',
        body: JSON.stringify(postData)
    });
}

function startGame() {
    postData = {
        "gameStart": true,
        "gameReset": false
    }
    changeGameSettings(postData)
    TimerId = setInterval(() => getCurrentPosition(), 2000);
}

function stopGame() {
    postData = {
        "gameStart": false,
        "gameReset": false
    }
    changeGameSettings(postData)
    clearTimeout(TimerId);
}

function getCurrentPosition() {
    fetch(SERVER_PATH + "/currentPosition", {
        method: 'get',
    })
    .then(function (response) {
        return response.json()
    })
    .then(function (data) {
        console.log('data', data)
        ctx.fillStyle = "green";
        ctx.clearRect(0, 0, 
                    data.mapSettings.maxX, 
                    data.mapSettings.maxY)
        for (let i = 0; i < data.positionPoint.length; i++) {
            ctx.fillRect(data.positionPoint[i].x * data.mapSettings.objX, 
                         data.positionPoint[i].y * data.mapSettings.objY, 
                         data.mapSettings.objX, 
                         data.mapSettings.objY)
        }
    });
}       