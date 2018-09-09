@extends('layouts.project')

@section('project')

            <div class="card">
                <div class="card-header">Static</div>
                <div class="card-body">
                	<div class="row mt-3">
                	<div class="col-md-6">
                		<div class="card">
	                    	<div class="card-body">
	                    		<h6 class="card-title">
	                    			User1 기여도
	                    		</h6>
	                    		<div class="row">
	                    			<div class="col-md-6">
	                    				<div class="card">
	                    					<div class="card-body">
	                    						<h6 class="card-title">commit 수</h6>
	                    						<p class="card-text">12</p>
	                    					</div>
	                    				</div>
	                    			</div>
	                    			<div class="col-md-6">
	                    				<div class="card">
	                    					<div class="card-body">
	                    						<h6 class="card-title">팀원 평가</h6>
	                    						<p class="card-text">12</p>
	                    					</div>
	                    				</div>
	                    			</div>
	                    		</div>
	                    	</div>
	                    </div>
	                  	
                	</div>
                	<div class="col-md-6">
                		<div class="card">
	                    	<div class="card-body">
	                    		<h6 class="card-title">
	                    			User2 기여도 
	                    		</h6>
	                    		<div class="row">
	                    			<div class="col-md-6">
	                    				<div class="card">
	                    					<div class="card-body">
	                    						<h6 class="card-title">commit 수</h6>
	                    						<p class="card-text">12</p>
	                    					</div>
	                    				</div>
	                    			</div>
	                    			<div class="col-md-6">
	                    				<div class="card">
	                    					<div class="card-body">
	                    						<h6 class="card-title">팀원 평가</h6>
	                    						<p class="card-text">12</p>
	                    					</div>
	                    				</div>
	                    			</div>
	                    		</div>
	                    	</div>
	                    </div>
	                  	
                	</div>
                </div>
                    
                    <div class="row mt-3">
                	<div class="col-md-6">
                		<div class="card">
	                    	<div class="card-body">
	                    		<h6 class="card-title">
	                    			User1 Commit 수
	                    		</h6>
	                    		<canvas id="chart-user1-commit"></canvas>
	                    	</div>
	                    </div>
	                  	<script>
	                        var config_user1_commit = {
	                            type: 'line',
	                            data: {
	                                datasets: [{
	                                    data: [
	                                        12, 33, 44, 22, 33, 44, 22
	                                    ],
	                                    label: 'Commit'
	                                }],
	                                labels: [
	                                    'January',
	                                    'February',
	                                    'March',
	                                    'April',
	                                    'May',
	                                    'June',
	                                    'July',
	                                ]
	                            },
	                            options: {
	                                responsive: true
	                            }
	                        };
	
	                    </script>
                	</div>
                	<div class="col-md-6">
                		<div class="card">
	                    	<div class="card-body">
	                    		<h6 class="card-title">
	                    			User2 Commit 수 
	                    		</h6>
	                    		<canvas id="chart-user2-commit"></canvas>
	                    	</div>
	                    </div>
	                  	<script>
	                        var config_user2_commit = {
	                            type: 'line',
	                            data: {
	                                datasets: [{
	                                    data: [
	                                        12, 33, 44, 22, 33, 44, 22
	                                    ],
	                                    label: 'Commit'
	                                }],
	                                labels: [
	                                    'January',
	                                    'February',
	                                    'March',
	                                    'April',
	                                    'May',
	                                    'June',
	                                    'July',
	                                ]
	                            },
	                            options: {
	                                responsive: true
	                            }
	                        };
	
	                    </script>
                	</div>
                </div>
                
                <div class="card mt-3">
                    	<div class="card-body">
                    		<h6 class="card-title">
                    			전체 Commit 수
                    		</h6>
                    		<canvas id="chart-all-commit"></canvas>
                    	</div>
                    </div>
                  	<script>
                        var config_all_commit = {
                            type: 'line',
                            data: {
                                datasets: [{
                                    data: [
                                        12, 33, 44, 22, 33, 44, 22
                                    ],
                                    label: 'Commit'
                                }],
                                labels: [
                                    'January',
                                    'February',
                                    'March',
                                    'April',
                                    'May',
                                    'June',
                                    'July',
                                ]
                            },
                            options: {
                                responsive: true
                            }
                        };

                    </script>
                </div>
            </div>
            <script>
            window.onload = function() {
                var ctx_all_commit = document.getElementById('chart-all-commit').getContext('2d');
                window.myPie = new Chart(ctx_all_commit, config_all_commit);
                
                var ctx_user1_commit = document.getElementById('chart-user1-commit').getContext('2d');
                window.myPie = new Chart(ctx_user1_commit, config_user1_commit);
                
                var ctx_user2_commit = document.getElementById('chart-user2-commit').getContext('2d');
                window.myPie = new Chart(ctx_user2_commit, config_user2_commit);
            };

        </script>
@endsection