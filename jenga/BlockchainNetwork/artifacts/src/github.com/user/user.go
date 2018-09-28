package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type keyType string

var project_id int = 0

const (
	passwordType keyType = "pw_"
	tokenType    keyType = "tk_"
	temporalType keyType = "tmp_"
	projectType  keyType = "prj_"
)

// UserChaincode example simple Chaincode implementation
type UserChaincode struct {
	stub     shim.ChaincodeStubInterface
	function string
	args     []string
}

type User struct {
	Id    string
	Pw    string
}
type UserProject struct {
	Pnum           int
	Pname          string
	PDes           string
	POkContributor []projectContributor
}

type projectContributor struct {
	Cname string
	C_ok  bool
}

func UserProjectInit(pnum int, pname string, pdes string, pokcontributor []string) UserProject {

	// Create struct and append it to the slice.
	up := UserProject{}
	pCs := []projectContributor{}

	up.Pnum = pnum
	up.Pname = pname
	up.PDes = pdes

	//okcontributor 배열
	for i := 0; i < len(pokcontributor); i++ {
		pC := projectContributor{}
		pC.Cname = pokcontributor[i]
		if i == 0 {
			pC.C_ok = true
		} else {
			pC.C_ok = false
		}
		pCs = append(pCs, pC)
	}
	up.POkContributor = pCs

	return up
}

func UserInit(id string, pw string) User {

	user := User{}
	user.Id = id
	user.Pw = pw
	
	return user
}

func (t *UserChaincode) call() pb.Response {
	function := t.function

	callMap := map[string]func() pb.Response{
		"signup":      t.signup,
		"signin":      t.signin,
		"getToken":    t.getToken,
		"addProject":  t.addProject,
		"loadProject": t.loadProject,
		"searchUser":  t.searchUser,
	}

	h := callMap[function]
	if h != nil {
		return callMap[function]()
	}

	res := make([]string, 0)
	for k := range callMap {
		res = append(res, `"`+k+`"`)
	}

	return shim.Error("Invalid invoke function name. Expecting " + strings.Join(res, ", "))
}

func (t *UserChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Portfolio is starting up")
	_, args := stub.GetFunctionAndParameters()

	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Expecting 0")
	}

	return shim.Success(nil)
}

func (t *UserChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("========================= Invoke =========================")

	function, args := stub.GetFunctionAndParameters()

	t.function = function
	t.args = args
	t.stub = stub

	return t.call()
}

func (t *UserChaincode) signup() pb.Response {

	fmt.Println("========================= signup =========================")

	args := t.args

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	
	id := args[1]
	pw := args[3]

	u := UserInit(id,pw)

	val, err := t.stub.GetState(string(passwordType) + id)
	if err == nil && val != nil && len(val) > 0 {
		return shim.Error(id + " has been registered.")
	}

	err = t.stub.PutState(string(passwordType)+id, []byte(pw))

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *UserChaincode) signin() pb.Response {

	fmt.Println("========================= signin =========================")

	args := t.args

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	id := args[1]
	pw := args[3]

	// id로 pw 찾을 수 있음
	pwb, err := t.stub.GetState(string(passwordType) + id)
	if err != nil {
		return shim.Error(id + " is not registered.")
	}

	if pw == string(pwb) {
		rand.Seed(time.Now().UnixNano())
		var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

		b := make([]rune, 10)
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}

		token := string(b)
		oldToken, _ := t.stub.GetState(string(tokenType) + id)

		if oldToken != nil && len(oldToken) > 0 {
			t.stub.DelState(string(temporalType) + string(oldToken))
		}

		// 새로운 토큰 값 id 를 key로 해서 저장
		err = t.stub.PutState(string(tokenType)+id, []byte(token))

		if err != nil {
			return shim.Error(err.Error())
		}

		err = t.stub.PutState(string(temporalType)+string([]byte(token)), []byte(string(id)))

		if err != nil {
			return shim.Error(err.Error())
		}

		return shim.Success([]byte("{ \"is_auth\": true, \"token\": \"" + token + "\"}"))
	}

	return shim.Error("{ \"is_auth\": false }")
}

func (t *UserChaincode) getToken() pb.Response {

	fmt.Println("========================= token =========================")

	//var id, pw string       // Entity
	var idVal, pwVal string // value
	args := t.args

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	idVal = args[1]
	pwVal = args[3]

	// id로 pw 찾기
	pwd, err := t.stub.GetState(string(passwordType) + idVal)

	fmt.Println(string(pwd))

	if err != nil {
		return shim.Error(idVal + "is not registered.")
	}

	if pwVal == string(pwd) {
		token, err := t.stub.GetState(string(tokenType) + idVal)

		if err != nil {
			return shim.Error(err.Error())
		}

		return shim.Success([]byte("{\"token\": \"" + string(token) + "\"}"))
	}

	return shim.Error("Incorrect password.")
}

func (t *UserChaincode) searchUser() pb.Response {
	fmt.Println("========================= searchUser =========================")

	args := t.args

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	id := args[1]

	// id로 pw 찾을 수 있음
	pw, err := t.stub.GetState(string(passwordType) + id)

	if err != nil {
		fmt.Println("id is not registered")
		return shim.Error(id + " is not registered")
	}

	if len(pw) != 0 {
		fmt.Println(id + " is registered")
	} else if len(pw) == 0 {
		return shim.Error(id + " is not registered")
	}
	return shim.Success(nil)

}

func (t *UserChaincode) getuserInfo(key string) ([]byte, error) pb.Response {
	return t.stub.GetState(string(temporalType) + key)
}
func (t *UserChaincode) addProject() pb.Response {

	/*

		 {
        "Pnum": 1,
        "Pname": "job10",
        "PDes": "qwe",
        "POkContributor": [
            {
                "Cname": "lye",
                "C_ok": true
            },
            {
                "Cname": "123q",
                "C_ok": false
            },
            {
                "Cname": "qwe",
                "C_ok": false
            }
        ]
    }

	*/

	fmt.Println("========================= addProject =========================")
	args := t.args
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}

	var contributorArray []string

	token := args[1]
	projectname := args[3]
	description := args[5]
	contributorlist := args[7]

	contributorArray = strings.Split(contributorlist, ",")

	//token 값으로 id 값 받아오기
	//userid, err := t.stub.GetState(string(temporalType) + token)
	userid, err := t.getuserInfo(token)

	if err != nil {
		return shim.Error(err.Error())
	}

	existProject, perr := t.stub.GetState(string(projectType) + string(userid))

	if perr != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("res -> ", string(existProject))

	//res := make([]string, 0)
	//_ = json.Unmarshal(existProject, &res)

	data := make([]*UserProject, 0)
	_ = json.Unmarshal(existProject, &data)

	// project 이름이 동일한 지 확인
	for _, v := range data {
		//fmt.Println("string(v) : ", string(v))
		if v.Pname == projectname {
			return shim.Error(projectname + " was already added.")
		}
	}
	project_id = project_id + 1
	p := UserProjectInit(project_id, projectname, description, contributorArray)

	data = append(data, &p)
	doc, jerr := json.MarshalIndent(data, "", "    ")

	if jerr != nil {
		return shim.Error("json Marshal Error")
	}

	fmt.Println("Transform JSON " + string(doc))

	err = t.stub.PutState(string(projectType)+string(userid), doc)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}

func (t *UserChaincode) loadProject() pb.Response {

	fmt.Println("========================= loadProject =========================")

	args := t.args

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	token := args[1]

	//token으로 id 찾기
	userid, err := t.stub.GetState(string(temporalType) + string(token))

	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println(string((userid)))

	doc, err := t.stub.GetState(string(projectType) + string(userid))

	if err != nil {
		return shim.Error(err.Error())
	}
	res := make([]*UserProject, 0)
	_ = json.Unmarshal(doc, &res)
	fmt.Println(string(doc))

	var data []string

	for _, v := range res {
		if len(v.Pname) != 0 {
			data = append(data, v.Pname)
		}
	}

	fmt.Println("project list :", data)
	/*
			projects := make([]map[string]interface{}, 0)

			for _, pid := range res {
				r := t.invokeChaincode("project", "common", "projectInfoWithoutToken", pid)
				var v map[string]interface{}
				_ = json.Unmarshal(r.Payload, &v)
				projects = append(projects, v)
			}

		b, _ := json.Marshal(projects)
	*/
	return shim.Success(nil)
}

func (t *UserChaincode) invokeChaincode(name, channel, function string, args ...string) pb.Response {
	q := toChaincodeArgs(append([]string{function}, args...)...)

	return t.stub.InvokeChaincode(name, q, channel)
}

func toChaincodeArgs(args ...string) [][]byte {
	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}
	return bargs
}

/*
// query callback representing the query of a chaincode
func (t *UserChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string // Entities
	var err error

	fmt.Println("portfolio query")

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}
*/

func main() {
	err := shim.Start(new(UserChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
