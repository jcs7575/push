<!DOCTYPE html>
<!DOCTYPE html>
<html lang="cn">
  <head>
    <meta charset="utf-8">
    <title>Push 提交</title>
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
                            <a href="#">提交</a>
                        </li>
                        <li>
                            <a href="/push/test">测试</a>
                        </li>
                        <li>
                            <a href="/push/query">查询</a>
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
      <fieldset>
        <legend>
            提交 Push：
        </legend>
        <div class="control-group">
          <label class="control-label" for="optionsCheckbox2">类型</label>
          <div class="controls">
            <div class="row show-grid">
                  <div class="span9">
                    <select id="pushType" name="pushType" required>
                        <option value="udid">UDID</option>
                        <option value="uid">UID</option>
                    </select>
                  </div>
              </div>
          </div>
        </div>

        <div class="control-group">
          <label class="control-label" for="focusedInput">标题</label>
          <div class="controls">
            <input type="text" id="title" name="title" value=""   required  >
            </div>
        </div>
        
        <div class="control-group">
          <label class="control-label" for="focusedInput">描述</label>
          <div class="controls">
            <input type="text" id="subTitle" name="subTitle" value=""  required   >
            </div>
        </div>
        <div class="control-group">
          <label class="control-label" for="focusedInput">Channel</label>
          <div class="controls">
            <input type="text" id="channel" name="channel" value=""  required   >
            </div>
        </div>
        <div class="control-group">
          <label class="control-label" for="focusedInput">发送时间</label>
          <div class="controls">
            <input type="text" id="pushDate" name="pushDate" value=""  required   >
            </div>
        </div>
        <div class="control-group">
          <label class="control-label" for="optionsCheckbox2">显示类型</label>
          <div class="controls">
            <div class="row show-grid">
                  <div class="span9">
                    <select id="disType" name="disType" required>
                        <option value="url">URL</option>
                        <option value="intent">INTENT</option>
                    </select>
                  </div>
              </div>
          </div>
        </div>
        <div class="control-group">
          <label class="control-label" for="focusedInput">url/intent</label>
          <div class="controls">
            <input type="text" id="url" name="url" value="{{.url}}" required>
            </div>
        </div>
        <div class="control-group">
            <label class="control-label" for="fileInput">上传文件</label>
            <div class="controls">
              <input class="input-file" name="fileInput" id="fileInput" type="file" required>
              {{.fileInputErr}}
            </div>
        </div>
       
  </fieldset>
    <div class="form-actions">
        <button type="submit" class="btn btn-primary">
            提交
        </button>
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
