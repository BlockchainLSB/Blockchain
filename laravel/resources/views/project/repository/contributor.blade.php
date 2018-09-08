@extends('layouts.repository')

@section('repository_content')

    <div class="card">
        <div class="card-header">Contributor</div>
        <div class="card-body">
	        <div class="row">
	        	@for($i = 0; $i < 3; $i++)
	        	<div class="col-md-3">
	        		<div class="card">
	        			<div class="card-body">
	        				<img src="/image/profile_happy.jpg" class="img-contributor"/>
	        			</div>
	        			<div class="card-body">
	        				<h6 class="card-title" style="text-align: center">Happy</h6>
	        				<p class="card-text" style="text-align: center">aaa@google.com</p>
	        			</div>
	        		</div>
	        	</div>
	        	@endfor
	        </div> 
        </div>
        
    </div>
   
@endsection
