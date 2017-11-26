<!doctype html>
<html lang="en" data-framework="intercoolerjs">

<head>
    <meta charset="utf-8">
    <title>Khronos</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link href="https://fonts.googleapis.com/css?family=Exo+2:300" rel="stylesheet">
    <link rel="stylesheet" href="/css/styles.css">
</head>

<body>
    <form ic-post-to='/tls' ic-target='tl_wrapper'>
        <select name="places" multiple size='2'>{{.Places}}</select>
        <button>Update</button>
    </form>
    <div class='tl_wrapper'>{{.TimelineLeft}}</div>
    <div class='tl_wrapper'>{{.TimelineRight}}</div>
</body>

</html>