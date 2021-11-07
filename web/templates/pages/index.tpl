{{ define "index" }}
<!doctype html>
<html lang="en">
{{ template "header" }}
    <body>        
        <canvas id="myCanvas" 
            width="640"
            height="480"
            style="border: 1px solid black;">
        </canvas>

        <script>
            var c = document.getElementById("myCanvas");
            var ctx = c.getContext("2d");

            ctx.fillStyle = "green";
            ctx.fillRect(10, 10, 10, 10);
        </script> 

        <div class="container">
            <button onclick="postPosition('up')"> Up </button>  
            <button onclick="postPosition('down')"> Down </button>  
            <button onclick="postPosition('left')"> Left </button>  
            <button onclick="postPosition('right')"> Right </button> 
        </div>  
    </body>
</html>
{{ end }}