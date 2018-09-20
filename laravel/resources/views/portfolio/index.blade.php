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
 	
 	<div class="row ">
 		<div class="col-md-5 clearfix mx-auto">
 			<div class="card h-100 shadow">
	 			<div class="card-header"><h3>언어 사용 빈도</h3></div>
	 			<div class="card-body text-center">
	 				<canvas id="lang-chart"></canvas>
	 			</div>
	 		</div>
	 	</div>
	 	
	 	<div class="col-md-5 clearfix mx-auto">
 			<div class="card h-100 shadow">
	 			<div class="card-header"><h3>기여도</h3></div>
	 			<div class="card-body pt-5">
	 				<div class="row">
 						<p class="h6 card-text text-muted mx-auto">평균commit 수</p>
 						<p class="h6 card-text text-muted mx-auto">총commit 수</p>
	 				</div>
	 				
	 				<div class="row">
 						<p class="h1 mx-auto pl-4 num" id="tot-commit">582</p>
 						<p class="h1 mx-auto pl-3 num" id="avg-commit">2341</p>
	 				</div>
	 			</div>
	 		</div>
	 	</div>
	 	
	 	<div class="col-md-11 clearfix mx-auto mt-4">
	 		<div class="card shadow">
		 		<div class="card-header"><h3>자격증</h3></div>
		 		<div class="card-body text-center">
		 			<div class="row">
		 				<div class="col">
		 					<p class="h6 card-text text-muted mx-auto">TOEIC</p>
		 					<canvas id="toeic-chart"></canvas>
		 				</div>
		 				
		 				<div class="col">
		 					<p class="h6 card-text text-muted mx-auto">TOPCIT</p>
		 					<canvas id="topcit-chart"></canvas>
		 				</div>
						
						<div class="col">
		 					<p class="h6 card-text text-muted mx-auto">TOEIC SPEAKING</p>
		 					<canvas id="toeicsp-chart"></canvas>
		 				</div> 			
	 				</div>
	 			</div>
		 	</div>
	 	</div>
	 	
	 </div>
 	</div>
</div>
<script>
	var lang_config = {
        type: 'pie',
        data: {
            datasets: [{
                data: [
                    10, 20, 30, 10
                ],
                backgroundColor: [
                    '#CEF19E',
                    '#A7DDA7',
                    '#78BE97',
                    '#398689'
                ],
                label: 'lang-chart'
            }],
            labels: [
                'Java',
                'Go',
                'C++',
                'C'
            ]
        },
        options: {
            responsive: true
        }
    };
    
    var toeic_config = {
        type: 'doughnut',
        data: {
            datasets: [{
                data: [
                    920,
                    80
                ],
                backgroundColor: [
                    '#78BE97',
                    '#DEDEDE'              
                ],
                label: 'Toeic'
            }],
        },
        options: {
            responsive: true,
            rotation: 1 * Math.PI,
      		circumference: 1 * Math.PI,
      		tooltips: {
	            callbacks: {
	                label: function(tooltipItem, data) {
	                    var label = data.datasets[tooltipItem.datasetIndex].label || '';
	
	                    if (label && tooltipItem.index != 1) {
	                        label += ': ';
	                        label += data.datasets[tooltipItem.datasetIndex].data[tooltipItem.index] || '';
	                    } else 
	                    	tooltipItem.enabled = false;
 
	                    return label;
	                }
	            }
        	},        
        	elements: {
				center: {
					text: '920',
			        color: '#398689',
			        fontStyle: 'Helvetica',
			        sidePadding: 20 
				}
			}
        }
    };
    
    var topcit_config = {
        type: 'doughnut',
        data: {
            datasets: [{
                data: [
                    400,
                    1000-400
                ],
                backgroundColor: [
                    '#78BE97',
                    '#DEDEDE'              
                ],
                label: 'topcit'
            }],
        },
        options: {
            responsive: true,
            rotation: 1 * Math.PI,
      		circumference: 1 * Math.PI,
      		tooltips: {
	            callbacks: {
	                label: function(tooltipItem, data) {
	                    var label = data.datasets[tooltipItem.datasetIndex].label || '';
	
	                    if (label && tooltipItem.index != 1) {
	                        label += ': ';
	                        label += data.datasets[tooltipItem.datasetIndex].data[tooltipItem.index] || '';
	                    } else 
	                    	tooltipItem.enabled = false;
 
	                    return label;
	                }
	            }
        	},        
        	elements: {
				center: {
					text: '400',
			        color: '#398689',
			        fontStyle: 'bold',
			        sidePadding: 20 
				}
			}
        }
    };
    
    var toeicsp_config = {
        type: 'doughnut',
        data: {
            datasets: [{
                data: [
                    700,
                    300
                ],
                backgroundColor: [
                    '#78BE97',
                    '#DEDEDE'              
                ],
                label: 'Toeic-Speaking'
            }],
        },
        options: {
        	responsive: true,
            rotation: 1 * Math.PI,
      		circumference: 1 * Math.PI,
      		tooltips: {
	            callbacks: {
	                label: function(tooltipItem, data) {
	                    var label = data.datasets[tooltipItem.datasetIndex].label || '';
	
	                    if (label && tooltipItem.index != 1) {
	                        label += ': ';
	                        label += data.datasets[tooltipItem.datasetIndex].data[tooltipItem.index] || '';
	                    } else 
	                    	tooltipItem.enabled = false;
 
	                    return label;
	                }
	            }
    		},        
        	elements: {
				center: {
					text: '700',
			        color: '#398689',
			        fontStyle: 'Helvetica',
			        sidePadding: 20 
				}
			}
        }
    };
    
    Chart.pluginService.register({
			beforeDraw: function (chart) {
				if (chart.config.options.elements.center) {
			        //Get ctx from string
			        var ctx = chart.chart.ctx;
			        
							//Get options from the center object in options
			        var centerConfig = chart.config.options.elements.center;
			      	var fontStyle = centerConfig.fontStyle || 'Arial';
					var txt = centerConfig.text;
			        var color = centerConfig.color || '#000';
			        var sidePadding = centerConfig.sidePadding || 20;
			        var sidePaddingCalculated = (sidePadding/100) * (chart.innerRadius * 2)
			        //Start with a base font of 30px
			        ctx.font = "50px " + fontStyle;
			        
							//Get the width of the string and also the width of the element minus 10 to give it 5px side padding
			        var stringWidth = ctx.measureText(txt).width;
			        var elementWidth = (chart.innerRadius * 2) - sidePaddingCalculated;
			
			        // Find out how much the font can grow in width.
			        var widthRatio = elementWidth / stringWidth;
			        var newFontSize = Math.floor(30 * widthRatio);
			        var elementHeight = (chart.innerRadius * 2);
		
			        // Pick a new font size so it will not be larger than the height of label.
			        var fontSizeToUse = Math.min(newFontSize, elementHeight);
			
							//Set font settings to draw it correctly.
			        ctx.textAlign = 'center';
			        ctx.textBaseline = 'middle';
			        var centerX = ((chart.chartArea.left + chart.chartArea.right) / 2);
			        var centerY = ((chart.chartArea.top + chart.chartArea.bottom)/2 + 50);
			        ctx.font = fontSizeToUse+"px " + fontStyle;
			        ctx.fillStyle = color;
			        
			        //Draw text in center
			        ctx.fillText(txt, centerX, centerY);
				}
		}
	});
		
 	window.onload = function() {
    	var lang_path = document.getElementById('lang-chart').getContext('2d');
    	window.myPie = new Chart(lang_path, lang_config);
    	
    	var toeic_path = document.getElementById('toeic-chart').getContext('2d');
    	window.myPie = new Chart(toeic_path, toeic_config);
    
    	
    	var ctx = document.getElementById('topcit-chart').getContext('2d');
    	window.myPie = new Chart(ctx, topcit_config);
    	
    	var toeicsp_path = document.getElementById('toeicsp-chart').getContext('2d');
    	window.myPie = new Chart(toeicsp_path, toeicsp_config);
    };    
</script>


@endsection	
