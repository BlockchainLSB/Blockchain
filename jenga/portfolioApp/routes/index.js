var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
	var sess = req.session;
	var login = sess.login;
	var user_id = sess.user_id
  res.render('index', { title: 'BlockchainLSB', login, user_id });
}); 


router.get('/portfolio', function(req, res, next){
	res.render('portfolio/index', {});
});

module.exports = router;
