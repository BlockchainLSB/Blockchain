@extends('layouts.master')

@section('content')

<div class="container">
	<!--
  	<div class="row py-3 flex-items-sm-center justify-content-center">
		<div class="col-md-6 clearfix">
		  <div class="card profile-card">
		    <figure>
		      <img src="http://cps-static.rovicorp.com/3/JPG_400/MI0003/711/MI0003711195.jpg?partner=allrovi.com" class="img-profile">
		    </figure>
		    <div class="card-body text-center">
		      <p class="h4 card-title font-weight-bold">김블록</p>
		      <p class="h7 card-subtitle text-muted">부산대학교</p>
		      <p class="h7 card-subtitle text-muted">정보컴퓨터공학</p>
		      <p class="h7 card-subtitle text-muted">hyper@pusan.ac.kr</p>
		    </div>
		  </div>
		</div>
	  </div>
	 -->
	 
	 <div class="col-md-8 clearfix mx-auto">
	 	<div class="card shadow">
	 		<div class="card-header"><h3>Profile</h3></div>
	 		<img src="/image/profile_happy.jpg" class="img-profile img-fluid z-depth-4">
	 		<div class="card-body text-center">
	 			<p class="h2 card-title font-weight-bold">김블록</p>
	 			<p class="h6 card-text text-muted">부산대학교</p>
		      	<p class="h6 card-text text-muted">정보컴퓨터공학</p>
		      	<p class="h6 card-text text-muted">hyper@pusan.ac.kr</p>
	 		</div>
	 	</div>
	 </div>
	 <!--
	 <div class="col-md-6 clearfix lang-freq">
	 	<div class="card shadow">
	 		<div class="card-body text-center">
	 			<p class="h4 card-title font-weight-bold">언어 사용빈도</p>
	 			<canvas id="lang-chart"></canvas>
	 		</div>
	 	</div>
	 </div>
	 -->
</div>

@endsection