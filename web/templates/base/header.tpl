{{ define "header" }}
<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">

    <title> hello </title>

    <style>
        .bd-placeholder-img {
          font-size: 1.125rem;
          text-anchor: middle;
          -webkit-user-select: none;
          -moz-user-select: none;
          -ms-user-select: none;
          user-select: none;
        }
  
        @media (min-width: 768px) {
          .bd-placeholder-img-lg {
            font-size: 3.5rem;
          }
        }

        .themed-grid-col {
            padding-top: 15px;
            padding-bottom: 15px;
            background-color: rgba(86, 61, 124, .15);
            border: 1px solid rgba(86, 61, 124, .2);
        }
      </style>

    <script>   
        function postData() {
            data = {
                	"id": "5", 
                    "title": "Special", 
                    "artist": "PR",  
                    "price": 9.99
                }
                
            fetch("/albums", {  
                    method: 'post', 
                    body: JSON.stringify(data)
                }) 
                .then(function (data) {  
                    console.log('Request succeeded with JSON response', data);  
                })  
                .catch(function (error) {  
                    console.log('Request failed', error);  
                });
        }
    </script>
</head>
{{end}}