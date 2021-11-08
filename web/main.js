const SERVER_PATH = "http://localhost:8080"

var GameTimerId 
var ChunkTimerId

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

function stopAllTimers() {
    clearTimeout(GameTimerId);
    clearTimeout(ChunkTimerId);
}

function startGame() {
    stopAllTimers()
    postData = {
        "gameStart": true,
        "gameReset": false
    }
    changeGameSettings(postData)
    GameTimerId = setInterval(() => getCurrentPosition(), 500);
    ChunkTimerId = setInterval(() => requestChunk(), 7000);
}

function stopGame() {
    postData = {
        "gameStart": false,
        "gameReset": false
    }
    changeGameSettings(postData)
    stopAllTimers()
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
    stopAllTimers()
}

function requestChunk() {
    fetch(SERVER_PATH + "/requestChunk", {
        method: 'get',
    })
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
            stopAllTimers()
        } else {
            ctx.fillStyle = "green";
        }
        ctx.clearRect(0, 0, 
                    MapSettings.maxX, 
                    MapSettings.maxY)
        ctx.beginPath();              
        for (let i = 0; i < data.positionPoint.length; i++) {
            ctx.fillRect(data.positionPoint[i].x * MapSettings.objX, 
                         data.positionPoint[i].y * MapSettings.objY, 
                         MapSettings.objX, 
                         MapSettings.objY)
        }
        ctx.closePath();
        ctx.fillStyle = "blue";
        for (let i = 0; i < data.chunkPoint.length; i++) {         
            ctx.fillRect(data.chunkPoint[i].x * MapSettings.objX, 
                         data.chunkPoint[i].y * MapSettings.objY, 
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