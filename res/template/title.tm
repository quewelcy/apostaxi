<!doctype html>
<html lang="en" data-framework="intercoolerjs">

<head>
  <meta charset="utf-8">
  <title>Stav Gennitria</title>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link href="https://fonts.googleapis.com/css?family=Exo+2:300" rel="stylesheet">
  <link href="/static/styles.css" rel="stylesheet">
</head>

<body>
  <div id="flex">

    <div class="pillar" id="pillarid">
      <ul id="pillarLeft" class="pillar-set">
        {{.dir}}
      </ul>
    </div>

    <div class="content" id="contentid">
      {{.content}}
    </div>

  </div>
  <script src="/static/split-1.3.5.min.js"></script>
  <script src="/static/app.js"></script>
</body>

</html>