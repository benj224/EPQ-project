<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- Latest compiled and minified CSS -->
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css">

	<!-- jQuery library -->
	<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>

	<!-- Popper JS -->
	<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.16.0/umd/popper.min.js"></script>

	<!-- Latest compiled JavaScript -->
	<script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.5.2/js/bootstrap.min.js"></script>
	<link rel="stylesheet" href="newCSS.css">

  </head>

  <body>
  	<div class="container-fluid">
  		<div class="jumbotron jumbotron-fluid" style="margin-bottom: 0px; margin-top: 15px;">
  			<h1 style="font-size: 3.75rem; text-align: center; margin-right: 50px;">Crypto Wallet and Co</h1>
  		</div>
  		<nav class="navbar navbar-expand-md bg-dark navbar-dark">
		  <!-- Brand -->
		  <a class="navbar-brand" href="#">
		  	<img src="bitcoinJpeg-removebg-preview.png" alt="Logo" style="width:40px;">
		  </a>

		  <!-- Toggler/collapsibe Button -->
		  <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#collapsibleNavbar">
		    <span class="navbar-toggler-icon"></span>
		  </button>

		  <!-- Navbar links -->
		  <div class="collapse navbar-collapse" id="collapsibleNavbar">
		    <ul class="navbar-nav">
	      	<li class="nav-item">
		        <a class="nav-link" href="index.html">HOME</a>
		      </li>
		      <li class="nav-item">
		        <a class="nav-link" href="login.php">LOGIN</a>
		      </li>
		      <li class="nav-item">
		        <a class="nav-link" href="logout.php">LOGOUT</a>
		      </li>
		      <li class="nav-item">
		      	<a class="nav-link" href="register.php">REGISTER</a>		      	
		      </li>
		    </ul>
		  </div>
		</nav>
        <div class="row">
			<div class="col-sm-10 offset-sm-1 col-md-6 offset-md-3">
				<form action="login.php" onsubmit="return formValidFunction()" method="POST">
				  <div class="form-group">
				    <label for="email">Email address:</label>
				    <input type="email" class="form-control" placeholder="Enter email" id="email" name="email">
				  </div>
				  <div class="form-group">
				    <label for="pwd">Password:</label>
				    <input type="password" class="form-control" placeholder="Enter password" id="pwd" name="pwd">
				  </div>
				  <div class="form-group form-check">
				    <label class="form-check-label">
				      <input class="form-check-input" type="checkbox" name="remember"> Remember me
				    </label>
				  </div>
				  <button type="submit" class="btn btn-primary" name="submit" value="submit">Submit</button>
				</form>
			</div>
		</div>
	</body>
</html>
