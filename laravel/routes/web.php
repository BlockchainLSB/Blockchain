<?php

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| contains the "web" middleware group. Now create something great!
|
*/

Route::get('/', function () {
    return view('welcome');
});

Route::get('/portfolio', function () {
	return view('portfolio/index');
});

Route::get('/project', function(){
	return view('project/index');
});

Route::get('/project/repository/description', function(){
	return view('project/repository/descriptionx');
});

Route::get('/project/repository/commit', function(){
	return view('project/repository/commit');
});

Route::get('/project/repository/contributor', function(){
	return view('project/repository/contributor');
});

Route::get('/static', function () {
    return view('/project/static');
});


Route::get('/evaluation', function () {
    return view('/project/evaluation');
});

Route::get('/addproject', function () {
    return view('/project/add');
});

Route::get('/register', function () {
    return view('/register.index');
});

