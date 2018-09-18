var express = require('express');
var router = express.Router();
var peer = 'peer';
var channel = 'myc';
var chaincode = 'lye';
var api_host = '172.21.0.3:7053';
var Client = require('node-rest-client').Client;
var client = new Client();
var temp;
/* json 파일 object 파일로 변환 */
var object = {};

var jsonheaders = {"Content-Type" : "application/json"};
object.headers = jsonheaders;


var invoke_user = function(fcn, args, callback){
	
	var api_url = api_host + '/channels/' + channel + '/chaincodes/'+chaincode;
	var jsonContent = {
						'peers' : peer,
						'fcn': fcn, 
						'args': args||[]
						};
	object.data = jsonContent;
	
	client.registerMethod("invokeUserMethod", api_url, "POST");
    client.methods.invokeUserMethod(object, function (data, response) {
		buf = new Buffer(JSON.stringify(data));

		var result =  buf.toString();
		var statusCode = response.statusCode;
		callback(result, statusCode);
	});
}

var query_user = function(fcn, args){
	var api_url = "https://eyecan.tk/rest_api/test_api";
	var jsonContent = {
						'device_id' : '1234',
						};
	object.data = jsonContent;

	client.registerMethod("queryUserMethod", api_url, "GET");
	client.methods.queryUserMethod(object, function (data, response) {
		buf = new Buffer(JSON.stringify(data));

		var result =  buf.toString();
		var statusCode = response.statusCode;
		callback(result, statusCode);
	});
}
router.get('/signup', function(req, res, next){
	res.render('user/signup', {});
})

router.post('/signup', function(req, res, next){
	var id = req.body.user_id;
	var passwd = req.body.user_passwd;
	/*
	invoke_user('searchUser', [id], function(data, statusCode){
		var result = data;
		var code = statusCode;
		var result_json = JSON.parse(result);
		console.log("result : " + result);
		console.log("status_code : " + code);
		if(!result_json.is_exist){
			res.render('user/signup', {error : '중복된 아이디 입니다.'});
		}
	});*/

	invoke_user('signup', ['id', id,'passwd', passwd], function(data, statusCode){
		var result = data;
		var code = statusCode;
		var result_json = JSON.parse(result);
		console.log("result : " + result);
		console.log("status_code : " + code);
		
		res.redirect('/');
	});
})

router.post('/signin', function(req, res, next){
	var id = req.body.user_id;
	var passwd = req.body.user_passwd;
	invoke_user('signin', [id, passwd], function(data, statusCode){
		var result = data;
		var code = statusCode;
		var result_json = JSON.parse(result);
		var token = result_json.token;

		console.log("result : " + result);
		console.log("status_code : " + code);
		console.log("token : " + token);
		var sess = req.session;
		sess.token= token;
		sess.login = true;
		res.redirect('/project?user_id='+id);
	});
	
})
module.exports = router;
