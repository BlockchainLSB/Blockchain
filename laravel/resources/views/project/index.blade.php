@extends('layouts.master')

@section('content')
<div class="container pl-7 pr-7">
	<a class="btn btn-secondary mb-3" href="/addproject">Add Project</a>
    <div class="row justify-content-center pr-6 pl-6">
        @for($i=0; $i<6; $i++)
        <div class="col-md-4 mb-4">
            <div class="card">
                <div class="card-header">
                	<a href="/project/repository/description">Name</a>
                </div>

                <div class="card-body">
                    
                    Description
                </div>
                <small class="card-footer text-muted">
                	language
                </small>
            </div>
        </div>
        @endfor
   </div>
        
    
</div>
@endsection
