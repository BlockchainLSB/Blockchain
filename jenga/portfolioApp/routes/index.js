var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
	var sess = req.session;
	var login = sess.login;
  res.render('index', { title: 'Express', login });
}); 


router.get('/portfolio', function(req, res, next){
	res.render('portfolio/index', {});
});

module.exports = router;
