var express = require('express');
var router = express.Router();
var Client = require('node-rest-client').Client;
var client = new Client();
var temp;
/* json 파일 object 파일로 변환 */
var object = {};

var api_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Mzc5NzgwMzQsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Ik9yZzEiLCJpYXQiOjE1Mzc5NDIwMzR9.GEqG7hFWyQTQVVlLUUGnDYmkQknNqSwKpE-AkaUX2_4";
var api_port = "4000";

var jsonheaders = {
					"Authorization": "Bearer " + api_token,
					"Content-Type" : "application/json"
					};
object.headers = jsonheaders;


var invoke_user = function(fcn, args, callback){
	
	var api_url = 'http://52.79.245.63:'+api_port+'/channels/mychannel/chaincodes/mycc'; 
	var jsonContent = {
						'peers' : ["peer0.org1.example.com","peer1.org1.example.com"],
						'fcn': fcn, 
						'args': args||[]
						};
	object.data = jsonContent;
	
	client.registerMethod("invokeUserMethod", api_url, "POST");
    client.methods.invokeUserMethod(object, function (data, response) {
    	var buf = new Buffer(data);
    	result = buf.toString('utf-8');
		var statusCode = response.statusCode;
		console.log('tx_id : ' + result);
		callback(statusCode);
	});
}

var query_user = function(fcn, args, callback){
	var api_url = 'http://52.79.245.63:'+api_port+'/channels/mychannel/chaincodes/mycc?peer=peer0.org1.example.com&fcn='+fcn+'&args='+JSON.stringify(args||null);

	
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
	
	query_user('searchUser', ['id', id], function(data, statusCode){
		var result = data;
		var code = statusCode; 
		//var result_json = JSON.parse(result);
		if(result.indexOf('Error') != -1){ 
			invoke_user('signup', ['id', id,'pw', passwd], function(statusCode){
				var code = statusCode;
				console.log("search user status_code : " + code);
				console.log(data);
				res.redirect('/');
			});
		}else{
			console.log("search user code : " + code);
			console.log('존재하는 id');
			res.render('user/signup', {error : '존재하는 id'});
		}
	});
/*
	*/
})

router.post('/signin', function(req, res, next){
	var id = req.body.user_id;
	var passwd = req.body.user_passwd;
	invoke_user('signin', ['id', id,'pw', passwd], function(statusCode){
		var code = statusCode;
		console.log("sign in status_code : " + code);
		
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
	});
	
	
})
module.exports = router;
