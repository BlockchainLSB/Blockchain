var express = require('express');
var router = express.Router();
var Client = require('node-rest-client').Client;
var client = new Client();
var temp;
/* json 파일 object 파일로 변환 */
var object = {};

var api_token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzgwMTI5NTAsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Ik9yZzEiLCJpYXQiOjE1Mzc5NzY5NTB9.Beub919ZD2RzUSG5y6W8kYfoIJ7U90QTrBGIjECo5qA";
var api_port = "4000";

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
		var result_json = JSON.parse(result);
		console.log('username : ' + result_json.Username);
		console.log('status code : ' + code);
		var projects = result_json.Projects;
		res.render('project/index', {login, projects});
	});
	
})

router.get('/detail', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var pnum = req.query.pnum;
	query_project('loadProjectdetail', ['pnum', pnum], function(data, statusCode){
		var result = data;
		var code = statusCode; 
		var result_json = JSON.parse(result);
		console.log('status code : ' + code);
		sess.project = result_json
		res.redirect('/project/description?pnum='+pnum);
	});
})

router.get('/description', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var token = sess.token;
	var project = sess.project;
	res.render('project/description', {login, project});
})

router.get('/affraise', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var token = sess.token;
	var project = sess.project;
	res.render('project/affraise', {login, project});
})


router.get('/contributes', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var token = sess.token;
	var project = sess.project;
	res.render('project/contributes', {login, project});
})

router.get('/contributors', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var token = sess.token;
	var project = sess.project;
	res.render('project/contributors', {login, project});
})


router.get('/addproject', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var token = sess.token;
	var user_id = sess.user_id;
	res.render('project/addproject', {login, user_id});
})

router.post('/addproject', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var token = sess.token;
	var pname = req.body.project_name;
	var pdes  = req.body.project_description;
	var contributors = req.body.contributor_list;
	invoke_project('addProject', ['token', token, 'pname', pname, 'pdes', pdes, 'contributors', contributors], function(data, statusCode){
		var result = data;
		var code = statusCode;

		console.log("result : " + result);
		console.log("status_code : " + code);
		res.redirect('/project?user_id='+sess.user_id);
	});
	
})



router.get('/accept', function(req, res, next){
	var sess = req.session;
	var token = sess.token;
	var pnum = req.query.pnum;
	console.log('accept token : ' + token);
	console.log('accept pnum : '+ pnum);
	
	invoke_project('acceptProject', ['token', token, 'pnum',pnum], function(statusCode){
		var code = statusCode;
		console.log("status_code : " + code);
		res.redirect('/project?user_id='+sess.user_id);
	});
})



module.exports = router;
