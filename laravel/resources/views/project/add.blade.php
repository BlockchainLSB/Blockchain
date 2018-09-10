@extends('layouts.master')

@section('content')
<div class="container pl-7 pr-7">
    <div class="col-md-12">
        <div class="card">
        	<div class="card-header">
        		Add Project
        	</div>
        	<div class="card-body">
        		<form action="/project">
				  <div class="form-group">
				    <label for="project_name">Project Name</label>
				    <input type="text" class="form-control" id="project_name"  placeholder="Project Name">
				  </div>
				  <div class="form-group">
				  	<label for="project_description">Project Description</label>
				    <input type="text" class="form-control" id="project_desciption"  placeholder="Project Description">
				  </div>
				  <div class="form-group">
				  	<label for="project_contributor">Contributor</label>
					  <div class="input-group mb-3">
						  <input type="text" class="form-control" placeholder="Contributor">
						  <div class="input-group-append">
						    <button class="btn btn-secondary" type="button">찾기</button>
						  </div>
						</div>
				  </div>
				  <button type="submit" class="btn btn-secondary">추가하기</button>
				</form>
        	</div>
        </div>
   </div>
        
    
</div>
@endsection
