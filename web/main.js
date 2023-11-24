const SERVER_PATH = "localhost:8080"

const _width = 640
const _height = 480
const _objWidth = 40
const _objHeight = 40

window.onload = function () {
    var conn;

    var c = document.getElementById("myCanvas");
    var ctx = c.getContext("2d");

    function Send(message) {
        if (!conn) {
            return false;
        }
        if (!message) {
            return false;
        }
        conn.send(message);
        return false;
    }

    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + SERVER_PATH + "/ws");
        conn.onclose = function (evt) {};
        conn.onmessage = function (evt) {
            var messages = evt.data.split('\n');
            for (var i = 0; i < messages.length; i++) {
                var data = JSON.parse(messages[i]);
                console.log(data)
                ctx.clearRect(0, 0, _width, _height)

                ctx.beginPath();
                ctx.fillStyle = "green";
                for (let i = 0; i < data.positionPoint.length; i++) {
                    ctx.fillRect(data.positionPoint[i].X * _objWidth,
                                data.positionPoint[i].Y * _objHeight,
                                _objWidth, _objHeight)
                }
                ctx.closePath();

                ctx.fillStyle = "blue";
                for (let i = 0; i < data.applePoint.length; i++) {         
                    ctx.fillRect(data.applePoint[i].X * _objWidth, 
                                data.applePoint[i].Y * _objHeight, 
                                _objWidth, _objHeight)              
                }  
            }
        };
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
    }

    function moveRect(e){
        switch(e.key){
            case "ArrowLeft": 
                Send('left')
                break;
            case "ArrowUp":   
                Send('up')
                break;
            case "ArrowRight":   
                Send('right')
                break;
            case "ArrowDown":
                Send('down')
                break;
        }
    }

    addEventListener("keydown", moveRect);
};