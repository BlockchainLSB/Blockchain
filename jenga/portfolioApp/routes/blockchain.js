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

var jsonheaders = {
					"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Mzc5NzgwMzQsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Ik9yZzEiLCJpYXQiOjE1Mzc5NDIwMzR9.GEqG7hFWyQTQVVlLUUGnDYmkQknNqSwKpE-AkaUX2_4",
					"Content-Type" : "application/json"
					};
object.headers = jsonheaders;

var query_portfolio = function(fcn, args, callback){
	var api_url = api_host + '/channels/mychannel/chaincodes/mycc?peer=peer0.org1.example.com&fcn='+fcn+'&args='+JSON.stringify(args||null);

	client.registerMethod("queryUserMethod", api_url, "GET");
    client.methods.queryUserMethod(object, function (data, response) {
    	console.log("data : " + data );
		var statusCode = response.statusCode;
		callback(data, statusCode);
	});
}

var query_chainInfo = function(callback) {
	var api_url = api_host + '/channels/mychannel?peer=peer0.org1.example.com';
	client.registerMethod("queryChainMethod", api_url, "GET");
	client.methods.queryChainMethod(object, function (data, response) {
    	console.log("data : " + JSON.stringify(data));
		var statusCode = response.statusCode;
		callback(data, statusCode);
	});
}

router.get('/', function(req, res, next){
	var sess = req.session;
	var token = sess.token;
	var login = sess.login;
	console.log('user token : ' + token);
	
	query_portfolio('getUserTransaction', ['token', token], function(data, statusCode){
		var result = data;
		var code = statusCode;
		var result_json = JSON.parse(result);
		//var trInfo = result_json.TransactionInfo;
		//var cn = trInfo.length;
		console.log("result: " + result);
		console.log("status code: " 	+ code);
		
		query_chainInfo(function(data, statusCode) {
			var cresult_json = data;
			var ccode = statusCode;
			var blockH = data.height.low;
			var currentBHoffset = data.currentBlockHash.offset;  
			console.log("status code: " + ccode);
			console.log("block height: " + blockH)
			res.render('blockchain/index', {cresult_json, result_json, login});		
		});
		
		
	});
})

module.exports = router;
