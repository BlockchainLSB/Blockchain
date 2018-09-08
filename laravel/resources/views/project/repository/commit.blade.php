@extends('layouts.repository')

@section('repository_content')
    <div class="card">
        <div class="card-header">Commit</div>
        <div class="card-body">
	        <div class="row" style="margin-left : 0px; margin-right : 0px;">   
	            <form class="form-inline">  
	                <div>
	                    <select class="form-group form-control">
	                        <option>branch : all</option>
	                        <option>branch : master</option>
	                        <option>branch : develop</option>
	                    </select>
	                </div> 
	                <button type="submit" class="select-category btn btn-outline-secondary ml-2">검색 <i class="fa fa-search"></i></button>	        
	            </form> 
	        </div>
        </div>
        <div class="card mb-4 mr-2 ml-2">
	    	<div class="card-body" style="width:100%;">
                <table class="table table-hover">
		            <thead style="background:#343A40;color:#FFFFFF">
		                <tr>
		                    <th scope="col">No.</th>
		                    <th scope="col">Branch</th>
		                    <th scope="col">User</th>
		                    <th scope="col">Message</th>
		                    <th scope="col">Time</th>
		                </tr>
		            </thead>
		            <tbody >
		                <div >
		                    @for($i = 1; $i < 11; $i++)
		                    <tr>
		                        <td><a href="#">commit#</a></td>
		                        <td>branch</td> 
		                        <td>user name</td>
		                        <td>commit message</td>
		                        <td>0000</td>
		                    </tr>
		                    @endfor
		                    </div>
		            </tbody>
		        </table>
	        </div>
	        
	        <div class="card-body">
	        	<nav>
				  <ul class="pagination">
				    <li class="page-item">
				      <a class="page-link" href="#">
				        <span aria-hidden="true">&laquo;</span>
				        <span class="sr-only">Previous</span>
				      </a>
				    </li>
				    <li class="page-item active"><a class="page-link" href="#">1</a></li>
				    <li class="page-item"><a class="page-link" href="#">2</a></li>
				    <li class="page-item"><a class="page-link" href="#">3</a></li>
				    <li class="page-item">
				      <a class="page-link" href="#" >
				        <span aria-hidden="true">&raquo;</span>
				        <span class="sr-only">Next</span>
				      </a>
				    </li>
				  </ul>
				</nav>
	        </div>
        </div>
    </div>
@endsection
