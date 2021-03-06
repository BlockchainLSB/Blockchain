/* module */
var fs = require('fs')
    ,express = require('express')
    ,bodyParser = require('body-parser');
const path = require('path');

/* REST API (POST) 사용 */
var app = express();
app.use(bodyParser.urlencoded({extended:true}));
app.use(bodyParser.json());
app.use(express.static('public'));
app.set('view engine', 'pug'); 
app.set('views', path.join(__dirname, 'views')); 

/* json 파일 object 파일로 변환 */
var contents = fs.readFileSync("json/token.json");
var heatContent = fs.readFileSync("json/template.json");
var jsonContent = JSON.parse(contents);
var heatJsonContent = JSON.parse(heatContent);
var object = {};
var heatObject = {};

object.data = jsonContent;
heatObject.data = heatJsonContent;
jsonheaders = {"Content-Type" : "application/json"};
object.headers = jsonheaders;

/* api */
var Client = require('node-rest-client').Client;
var client = new Client();

var countNum = 0;

/* 메인 페이지 */
app.get('/', function(req, res){
    fs.readFile('index.html', function(err,data){
        res.writeHead(200, {"Content-Type" : "text/html"});
        res.end(data);
    });
})

app.get('/select', function(req, res){
    fs.readFile('select.html', function(err,data){
        res.writeHead(200, {"Content-Type" : "text/html"});
        res.end(data);
    });
})

/* 로그인 */
app.post('/select', function(req, res, next){
    //console.log(req.body);
    var id = req.body.id;
    var pwd = req.body.pwd;

    /* token.json 바꾸기 */
    jsonContent.auth.identity.password.user.name = id;
    jsonContent.auth.identity.password.user.password = pwd;

    /* api 보내기 */
    client.registerMethod("postMethod", "http://164.125.70.26:801/identity/v3/auth/tokens", "POST");
    client.registerMethod("jsonMethod", "http://remote.site/rest/json/method", "GET");
    
    client.methods.postMethod(object, function (data, response) {
        // parsed response body as js object
        // raw response
        var headers = JSON.stringify(response.headers);
        var jsonValue = JSON.parse(headers);
        console.log("--------------- token value ---------------\n", jsonValue['x-subject-token']);

        var stream = fs.createWriteStream("token.txt");
        stream.once('open', function(fd){
            stream.write(jsonValue['x-subject-token']);
            stream.end();
        });
    });
    res.redirect('/select');
});

app.get('/select/nonfunc', function(req, res){

    fs.readFile('nonfunctional.html', function(err, data){
        res.writeHead(200, {"Content-Type" : "text/html"});
        res.end(data);
    });
});

app.post('/select/nonfunc', function(req, res, next){
    // 총 기기 수
    //var countNum = 0;
    var funcContent = JSON.parse(JSON.stringify(req.body));
    countNum = countMachine(funcContent);
    

    //function select
    res.redirect('/select/nonfunc');
    //res.render('selectspeed', {name : req.body.machineName , nameq : req.body.systemName});

});

app.get('/select/nonfunc/createstack', function(req, res){
    fs.readFile('createstack.pug', function(err,data){
        res.writeHead(200, {"Content-Type" : "text/html"});
        res.end(data);
    });


});

app.post('/select/nonfunc/createstack', function(req, res, next){

    console.log(req.body);

    var Jsonimage={"image" : "09e1182b-d088-4a1b-bdc8-b6b9e67cbcd3"};
    var nonfuncContent = JSON.parse(JSON.stringify(req.body));
    var Jsonflavor = DistinctResource(nonfuncContent, countNum);
    //flavor = JSON.parse(flavor);

    //console.log("flavor : ", flavor);
    //console.log("type: ", typeof(req.body.fullname));

    heatJsonContent.stack_name = req.body.fullname;
    heatJsonContent.parameters.flavor = Jsonflavor.flavor; //?
    heatJsonContent.template.resources.hello_world.properties.image = Jsonimage.image;
    heatObject.data = heatJsonContent;

    var getToken = fs.readFileSync("token.txt", 'utf8');
    var token = JSON.stringify(getToken);
    var heatHeaders = {"Content-Type" : "application/json",
                      "X-Auth-Token" :getToken};
    heatObject.headers = heatHeaders;

    client.post("http://164.125.70.26:801/heat-api/v1/48a1d38373a14cb9b89a6ddcea0ffc0f/stacks", heatObject, function(data, response) {
        console.log("CompleteCreateStack!");
    });

    console.log("#####", heatObject);

    var systemJson = req.body.systemName;
    var machineJson = req.body.machineName;
    var bmesJson = req.body.bmesName;
    var outsideJson = req.body.outsideName;

    res.render('createstack', {name : req.body.systemName, nameq : req.body.machineName});
});

app.listen(3023, function(){
    console.log('Server start.');
});

app.all('*', function(req, res) {
	res.status(404).send('<h1>ERROR - 페이지를 찾을 수 없습니다.</h1>');
});


function countMachine(data){

    var machineJson = data.machineName;
    var outsideJson = data.outsideName;
    var Count=0;

    for(var i=0; i<machineJson.length; i++){
        Count += parseInt(machineJson[i],10);
    }
    for(var j=0; j<outsideJson.length; j++){
        Count += parseInt(outsideJson[j],10);
    }
    return Count;

}


function DistinctResource(data, count){

    var speed = data.speedName;
    var flavor;
    var Jsonflavor;
    if(speed == "LOW" && ((1<=count) && (count <=20))) 
    if((1<=count) && (count <= 20))
    {
        if(speed == "LOW") flavor = "m1.tiny";
        else if(speed == "MIDDLE") flavor = "m1.custom2";
        else flavor = "m1.small";
    }
    if((21<=count) && (count <= 40))
    {
        if(speed == "LOW") flavor = "m1.custom2";
        else if(speed == "MIDDLE") flavor = "m1.small";
        else flavor = "m1.medium";
    }
    if((41<=count) && (count <= 60))
    {
        if(speed == "LOW") flavor = "m1.small";
        else if(speed == "MIDDLE") flavor = "m1.medium";
        else flavor = "m1.large";
    }
    if((61<=count) && (count <= 80))
    {
        if(speed == "LOW") flavor = "m1.medium";
        else flavor = "m1.large";
    }
    if((81<=count) && (count <= 100))
    {
        flavor = "m1.xlarge";
    }

    Jsonflavor = {"flavor" : flavor};
    return Jsonflavor;
}