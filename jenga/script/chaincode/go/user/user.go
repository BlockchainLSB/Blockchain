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

const (
	passwordType    keyType = "pw_"
	tokenType       keyType = "tk_"
	temporalType    keyType = "tmp_"
	projectType     keyType = "prj_"
	starProjectType keyType = "sprj_"
)

// UserChaincode manges users and token.
type UserChaincode struct {
	stub     shim.ChaincodeStubInterface
	function string
	args     []string
}

func (t *UserChaincode) call() pb.Response {
	function := t.function

	callMap := map[string]func() pb.Response{
		"signup":              t.signup,
		"signin":              t.signin,
		"token":               t.token,
		"userByToken":         t.userByToken,
		"addProject":          t.addUserProject,
		"addFavoriteProject":  t.addFavoriteProject,
		"loadProject":         t.loadUserProject,
		"loadFavoriteProject": t.loadFavoriteProject,
		"delFavoriteProject":  t.delFavoriteProject,
		"delProject":          t.delUserProject,
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

func (t *UserChaincode) state(typ keyType, id string) ([]byte, error) {
	return t.stub.GetState(string(typ) + id)
}

func (t *UserChaincode) putState(typ keyType, id string, data []byte) error {
	return t.stub.PutState(string(typ)+id, data)
}

func (t *UserChaincode) tokenState(key string) ([]byte, error) {
	return t.stub.GetState(string(temporalType) + key)
}

func (t *UserChaincode) delTokenState(key string) error {
	return t.stub.DelState(string(temporalType) + key)
}

func (t *UserChaincode) putTokenState(id string, token []byte) error {
	return t.stub.PutState(string(temporalType)+string(token), []byte(id))
}

func (t *UserChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	_, args := stub.GetFunctionAndParameters()

	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Expecting 0")
	}

	return shim.Success(nil)
}

func (t *UserChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	t.function = function
	t.args = args
	t.stub = stub

	return t.call()
}

func (t *UserChaincode) signup() pb.Response {
	args := t.args

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	id := args[0]
	pw := args[1]

	val, err := t.state(passwordType, id)
	if err == nil && val != nil && len(val) > 0 {
		return shim.Error(id + " has been registered.")
	}

	err = t.putState(passwordType, id, []byte(pw))

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *UserChaincode) signin() pb.Response {
	args := t.args

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	id := args[0]
	pw := args[1]

	pwb, err := t.state(passwordType, id)
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
		oldToken, _ := t.state(tokenType, id)

		if oldToken != nil && len(oldToken) > 0 {
			t.delTokenState(string(oldToken))
		}

		err = t.putState(tokenType, id, []byte(token))

		if err != nil {
			return shim.Error(err.Error())
		}

		err = t.putTokenState(id, []byte(token))

		if err != nil {
			return shim.Error(err.Error())
		}

		return shim.Success([]byte("{ \"is_auth\": true, \"token\": \"" + token + "\"}"))
	}

	return shim.Error("Incorrect password.")
}

func (t *UserChaincode) token() pb.Response {
	args := t.args

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	id := args[0]
	pw := args[1]

	pwb, err := t.state(passwordType, id)
	if err != nil {
		return shim.Error(id + " is not registered.")
	}

	if pw == string(pwb) {
		token, err := t.state(tokenType, id)

		if err != nil {
			return shim.Error(err.Error())
		}

		return shim.Success(token)
	}

	return shim.Error("Incorrect password.")
}

func (t *UserChaincode) userByToken() pb.Response {
	args := t.args

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	token := args[0]

	user, err := t.tokenState(token)

	if err != nil {
		return shim.Error(err.Error())
	}

	if string(user) == "" {
		return shim.Error("Invaild token.")
	}

	return shim.Success(user)
}

func (t *UserChaincode) addProject(token, pn string, typ keyType) pb.Response {
	user, err := t.tokenState(token)

	if err != nil {
		return shim.Error(err.Error())
	}

	la, err := t.state(typ, string(user))

	res := make([]string, 0)

	_ = json.Unmarshal(la, &res)

	for _, v := range res {
		if string(v) == pn {
			return shim.Error(pn + "was already added.")
		}
	}

	res = append(res, pn)

	jr, err := json.Marshal(res)

	if err != nil {
		return shim.Error(err.Error())
	}

	if err = t.putState(typ, string(user), jr); err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *UserChaincode) addUserProject() pb.Response {
	args := t.args

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	token := args[0]
	pn := args[1]

	return t.addProject(token, pn, projectType)
}

func (t *UserChaincode) addFavoriteProject() pb.Response {
	args := t.args

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	token := args[0]
	pn := args[1]

	return t.addProject(token, pn, starProjectType)
}

func (t *UserChaincode) loadProject(token string, typ keyType) []string {
	user, err := t.tokenState(token)

	res := make([]string, 0)

	if err != nil {
		return res
	}

	la, err := t.state(typ, string(user))

	if err != nil {
		return res
	}

	_ = json.Unmarshal(la, &res)

	return res
}

func (t *UserChaincode) loadProjectInfo(p []string) pb.Response {
	projects := make([]map[string]interface{}, 0)

	for _, pid := range p {
		r := t.invokeChaincode("project", "common", "projectInfoWithoutToken", pid)
		var v map[string]interface{}
		_ = json.Unmarshal(r.Payload, &v)
		projects = append(projects, v)
	}

	b, _ := json.Marshal(projects)

	return shim.Success(b)
}

func toChaincodeArgs(args ...string) [][]byte {
	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}
	return bargs
}

func (t *UserChaincode) invokeChaincode(name, channel, function string, args ...string) pb.Response {
	q := toChaincodeArgs(append([]string{function}, args...)...)

	return t.stub.InvokeChaincode(name, q, channel)
}

func (t *UserChaincode) loadUserProject() pb.Response {
	args := t.args

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	token := args[0]
	p := t.loadProject(token, projectType)

	return t.loadProjectInfo(p)
}

func (t *UserChaincode) loadFavoriteProject() pb.Response {
	args := t.args

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	token := args[0]

	p := t.loadProject(token, starProjectType)

	b, _ := json.Marshal(p)
	return shim.Success(b)
}

func (t *UserChaincode) delProject(token, pid string, typ keyType) pb.Response {
	user, _ := t.tokenState(token)

	p := t.loadProject(token, typ)

	for i := 0; i < len(p); i++ {
		if p[i] == pid {
			var res []string

			if (i + 1) == len(p) {
				res = p[:i]
			} else {
				res = append(p[:i], p[i+1:]...)
			}

			b, _ := json.Marshal(&res)

			t.putState(typ, string(user), b)
		}
	}

	return shim.Success(nil)
}

func (t *UserChaincode) delFavoriteProject() pb.Response {
	args := t.args

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	token := args[0]
	pid := args[1]

	return t.delProject(token, pid, starProjectType)
}

func (t *UserChaincode) delUserProject() pb.Response {
	args := t.args

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	token := args[0]
	pid := args[1]

	return t.delProject(token, pid, projectType)
}

func main() {
	err := shim.Start(new(UserChaincode))
	if err != nil {
		fmt.Printf("Error starting User chaincode: %s", err)
	}
}
