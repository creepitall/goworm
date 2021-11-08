const SERVER_PATH = "http://localhost:8080"

var TimerId 

var c = document.getElementById("myCanvas");
var ctx = c.getContext("2d");

var MapSettings = {
    "maxX": 0,
    "maxY": 0,
    "objX": 0,
    "objY": 0
}

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
    })
    .then(function (response) {
        return response.json()
    })
    .then(function (data) {
        console.log('data', data)
        MapSettings.maxX = data.mapSettings.maxX
        MapSettings.maxY = data.mapSettings.maxY
        MapSettings.objX = data.mapSettings.objX
        MapSettings.objY = data.mapSettings.objY
    });
}

function startGame() {
    clearTimeout(TimerId);
    postData = {
        "gameStart": true,
        "gameReset": false
    }
    changeGameSettings(postData)
    TimerId = setInterval(() => getCurrentPosition(), 1000);
}

function stopGame() {
    postData = {
        "gameStart": false,
        "gameReset": false
    }
    changeGameSettings(postData)
    clearTimeout(TimerId);
}

function resetGame() {
    postData = {
        "gameStart": false,
        "gameReset": true
    }
    changeGameSettings(postData)
    ctx.clearRect(0, 0, 
        MapSettings.maxX, 
        MapSettings.maxY)
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
        if (data.Death == true) {
            ctx.fillStyle = "red";
            clearTimeout(TimerId);
        } else {
            ctx.fillStyle = "green";
        }
        ctx.clearRect(0, 0, 
                    MapSettings.maxX, 
                    MapSettings.maxY)
        for (let i = 0; i < data.positionPoint.length; i++) {
            ctx.fillRect(data.positionPoint[i].x * MapSettings.objX, 
                         data.positionPoint[i].y * MapSettings.objY, 
                         MapSettings.objX, 
                         MapSettings.objY)
        }
    });
}       

function moveRect(e){
    switch(e.key){  
        case "ArrowLeft":  // если нажата клавиша влево
            postPosition('left')
            break;
        case "ArrowUp":   // если нажата клавиша вверх
            postPosition('up')
            break;
        case "ArrowRight":   // если нажата клавиша вправо
            postPosition('right')
            break;
        case "ArrowDown":   // если нажата клавиша вниз
            postPosition('down')
            break;
    }
}
 
addEventListener("keydown", moveRect);