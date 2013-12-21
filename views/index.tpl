<!DOCTYPE html>

<html>
  	<head>
    	<title>Push 管理系统</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    	<link href="static/css/bootstrap.min.css" rel="stylesheet">
    	 <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
	    <script src="https://code.jquery.com/jquery.js"></script>
	    <!-- Include all compiled plugins (below), or include individual files as needed -->
	    <script src="static/js/bootstrap.min.js"></script>
	    <!-- Custom styles for this template -->
    	<link href="static/css/signin.css" rel="stylesheet">
	</head>
  	
  	<body>
  		<div class="container">
	      <form class="form-signin" action="/push/test">
	        <h2 class="form-signin-heading">Please sign in</h2>
	        <input type="text" name="username" class="form-control" placeholder="Email address" required autofocus>
	        <input type="password" class="form-control" placeholder="Password" required>
	        <label class="checkbox">
	          <input type="checkbox" value="remember-me"> Remember me
	        </label>
	        <button class="btn btn-lg btn-primary btn-block" type="submit">Sign in</button>
	      </form>
    	</div>	
	</body>
</html>
