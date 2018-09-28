var express = require('express');
var router = express.Router();
var Client = require('node-rest-client').Client;
var client = new Client();
/* json 파일 object 파일로 변환 */
var object = {};




var invoke_project = function(api_token, api_port, fcn, args, callback){
	var jsonheaders = {
		"Authorization": "Bearer " + api_token,
		"Content-Type" : "application/json"
		};
	object.headers = jsonheaders;

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

var query_project = function(api_token, api_port, fcn, args, callback){ 
	var jsonheaders = {
		"Authorization": "Bearer " + api_token,
		"Content-Type" : "application/json"
		};
	object.headers = jsonheaders;

	var api_url = 'http://52.79.245.63:'+api_port+'/channels/mychannel/chaincodes/mycc?peer=peer0.org1.example.com&fcn='+fcn+'&args='+JSON.stringify(args||null);
	
	
	client.registerMethod("queryProjectMethod", api_url, "GET");
    client.methods.queryProjectMethod(object, function (data, response) {
    	console.log("data : " + data );
		var statusCode = response.statusCode;
		callback(data, statusCode);
	});
}

router.get('/', function(req, res, next){
	console.log('user_id : ' + req.body.user_id);
	var sess = req.session;
	console.log('sess.token : ' + sess.token);
	console.log('sess.login : ' + sess.login);
	var token = sess.token;
	var login = sess.login;
	var user_id = sess.user_id;
	var api_token = sess.api_token;
	var api_port = sess.api_port;
	var search_id = req.query.user_id;
	query_project(api_token, api_port, 'searchProject', ['userid', search_id], function(data, statusCode){
		var result = data;
		var code = statusCode; 
		var result_json = JSON.parse(result);
		console.log('username : ' + result_json.Username);
		console.log('status code : ' + code);
		var projects = result_json.Projects;
		res.render('project/index', {login, user_id, projects, api_token, api_port, search_id});
	});
	
})

router.get('/detail', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var pnum = req.query.pnum;
	var api_token = sess.api_token;
	var api_port = sess.api_port;
	query_project(api_token, api_port, 'loadProjectdetail', ['pnum', pnum], function(data, statusCode){
		var result = data;
		var code = statusCode; 
		var result_json = JSON.parse(result);
		console.log('status code : ' + code);
		sess.project = result_json;
		res.redirect('/project/description?pnum='+pnum);
	});
})

router.get('/description', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var token = sess.token;
	var api_token = sess.api_token;
	var api_port = sess.api_port;
	var project = sess.project;
	var user_id = sess.user_id;
	res.render('project/description', {login, user_id, project, api_token, api_port});
})

router.get('/appraise', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var token = sess.token;
	var project = sess.project;
	var pnum = project.Pnum;
	var api_token = sess.api_token;
	var api_port = sess.api_port;
	var user_id = sess.user_id;
	query_project(api_token, api_port, 'requestedConlist', ['token', token, 'pnum', ''+pnum], function(data, statusCode){
		console.log('pnum : '+ pnum);
		var result = data;
		var code = statusCode; 
		var request_list = JSON.parse(result);
		console.log('status code : ' + code);
		res.render('project/appraise', {login, user_id, request_list, project, api_token, api_port});
	});
	
})

router.post('/affraise', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var token = sess.token;
	var project = sess.project;
	var pnum = project.Pnum;
	var pdes = req.body.contribution
	var api_token = sess.api_token;
	var api_port = sess.api_port;

	invoke_project(api_token, api_port, 'addContribution', ['token', token, 'pnum', ''+pnum, 'pdes', pdes], function(statusCode){
		console.log('pnum : '+ pnum);
		var code = statusCode;

		console.log("status_code : " + code);
		res.redirect('/project/appraise?pnum='+pnum);
	})
})


router.get('/contributes', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var token = sess.token;
	var project = sess.project;
	var pnum = project.Pnum;
	var api_token = sess.api_token;
	var api_port = sess.api_port;
	var user_id = sess.user_id;

	query_project(api_token, api_port, 'allacceptedConlist', ['pnum', ''+pnum], function(data, statusCode){
		console.log('pnum : '+ pnum);
		var result = data;
		var code = statusCode; 
		var accepted_list = JSON.parse(result);
		console.log('status code : ' + code);
		res.render('project/contributes', {login, user_id, accepted_list ,project, api_token, api_port});
	});
})

router.get('/contributors', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var token = sess.token;
	var project = sess.project;
	var user_id = sess.user_id;
	var api_token = sess.api_token;
	var api_port = sess.api_port;
	res.render('project/contributors', {login, user_id, project, api_token, api_port});
})


router.get('/addproject', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var token = sess.token;
	var user_id = sess.user_id;
	var api_token = sess.api_token;
	var api_port = sess.api_port;
	res.render('project/addproject', {login, user_id, api_token, api_port});
})

router.post('/addproject', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var token = sess.token;
	var pname = req.body.project_name;
	var pdes  = req.body.project_description;
	var contributors = req.body.contributor_list;
	var api_token = sess.api_token;
	var api_port = sess.api_port;
	invoke_project(api_token, api_port, 'addProject', ['token', token, 'pname', pname, 'pdes', pdes, 'contributors', contributors], function(data, statusCode){
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
	var api_token = sess.api_token;
	var api_port = sess.api_port;

	invoke_project(api_token, api_port, 'acceptProject', ['token', token, 'pnum',pnum], function(statusCode){
		var code = statusCode;
		console.log("status_code : " + code);
		res.redirect('/project?user_id='+sess.user_id);
	});
})

router.get('/accept_contribution', function(req, res, next){
	var sess = req.session;
	var token = sess.token;
	var project = sess.project;
	var pnum = project.Pnum;
	var pindex = req.query.pindex;
	var api_token = sess.api_token;
	var api_port = sess.api_port;
	
	invoke_project(api_token, api_port, 'acceptContribution', ['token', token, 'pnum',''+pnum, 'pindex', ''+pindex], function(statusCode){
		console.log('pnum : '+ pnum);
		var code = statusCode;
		console.log("status_code : " + code);
		res.redirect('/project/appraise?pnum='+pnum);
	});
})


module.exports = router;
