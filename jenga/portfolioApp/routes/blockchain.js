var express = require('express');
var router = express.Router();
var peer = 'peer';
var channel = 'mychannel';
var chaincode = 'mycc';
var api_host = 'http://52.79.245.63:4001';
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


var invoke_portfolio = function(fcn, args, callback){
	
	var api_url = api_host + '/channels/mychannel/chaincodes/mycc';
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

var query_portfolio = function(fcn, args, callback){
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
	var token = sess.token;
	console.log('user token : ' + token);
	
	query_portfolio('getUserTransaction', ['token', token], function(data, statusCode){
		var result = data;
		var code = statusCode;
		var result_json = JSON.parse(result);
		//var trInfo = result_json.TransactionInfo;
		//var cn = trInfo.length;
		console.log("result: " + result);
		console.log("status code: " + code);
		res.render('blockchain/index', {result_json});
	});
})

module.exports = router;
