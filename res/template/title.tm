<!doctype html>
<html lang="en" data-framework="intercoolerjs">

<head>
  <meta charset="utf-8">
  <title>Stav Gennitria</title>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link href="https://fonts.googleapis.com/css?family=Exo+2:300" rel="stylesheet">
  <link rel="stylesheet" href="/css/styles.css">
</head>

<body>
  <div id="flex">
    <div class="pillar" id="pillarid">

      <ul id="pillarLeft" class="pillar-set">
        {{.Pillar}}
      </ul>

      <ul id="pillarRight" class="pillar-set">

      </ul>
    </div>

    <div class="content" id="contentid"></div>
  </div>
  <script src="/js/vendor/jquery-3.1.1.min.js"></script>
  <script src="/js/vendor/split-1.3.5.min.js"></script>
  <script src="/js/vendor/intercooler-1.1.1.min.js"></script>
  <script src="/js/app.js"></script>
</body>

</html>