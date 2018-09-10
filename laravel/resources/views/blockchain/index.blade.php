@extends('layouts.master')

@section('content')
<div class="container">
	<div class="col-md-9 clearfix mx-auto pt-4 pb-4 h-300">
		<div class="card shadow">
			<div class="card-header"><h3>Transactions</h3></div>
			<div class="card-body text-center">
				<table class="table table-bordered table-sm table-hover">
				    <thead>
				      <tr>
				        <th>No.</th>
				        <th>name</th>
				        <th>time</th>
				      </tr>
				    </thead>
				    <tbody>
				      <div >
	                    @for($i = 1; $i < 6; $i++)
	                    <tr>
	                        <td>{{$i}}</td>
	                        <td>sign in {{$i}}</td>
	                        <td>{{$i * 2}}</td>
	                    </tr>
	                    @endfor
	                    </div>
				    </tbody>
				  </table>
				  	
			  		<ul class="pagination">
					    <li class="page-item">
					      <a class="page-link" href="#" aria-label="Previous">
					        <span aria-hidden="true">&laquo;</span>
					        <span class="sr-only">Previous</span>
					      </a>
					    </li>
					    <li class="page-item"><a class="page-link" href="#">1</a></li>
					    <li class="page-item"><a class="page-link" href="#">2</a></li>
					    <li class="page-item"><a class="page-link" href="#">3</a></li>
					    <li class="page-item">
					      <a class="page-link" href="#" aria-label="Next">
					        <span aria-hidden="true">&raquo;</span>
					        <span class="sr-only">Next</span>
					      </a>
					    </li>
				  	</ul>
			</div>
		</div>
	</div>
	 	
<div class="col-md-9 clearfix mx-auto pt-4 pb-4 h-300">
		<div class="card shadow">
			<div class="card-header"><h3>Blocks</h3></div>
			<div class="card-body text-center">
				<table class="table table-bordered table-sm table-hover">
				    <thead>
				      <tr>
				        <th>No.</th>
				        <th>name</th>
				        <th>time</th>
				      </tr>
				    </thead>
				    <tbody>
				      <div >
	                    @for($i = 1; $i < 6; $i++)
	                    <tr>
	                        <td>{{$i}}</td>
	                        <td>sign in {{$i}}</td>
	                        <td>{{$i * 2}}</td>
	                    </tr>
	                    @endfor
	                    </div>
				    </tbody>
				  </table>
				  	
			  		<ul class="pagination">
					    <li class="page-item">
					      <a class="page-link" href="#" aria-label="Previous">
					        <span aria-hidden="true">&laquo;</span>
					        <span class="sr-only">Previous</span>
					      </a>
					    </li>
					    <li class="page-item"><a class="page-link" href="#">1</a></li>
					    <li class="page-item"><a class="page-link" href="#">2</a></li>
					    <li class="page-item"><a class="page-link" href="#">3</a></li>
					    <li class="page-item">
					      <a class="page-link" href="#" aria-label="Next">
					        <span aria-hidden="true">&raquo;</span>
					        <span class="sr-only">Next</span>
					      </a>
					    </li>
				  	</ul>
			</div>
		</div>
	</div>
</div>

@endsection	