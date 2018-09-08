@extends('layouts.master')

@section('content')
<div class="container">
	<div class="card mb-4" style="background:#343A40">
	    <div class="card-body">
	        <h5 class="d-inline" style="color:#FFFFFF"> Project Name</h5>
	    </div>
	</div>
	
    <div class="row">
        <div class="col-md-2">
                <ul class="list-group">
				  <li class="list-group-item dropdown"><a class="dropdown-toggle" data-toggle="dropdown" href="#" style="color:#000000">Repository <span class="caret"></span></a>
			        <ul class="dropdown-menu">
			          <li class="list-group-item"><a href="/project/repository/description" style="color:#000000">Description</a></li>
			          <li class="list-group-item"><a href="/project/repository/commit" style="color:#000000">Commit</a></li>
			          <li class="list-group-item"><a href="/project/repository/contributor" style="color:#000000">Contributor</a></li>
			        </ul>
			      </li>
				  <li class="list-group-item">
				  	<a href="/static" style="color:#000000">Static</a>
				  </li>
				  <li class="list-group-item">
				  	<a href="/evaluation" style="color:#000000">Evaluation</a>
				  </li>
				</ul>
				
            </div>
        <div class="col-md-10 mb-5">
                 @yield('project')
            
        </div>
    </div>
    
</div>
@endsection
