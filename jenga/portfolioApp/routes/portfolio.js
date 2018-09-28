var express = require('express');
var router = express.Router();
var peer = 'peer';
var channel = 'mychannel';
var chaincode = 'mycc';
var port = '4000'
var api_host = 'http://52.79.245.63:' + port;
var Client = require('node-rest-client').Client;
var client = new Client();
var temp;
/* json 파일 object 파일로 변환 */
var object = {};


var query_portfolio = function(api_token, api_port, fcn, args, callback){
	var jsonheaders = {
					"Authorization": "Bearer " + api_token,
					"Content-Type" : "application/json"
					};
	object.headers = jsonheaders;
	var api_host = 'http://52.79.245.63:' + api_port;
	var api_url = api_host + '/channels/mychannel/chaincodes/mycc?peer=peer0.org1.example.com&fcn='+fcn+'&args='+JSON.stringify(args||null);

	client.registerMethod("queryUserMethod", api_url, "GET");
    client.methods.queryUserMethod(object, function (data, response) {
    	console.log("data : " + data );
		var statusCode = response.statusCode;
		callback(data, statusCode);
	});
}

router.get('/', function(req, res, next){
	var sess = req.session;
	var login = sess.login;
	var user_id = sess.user_id;
	var api_token = sess.api_token;
	var api_port = sess.api_port;
	console.log("userid: " + user_id)
	query_portfolio(api_token, api_port, 'searchPortfolio', ['userid', user_id], function(data, statusCode){
		var result = data;
		var code = statusCode;
		var jportfolio = JSON.parse(result);
		
		query_portfolio(api_token, api_port, 'searchProject', ['userid', user_id], function(data, statusCode){
		var result = data;
		var code = statusCode; 
		var result_json = JSON.parse(result);
		var projectl = result_json.Projects;
		var cn;
		if(projectl == null) {
			console.log("no projects")
			cn = 0;
		}
		else {
			cn = projectl.length;
		}
		
		console.log('username : ' + result_json.Username);
		console.log('status code : ' + code);
		
		res.render('portfolio/index', {login, user_id, jportfolio, cn, projectl, api_token, api_port});
	});
	
			
	});
})

module.exports = router;