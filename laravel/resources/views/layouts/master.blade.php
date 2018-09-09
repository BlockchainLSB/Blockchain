<!doctype html>
<html lang="ko">
 
<head>
    <meta http-equiv="Content-Type" content="text/html, charset=utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link href="https://fonts.googleapis.com/css?family=Poppins:300,500,600" rel="stylesheet">
    
    <script src="/js/jquery-3.3.1.min.js"></script>
    <script src="/js/popper.min.js"></script>
    <script src="/js/bootstrap.min.js"></script>
    <script src="/js/Chart.js"></script>
    <script src="js/Chart.bundle.js"></script>
    
    <link rel="stylesheet" href="/css/bootstrap.min.css">
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Source+Sans+Pro:300,400,700">
	<link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.3.1/css/all.css" integrity="sha384-mzrmE5qonljUremFsqc01SB46JvROS7bZs3IO2EmfFsd15uHvIt+Y8vEf7N7fWAU" crossorigin="anonymous">
    
    <link rel="stylesheet" href="/css/bokyoung_style.css">


    <title>Blockchain LSB</title>
</head>

<body>
	<header>
	 <nav class="navbar navbar-expand-lg navbar-dark bg-dark fixed-top " id="mainNav">
      <div class="container">
        <a class="navbar-brand js-scroll-trigger" href="#page-top">Blockchain LSB</a>
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarResponsive">
          <ul class="navbar-nav ml-auto">
          	<li class="nav-item dropdown">
	            <button class="btn btn-secondary" href="/#"data-toggle="dropdown" aria-haspopup="true" aria-expanded="true">
	            	Login</button>
	            <div class="dropdown-menu p-3" style="width:200px">
	              <form>
	              	<div class="form-group">
					  <div class="input-group mb-1">
					  	  <div class="input-group-prepend">
						    <span class="input-group-text">ID  </span>
						  </div>
						  <input type="text" class="form-control" placeholder="ID">
					  </div>
					  <div class="input-group mb-1">
					  	<div class="input-group-prepend">
						    <span class="input-group-text">PW</span>
						  </div>
						  <input type="password" class="form-control" placeholder="Password">
					  </div>
					  <div class="input-group mb-1">
					  	<div class="input-group-prepend mr-2">
						    <a class="btn btn-secondary" href="#">Login</a>
						  </div>
						  <div class="input-group-append">
						    <a class="btn btn-secondary" href="/register">Register</a>
						  </div>
					  </div>
				  </div>
	              </form>
	            </div>
	        </li>
          	<li class="nav-item dropdown">
            	<a class="nav-link" data-toggle="dropdown"><i class="fas fa-search"></i></a>
            	<div class="dropdown-menu p-3" style="width:300px">
	              <form>
	              	<div class="form-group">
					  <div class="input-group">
						  <input type="text" class="form-control" placeholder="Contributor">
						  <div class="input-group-append">
						    <button class="btn btn-secondary" type="submit">찾기</button>
						  </div>
						</div>
				  </div>
	              </form>
	            </div>
            </li>
            
            <li class="nav-item">
              <a class="nav-link js-scroll-trigger" href="#">Portfolio</a>
            </li>
            <li class="nav-item">
              <a class="nav-link js-scroll-trigger" href="/project">Project</a>
            </li>
            <li class="nav-item">
              <a class="nav-link js-scroll-trigger" href="#">Blockchain</a>
            </li>
            
            
          </ul>
        </div>
      </div>
    </nav>
    </header>
        <div class="content mb-5">

	          @yield('content')    
	    
	    </div>
	    
	    <script>
	    	var reference = document.querySelector('#button-a');
			var popper = document.querySelector('#popper-a');
			var anotherPopper = new Popper(reference, popper, {
			  placement: 'bottom'
			});
	    </script>

</body>
</html>