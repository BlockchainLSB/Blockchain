@extends('layouts.master')

@section('content')

<div class="container">	 
	 <div class="col-md-9 clearfix mx-auto pb-4">
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
	 
	 
 	<div class="col-md-9 clearfix mx-auto pb-4">
 		<div class="card shadow">
	 		<div class="card-header"><h3>Project</h3></div>
				
  				<div class="carousel slide projects" data-ride="carousel" id="multi-project">
  					
  					<div class="controls-mid pb-2 text-center">
    					<a class="btn-floating btn-lg" href="#multi-project" data-slide="prev"><i class="fa fa-chevron-circle-left"></i></a>
    					<a class="btn-floating btn-lg" href="#multi-project" data-slide="next"><i class="fa fa-chevron-circle-right"></i></a>
  					</div>
  					
    	      		<div class="carousel-inner pr-2">
  	          		<div class="carousel-item active">
                		<div class="row">
                    		<div class="col-sm">
                    			<div class="card border-success "><div class="card-body text-center">
                    				<p class="h5 card-title font-weight-bold">Project1</p>
                    				<p class="h6 card-text text-secondary mt-3">블록체인 네트워크 구축</p>
                    				<a href="#" class="btn btn-success mt-3">Go</a>
                				</div></div>
                			</div>
                			
                    		<div class="col-sm">	
                    			<div class="card border-success"><div class="card-body text-center">
                    				<p class="h5 card-title font-weight-bold">Project2</p>
                    				<p class="h6 card-text text-secondary mt-3">길찾기 API 사용</p>
                    				<a href="#" class="btn btn-success mt-3">Go</a>
                				</div></div>
            				</div>
            				
            				<div class="col-sm">	
                    			<div class="card border-success"><div class="card-body text-center">
                    				<p class="h5 card-title font-weight-bold">Project3</p>
                    				<p class="h6 card-text text-secondary mt-3">머신러닝</p>
                    				<a href="#" class="btn btn-success mt-3">Go</a>
                				</div></div>
            				</div>
                		</div>
            		</div>
            		
            		<div class="carousel-item">
                		<div class="row">
                    		<div class="col-sm">
                    			<div class="card border-success"><div class="card-body text-center">
                    				<p class="h5 card-title font-weight-bold">Project4</p>
                    				<p class="h6 card-text text-secondary mt-3">월월월</p>
                    				<a href="#" class="btn btn-success mt-3">Go</a>
                				</div></div>
                			</div>
                			
                    		<div class="col-sm">	
                    			<div class="card border-success"><div class="card-body text-center">
                    				<p class="h5 card-title font-weight-bold">Project5</p>
                    				<p class="h6 card-text text-secondary mt-3">10 commit</p>
                    				<a href="#" class="btn btn-success mt-3">Go</a>
                				</div></div>
            				</div>
            				
            				<div class="col-sm">	
                    			<div class="card border-success"><div class="card-body text-center">
                    				<p class="h5 card-title font-weight-bold">Project6</p>
                    				<p class="h6 card-text text-secondary mt-3">10 commit</p>
                    				<a href="#" class="btn btn-success mt-3">Go</a>
                				</div></div>
            				</div>
                		</div>
            		</div>
            		
        		</div>
 			</div>
 		</div>
 	</div>
 	<div class="row pb-4">
 		<div class="col-md-5 clearfix mx-auto">
 			<div class="card shadow">
	 			<div class="card-header"><h3>언어 사용 빈도</h3></div>
	 		</div>
	 	
	 		<div class="card-body text-center">
	 			<canvas id="lang-chart"></canvas>
	 		</div>
	 	</div>
	 	
	 	<div class="col-md-5 clearfix mx-auto">
 			<div class="card shadow">
	 			<div class="card-header"><h3>기여도</h3></div>
	 		</div>
	 	
	 		<div class="card-body text-center">
	 			<canvas id="lang-chart"></canvas>
	 		</div>
	 	</div>
	 	
	 </div>
 	</div>
</div>
<script>
		var config_path = {
		                            type: 'pie',
		                            data: {
		                                datasets: [{
		                                    data: [
		                                        10, 20,
		                                    ],
		                                    backgroundColor: [
		                                        '#B8937F',
		                                        '#F4A3C0',
		                                    ],
		                                    label: 'Dataset 1'
		                                }],
		                                labels: [
		                                    '최근 영상 보기',
		                                    '검색하기',
		                                ]
		                            },
		                            options: {
		                                responsive: true
		                            }
		                        };
	     window.onload = function() {
                var ctx_path = document.getElementById('lang-chart').getContext('2d');
                window.myPie = new Chart(ctx_path, config_path);
            };    
</script>


@endsection	
