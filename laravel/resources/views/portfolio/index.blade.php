@extends('layouts.master')

@section('content')

<div class="container">	 
	 <div class="col-md-9 clearfix mx-auto p-bottom">
	 	<div class="card shadow">
	 		<div class="card-header"><h3>Profile</h3></div>
	 		<img src="/image/profile_happy.jpg" class="rounded-circle ">
	 		<div class="card-body text-center">
	 			<p class="h2 card-title font-weight-bold">김블록</p>
	 			<p class="h6 card-text text-muted">부산대학교</p>
		      	<p class="h6 card-text text-muted">정보컴퓨터공학</p>
		      	<p class="h6 card-text text-muted">hyper@pusan.ac.kr</p>
	 		</div>
	 	</div>
	 </div>
	 
	 
 	<div class="col-md-9 clearfix mx-auto">
 		<div class="card shadow">
	 		<div class="card-header"><h3>Project</h3></div>
			<div class="carousel slide projects" data-ride="carousel">
        		<div class="carousel-inner">
            		<div class="carousel-item active">
                		<div class="row">
                    		
                    		<div class="col-sm">
                    			<div class="card border-success"><div class="card-body text-center">
                    				<p class="h5 card-title font-weight-bold">BlockchainLSB</p>
                    				<div class="row"><p class="h6 card-text text-secondary">10 commit</p></div>
                    		
                    				<a href="#" class="btn btn-success">Go</a>
                				</div></div>
                			</div>
                			
                    		<div class="col-sm"><h4>Project2</h4></div>
                    		<div class="col-sm"><h4>Project3</h4></div>
                		</div>
            		</div>
            		<!--
	            	<div class="carousel-item">
	                	<div class="row">
	                    	<div class="col-sm"><h4>Project4</h4></div>
	                    	<div class="col-sm"><h4>Project5</h4></div>
	                    	<div class="col-sm"><h4>Project6</h4></div>
	                	</div>
	            	</div>
	            	-->
        		</div>
 			</div>
 		</div>
 	</div>
 	
	 <!--
	 <div class="col-md-6 clearfix lang-freq">
	 	<div class="card shadow">
	 		<div class="card-body text-center">
	 			<p class="h4 card-title font-weight-bold">언어 사용빈도</p>
	 			<canvas id="lang-chart"></canvas>
	 		</div>
	 	</div>
	 </div>
	 -->
</div>

@endsection