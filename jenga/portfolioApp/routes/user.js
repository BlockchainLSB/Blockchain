var express = require('express');
var router = express.Router();
var peer = 'peer';
var channel = 'mychannel';
var chaincode = 'mycc';
var api_host = 'http://localhost:4000';
var Client = require('node-rest-client').Client;
var client = new Client();
var temp;
/* json 파일 object 파일로 변환 */
var object = {};

var jsonheaders = {
					"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Mzc5MTA5MzEsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Ik9yZzEiLCJpYXQiOjE1Mzc4NzQ5MzJ9.4vsYJ7t0xKcEGnLrR4S_-Lbo-m1PzeQCsrBGoWm2VQQ",
					"Content-Type" : "application/json"
					};
object.headers = jsonheaders;


var invoke_user = function(fcn, args, callback){
	
	var api_url = 'http://localhost:4000/channels/mychannel/chaincodes/mycc';
	var jsonContent = {
						'peers' : ["peer0.org1.example.com","peer1.org1.example.com"],
						'fcn': fcn, 
						'args': args||[]
						};
	object.data = jsonContent;
	
	client.registerMethod("invokeUserMethod", api_url, "POST");
    client.methods.invokeUserMethod(object, function (data, response) {
		var statusCode = response.statusCode;
		callback(statusCode);
	});
}

var query_user = function(fcn, args, callback){
	var api_url = 'http://localhost:4000/channels/mychannel/chaincodes/mycc?peer=peer0.org1.example.com&fcn='+fcn+'&args='+JSON.stringify(args||null);

	
	client.registerMethod("queryUserMethod", api_url, "GET");
    client.methods.queryUserMethod(object, function (data, response) {
    	console.log("data : " + data );
		var statusCode = response.statusCode;
		callback(data, statusCode);
	});
}
router.get('/signup', function(req, res, next){
	res.render('user/signup', {});
})

router.post('/signup', function(req, res, next){
	var id = req.body.user_id;
	var passwd = req.body.user_passwd;
	
	query_user('searchUser', [id], function(data, statusCode){
		var result = data;
		var code = statusCode;
		var result_json = JSON.parse(result);
		console.log("result : " + result);
		console.log("status_code : " + code);
		if(!result_json.is_exist){
			res.render('user/signup', {error : '중복된 아이디 입니다.'});
		}
	});

	invoke_user('signup', ['id', id,'pw', passwd], function(statusCode){
		var code = statusCode;
		
		
		console.log("status_code : " + code);
		
		res.redirect('/');
	});
})

router.post('/signin', function(req, res, next){
	var id = req.body.user_id;
	var passwd = req.body.user_passwd;
	query_user('getToken', ['id' , id, 'pw', passwd], function(data, statusCode){
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
