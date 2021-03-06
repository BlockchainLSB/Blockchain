var express = require('express');
var router = express.Router();
var peer = 'peer';
var channel = 'mychannel';
var chaincode = 'mycc';

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

var query_chainInfo = function(api_token, api_port,callback) {
	var jsonheaders = {
					"Authorization": "Bearer " + api_token,
					"Content-Type" : "application/json"
					};
	object.headers = jsonheaders;
	var api_host = 'http://52.79.245.63:' + api_port;
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
	var user_id = sess.user_id;
	
	var api_token = sess.api_token;
	var api_port = sess.api_port;
	query_portfolio(api_token, api_port,'getUserTransaction', ['token', token], function(data, statusCode){
		var result = data;
		var code = statusCode;
		var result_json = JSON.parse(result);
		
		console.log("result: " + result);
		console.log("code: " 	+ code);
		var page_size = 5; // 5 row per 1page
		var page_list_size = 5; // # of pages
		var no = ""; // var limit
		var totalPageCount;  // total # of row
		var curPage = req.query.cur; // current Page
		var transactionInfos = null;
		var startPageNum, totalPage, totalSet, curSet, startPage, endPage, iStart, iEnd, previous, next, totalSet;
		
		if(result_json.TransactionInfo == null) {
			totalPageCount = 0;
			startPageNum = 0;
			totalPage = 0;
			totalSet = 0;
			curSet = 0;
			startPage = 0;
			endPage = 0;
			iStart = 0;
			iEnd = 0;
			previous = 1;
			next = 1;
		}
		else {
			totalPageCount = result_json.TransactionInfo.length;
			// TxId, Value,Timestamp
			transactionInfos = result_json.TransactionInfo;
			totalPage = Math.ceil(totalPageCount / page_size); // total # of pages
			totalSet = Math.ceil(totalPage / page_list_size); // total # of sets
			curSet = Math.ceil(curPage / page_list_size); // current set #
			startPage = ((curSet - 1) * 5) + 1 // 현재 세트내 출력될 시작 페이지
			endPage = (startPage + page_list_size) - 1; // 현재 세트내 출력될 마지막 페이지
			iStart = (curPage*page_size) - page_size;
			iEnd = (curPage*page_size);
			
			if (iEnd > totalPageCount) {
				iEnd = totalPageCount
			}
			if(endPage > totalPage) {
				endPage = totalPage;
			}
			
			if(curSet == 1) {
				previous = 1;
			}
			else {
				previous = (curSet - 1) * page_list_size 
			}
			
			if(curSet >= totalSet) {
				console.log("curset >= totalSet")
				next = totalPage;
			}
			else {
				console.log("before curSet != total set - curPage: " + curPage)
				next = Number(curPage) + 1;
			}
		}
		
		if(curPage < 0) {
			no = 0;
		}
		else {
			no = (curPage - 1) * 10
		}
		
		console.log('[0] curPage: ' + curPage + ' | [1] page_list_size: ' + page_list_size );
		console.log('page_size: ' + page_size + ' ,totalPage' + totalPage + ' ,totalSet' + totalSet + ' ,curSet' + curSet + ' ,startPage' + startPage + ' ,endPage' + endPage);
		console.log("previous: " + previous + ",next: " + next);
		var pageInfo = {
			"curPage": curPage,
			"page_list_size": page_list_size,
			"page_size": page_size,
			"totalPage": totalPage,
			"totalSet": totalSet,
			"curSet": curSet,
			"startPage": startPage,
			"endPage": endPage,
			"iStart": iStart,
			"iEnd": iEnd,
			"previous": previous,
			"next": next,
			"tot": totalPageCount
		}
		console.log("iStart: " + iStart + ",iEnd: " + iEnd);
		query_chainInfo(api_token, api_port,function(data, statusCode) {
			var cresult_json = data;
			var ccode = statusCode;
			  
			console.log("status code: " + ccode);
			res.render('blockchain/index', {cresult_json, transactionInfos, login, pageInfo, user_id, api_token, api_port});		
		});
		
		
	});
})

module.exports = router;