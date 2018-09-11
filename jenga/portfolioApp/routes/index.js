var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { title: 'Express' });
});


router.get('/project', function(req, res, next){
	res.render('project/index', {});
});

router.get('/portfolio', function(req, res, next){
	res.render('portfolio/index', {});
});

router.get('/blockchain', function(req, res, next){
	res.render('blockchain/index', {});
});

module.exports = router;