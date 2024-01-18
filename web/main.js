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

    function FillText(ctx, x, y) {
        ctx.font = "14px Arial";
        ctx.fillStyle = "white";
        pos = x + ',' + y           
        ctx.fillText(pos, x * _objWidth + 5, y * _objHeight + 20);  
    }

    function FillRect(ctx, x, y, color) {
        ctx.fillStyle = color;
        ctx.fillRect(x * _objWidth,
                     y * _objHeight,
                    _objWidth, _objHeight)    
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
                for (let i = 0; i < data.positionPoint.length; i++) {
                    FillRect(ctx, data.positionPoint[i].X, data.positionPoint[i].Y, "green")
                    FillText(ctx, data.positionPoint[i].X, data.positionPoint[i].Y)                  
                }
                ctx.closePath();

                for (let i = 0; i < data.applePoint.length; i++) { 
                    FillRect(ctx, data.applePoint[i].X, data.applePoint[i].Y, "red")     
                    FillText(ctx, data.applePoint[i].X, data.applePoint[i].Y)                                    
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