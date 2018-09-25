var express = require('express');
var router = express.Router();

/* GET home page. */

router.get('/', function(req, res, next){
	sess = req.session;
	console.log('sess.token : ' + sess.token);
	console.log('sess.login : ' + sess.login);
	res.render('project/index', {});
})

router.get('/repository/description', function(req, res, next){
	res.render('project/repository/description', {});
})

router.get('/repository/commit', function(req, res, next){
	res.render('project/repository/commit', {});
})

router.get('/repository/contributor', function(req, res, next){
	res.render('project/repository/contributor', {});
})

router.get('/static', function(req, res, next){
	res.render('project/static', {});
})

router.get('/addproject', function(req, res, next){
	res.render('project/addproject', {});
})

router.get('/evaluation', function(req, res, next){
	res.render('project/evaluation', {});
})



module.exports = router;
