<!DOCTYPE html>
<!DOCTYPE html>
<html lang="cn">
  <head>
    <meta charset="utf-8">
    <title>Push 结果</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link href="static/css/bootstrap.css" rel="stylesheet">
    
     <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <script src="https://code.jquery.com/jquery.js"></script>
    <script src="http://code.jquery.com/jquery-1.10.2.min.js"></script>
    <!-- Include all compiled plugins (below), or include individual files as needed -->
    <script src="static/js/bootstrap.min.js"></script>
    <style type="text/css">
      body {
        padding-top: 60px;
        padding-bottom: 40px;
      }
      .sidebar-nav {
        padding: 9px 0;
      }
    </style>
    <link href="static/css/bootstrap-responsive.css" rel="stylesheet">
    <link type="text/css" rel="stylesheet" href="http://s.wdjimg.com/style/min/g=bootstrap" />
    <div class="navbar navbar-fixed-top">
        <div class="navbar-inner">
            <div class="container">
                <a class="brand" href="#">Push</a>
                <div class="nav-collapse">
                    <ul class="nav">
                        <li>
                            <a href="/push/create">提交</a>
                        </li>
                        <li>
                            <a href="/push/test">测试</a>
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
  </head>

  <body>

    <div class="container">

      <form method="post" action="/push/create" id="form" enctype="multipart/form-data" class="form-horizontal">
      <div class="alert alert-info">
            <strong><h1>{{.result}}</h1></strong> 
      </div>
      <div class="alert alert-error">
      	<strong><h1>{{.errors}}</h1></strong>
      </div>
	  </form>
      <hr>

      <footer>
        <p>© Company 2013</p>
      </footer>

    </div>
    <!-- Le javascript
    ================================================== -->
    <!-- Placed at the end of the document so the pages load faster -->
    <script src="https://code.jquery.com/jquery.js"></script>
      <!-- Include all compiled plugins (below), or include individual files as needed -->
    <script src="static/js/bootstrap.min.js"></script>

  </body>
</html>
