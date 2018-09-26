var express = require('express');
var router = express.Router();
var Client = require('node-rest-client').Client;
var client = new Client();
var temp;
/* json 파일 object 파일로 변환 */
var object = {};

var api_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Mzc5NzgwMzQsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Ik9yZzEiLCJpYXQiOjE1Mzc5NDIwMzR9.GEqG7hFWyQTQVVlLUUGnDYmkQknNqSwKpE-AkaUX2_4";
var api_port = "4001";

var jsonheaders = {
					"Authorization": "Bearer " + api_token,
					"Content-Type" : "application/json"
					};
object.headers = jsonheaders;


var invoke_project = function(fcn, args, callback){
	
	var api_url = 'http://52.79.245.63:'+api_port+'/channels/mychannel/chaincodes/mycc';
	var jsonContent = {
						'peers' : ["peer0.org1.example.com","peer1.org1.example.com"],
						'fcn': fcn, 
						'args': args||[]
						};
	object.data = jsonContent;
	
	client.registerMethod("invokeProjectMethod", api_url, "POST");
    client.methods.invokeProjectMethod(object, function (data, response) {
    	var buf = new Buffer(data);
    	result = buf.toString('utf-8');
		var statusCode = response.statusCode;
		console.log('tx_id : ' + result);
		callback(statusCode);
	});
}

var query_project = function(fcn, args, callback){ 
	var api_url = 'http://52.79.245.63:'+api_port+'/channels/mychannel/chaincodes/mycc?peer=peer0.org1.example.com&fcn='+fcn+'&args='+JSON.stringify(args||null);
	
	
	client.registerMethod("queryProjectMethod", api_url, "GET");
    client.methods.queryProjectMethod(object, function (data, response) {
    	console.log("data : " + data );
		var statusCode = response.statusCode;
		callback(data, statusCode);
	});
}

router.get('/', function(req, res, next){
	var sess = req.session;
	console.log('sess.token : ' + sess.token);
	console.log('sess.login : ' + sess.login);
	var token = sess.token;
	var login = sess.login;
	query_project('loadProject', ['token', token], function(data, statusCode){
		var result = data;
		var code = statusCode; 
		//var result_json = JSON.parse(result);
		res.render('project/index', {login});
	});
	
})

router.post('/', function(req, res, next){
	var sess = req.session;
	var token = sess.token;
	var project_name = req.body.project_name;
	var project_description = req.body.project_description;
	invoke_project('addProject', [token, project_name, project_description], function(data, statusCode){
		var result = data;
		var code = statusCode;
		var result_json = JSON.parse(result);

		console.log("result : " + result);
		console.log("status_code : " + code);
		res.redirect('/project');
	});
})

router.get('/repository/description', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var token = sess.token;
	res.render('project/repository/description', {login});
})

router.get('/repository/commit', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var token = sess.token;
	res.render('project/repository/commit', {login});
})

router.get('/repository/contributor', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var token = sess.token;
	res.render('project/repository/contributor', {login});
})

router.get('/static', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var token = sess.token;
	res.render('project/static', {login});
})

router.get('/addproject', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var token = sess.token;
	res.render('project/addproject', {login});
})

router.post('/addproject', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var token = sess.token;
	var pname = req.body.project_name;
	var pdes  = req.body.project_description;
	res.redirect('/project');
})


router.get('/evaluation', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var token = sess.token;
	res.render('project/evaluation', {login});
})



module.exports = router;
