var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
	var sess = req.session;
	var api_token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzgxNTgxODUsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Ik9yZzEiLCJpYXQiOjE1MzgxMjIxODV9.K72ivhB2Xxau7pa5dRugfLcDz6XtQn-sTi0WDHk_l0U';
	var api_port = 4000;
	sess.api_token = api_token;
	sess.api_port = api_port;
	var login = sess.login;
	var user_id = sess.user_id
  	res.render('index', { title: 'BlockchainLSB', login, user_id, api_token, api_port });
}); 

module.exports = router;
