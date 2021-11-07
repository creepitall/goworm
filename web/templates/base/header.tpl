{{ define "header" }}
<head>
    <meta charset="utf-8">
 
    <title> hello </title>

    <style></style>

    <script>   
        function postPosition(positionName = '') {
          postData = {
            "position": positionName
          }
          fetch("/move", {
            method: 'post',
            body: JSON.stringify(postData)
          })
          .then(function (response) {
            return response.json()
          })
          .then(function (data) {
            console.log('data', data)
            ctx.clearRect(0, 0, 640, 480)
            ctx.fillRect(data.x, data.y, 10, 10)
          });
        }
    </script>
</head>
{{end}}