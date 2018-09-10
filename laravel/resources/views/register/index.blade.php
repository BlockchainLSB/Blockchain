@extends('layouts.master')

@section('content')
<div class="container pl-7 pr-7">
    <div class="col-md-12">
        <div class="card">
        	<div class="card-header">
        		Register User
        	</div>
        	<div class="card-body">
        		<form action="/project">
        			<div class="form-group">
        				<div class="row">
        					<div class="col-md-2">
	        					<div class="card">
	        						<div class="card-body">
	        							<img src="/image/profile_img.png" style="width: 100px; height: auto"/>
	        						</div>
	        					</div>
        					</div>
        				</div>
        				<div class="row mt-3">
        					<div class="col-md-5">
        						<div class="input-group mb-3">
							  <div class="input-group-prepend">
							    <span class="input-group-text">Profile</span>
							  </div>
							  <div class="custom-file">
							    <input type="file" class="custom-file-input" id="profile_imh">
							    <label class="custom-file-label" for="profile_img">Upload Profile Image(.jpg)</label>
							  </div>
							</div>
        					</div>
        				</div>
        			</div>
				  <div class="form-group">
				  	<div class="row">
				  		<div class="col-md-5">
				  			<label for="name">Name</label>
				    		<input type="text" class="form-control" id="name"  placeholder="Name">
				  		</div>
				  		<div class="col-md-5">
						  	<label for="email">ID</label>
						    <input type="text" class="form-control" id="id"  placeholder="ID">
				    	</div>
				  	</div>
				  </div>
				  <div class="form-group">
				  	<div class="row">
				  		<div class="col-md-5">
				  			<label for="email">Email</label>
				    		<input type="email" class="form-control" id="email"  placeholder="Email">
				  		</div>
				  	</div>
				  </div>
				  <div class="form-group">
				  	<div class="row">
				  		<div class="col-md-5">
				  			<label for="password">Password</label>
				    		<input type="password" class="form-control" id="password"  placeholder="Password">
				  		</div>
				  		<div class="col-md-5">
						  	<label for="password_confirm">Password Confirm</label>
				    		<input type="password" class="form-control" id="password_confirm"  placeholder="Password Confirm">
				    	</div>
				  	</div>
				  </div>
				   <div class="form-group">
				  	<div class="row">
				  		<div class="col-md-5">
				  			<label for="school">School</label>
				    		<input type="text" class="form-control" id="school"  placeholder="School">
				  		</div>
				  		<div class="col-md-5">
						  	<label for="major">Major</label>
				    		<input type="text" class="form-control" id="major"  placeholder="Major">
				    	</div>
				  	</div>
				  </div>
				  <div class="form-group">
				  	<div class="row">
				  		<div class="col-md-10">
				  			<div class="card">
				  				<div class="card-header">Git 연동 ,자격증 점수 입력</div>
				  				<div class="card-body">
				  					<button class="btn btn-secondary mb-4">Git 연동하기</button>
				  					<div class="row">
				  					<div class="col-md-4">
				  					<label for="toeic">Toeic</label>
				    				<input type="text" class="form-control" id="toeic"  placeholder="Toeic">
				    				</div>
				    				<div class="col-md-4">
				    				<label for="toeic_speaking">Toeic Speaking</label>
				    				<input type="text" class="form-control" id="toeic_speaking"  placeholder="Toeic Speaking">
				    				</div>
				    				<div class="col-md-4">
				    				<label for="topcit">Topcit</label>
				    				<input type="text" class="form-control" id="topcit"  placeholder="Topcit">
				    				</div>
				    				</div>
				  				</div>
				  			</div>
				  		</div>
				  	</div>
				  </div>
				  <button type="submit" class="btn btn-primary">가입하기</button>
				</form>
        	</div>
        </div>
        
    
</div>
@endsection