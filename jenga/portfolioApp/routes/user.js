var express = require('express');
var request = require('request');
var router = express.Router();
var peer = 'peer';
var channel = 'common';
var api_host = '127.0.0.1';

var invoke_user = function(fcn, args){
	var api_url = api_host + '/channels/' + channel + '/channelcodes/user';
	var options = {
		url: api_url,
		form: {	
			'peers' : peer,
			'fcn': fcn, 
			'args': JSON.stringify(args||null) },
		headers: { 
			'Accept': 'plan/text',
			'Content-Type': 'application/json; charset=UTF-8',
   			'Accept-Language': 'ko' },
	};
	var _req = request.post(options).on('response', function(response) {
		console.log(response.statusCode); // 200
		console.log(response.headers['content-type']);
	});
}

var query_user = function(fcn, args){
	var api_url = api_host + '/channels/' + channel + '/channelcodes/user';
	var options = {
		url: api_url,
		form: {	
			'peers' : peer,
			'fcn': fcn, 
			'args': JSON.stringify(args||null) },
		headers: { 
			'Accept': 'plan/text',
			'Content-Type': 'application/json; charset=UTF-8',
   			'Accept-Language': 'ko' },
	};
	var _req = request.get(options).on('response', function(response) {
		console.log(response.statusCode); // 200
		console.log(response.headers['content-type']);
	});
}
router.get('/signup', function(req, res, next){
	res.render('user/signup', {});
})

router.post('/signup', function(req, res, next){
	var id = req.body.user_id;
	var passwd = req.body.user_passwd;
	//invoke_user('signup', [id, passwd]);
	res.redirect('/');
})

module.exports = router;
