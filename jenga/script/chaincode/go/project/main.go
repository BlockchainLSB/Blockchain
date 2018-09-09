package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type keyType string

const (
	commonChannel string = "common"
)

const (
	userChaincode string = "user"
)

const (
	tokenFunction        string = "token"
	userByTokenrFunction        = "userByToken"
	addProjectFunction          = "addProject"
)

const (
	userType    keyType = "us_"
	projectType keyType = "pr_"
)

func toChaincodeArgs(args ...string) [][]byte {
	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}
	return bargs
}

type ProjectChaincode struct {
	stub     shim.ChaincodeStubInterface
	function string
	args     []string
}

func (t *ProjectChaincode) call() pb.Response {
	function := t.function

	callMap := map[string]func() pb.Response{
		"create":                  t.create,
		"overwrite":               t.overwrite,
		"projectInfo":             t.loadProjectInformation,
		"addDone":                 t.addDone,
		"addAttendee":             t.addAttendee,
		"signDone":                t.signDone,
		"projectInfoWithoutToken": t.loadProjectInformationWithoutToken,
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

func (t *ProjectChaincode) checkToken(token string) (string, error) {
	res := t.invokeChaincode(userChaincode, commonChannel, userByTokenrFunction, token)

	if string(res.Payload) == "" {
		return "", errors.New("Invalid token")
	}

	return string(res.Payload), nil
}

func (t *ProjectChaincode) state(typ keyType, id string) ([]byte, error) {
	return t.stub.GetState(string(typ) + id)
}

func (t *ProjectChaincode) putState(typ keyType, id string, data []byte) error {
	return t.stub.PutState(string(typ)+id, data)
}

func (t *ProjectChaincode) putProjectState(prj project) error {
	b, err := json.Marshal(prj)

	if err != nil {
		return err
	}

	var v map[string]interface{}
	err = json.Unmarshal(b, &v)

	if err != nil {
		return err
	}

	return t.putMapState(string(projectType)+prj.Admin+"_"+prj.Name+"_", v)
}

func (t *ProjectChaincode) putMapState(prefix string, v map[string]interface{}) (err error) {
	for k, val := range v {
		if err := t.putTypeAssertion(keyType(prefix), k, val); err != nil {
			return err
		}
	}

	return nil
}

func (t *ProjectChaincode) putArrayState(prefix string, v []interface{}) error {
	t.putState(keyType(prefix), "len", []byte(strconv.Itoa(len(v))))

	for i, val := range v {
		k := strconv.Itoa(i)
		if err := t.putTypeAssertion(keyType(prefix), k, val); err != nil {
			return err
		}
	}

	return nil
}

func (t *ProjectChaincode) putTypeAssertion(prefix keyType, k string, v interface{}) (err error) {
	switch tv := v.(type) {
	case float64:
		b := strconv.FormatInt(int64(tv), 10)
		err = t.putState(keyType(prefix), k, []byte(b))
	case string:
		err = t.putState(keyType(prefix), k, []byte(tv))
	case bool:
		b := make([]byte, 1)
		if tv {
			b[0] = 1
		}
		err = t.putState(keyType(prefix), k, b)
	case []interface{}:
		err = t.putArrayState(string(prefix)+k+"_", tv)
	case map[string]interface{}:
		err = t.putMapState(string(prefix)+k+"_", tv)
	default:
		typ := reflect.TypeOf(tv)
		err = errors.New("Cannot determine a type of " + k + "'s value(" + typ.Name() + ")")
	}

	return
}

func (t *ProjectChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	_, args := stub.GetFunctionAndParameters()

	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Expecting 0")
	}

	return shim.Success(nil)
}

func (t *ProjectChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	t.function = function
	t.args = args
	t.stub = stub

	return t.call()
}

func (t *ProjectChaincode) create() pb.Response {
	args := t.args

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	token := args[0]

	user, err := t.checkToken(token)

	if err != nil {
		return shim.Error(err.Error())
	}

	var p project
	err = json.Unmarshal([]byte(args[1]), &p)
	if err != nil {
		return shim.Error(err.Error())
	}

	if user != p.Admin {
		return shim.Error("You don't have a admin permission to create the project")
	}

	p.ID = p.Admin + "_" + p.Name
	err = t.putProjectState(p)

	if err != nil {
		return shim.Error(err.Error())
	}

	return t.invokeChaincode(userChaincode, commonChannel, addProjectFunction, token, p.ID)
}

func (t *ProjectChaincode) overwrite() pb.Response {
	args := t.args

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	token := args[0]

	user, err := t.checkToken(token)

	if err != nil {
		return shim.Error(err.Error())
	}

	var p project
	err = json.Unmarshal([]byte(args[1]), &p)
	if err != nil {
		return shim.Error(err.Error())
	}

	if user != p.Admin {
		return shim.Error("You don't have a admin permission to create the project")
	}

	p.ID = p.Admin + "_" + p.Name
	err = t.putProjectState(p)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *ProjectChaincode) loadProjectInformation() pb.Response {
	args := t.args

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	token := args[0]
	pn := args[1]

	user, err := t.checkToken(token)

	if err != nil {
		return shim.Error(err.Error())
	}

	// TODO: check whether the user is an attendee in the project.
	_ = user

	pj := new(project)

	if err = pj.load(t.stub, pn); err != nil {
		return shim.Error(err.Error())
	}

	res, err := json.Marshal(pj)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte(res))
}

func (t *ProjectChaincode) loadProjectInformationWithoutToken() pb.Response {
	args := t.args

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	pn := args[0]

	pj := new(project)

	if err := pj.load(t.stub, pn); err != nil {
		return shim.Error(err.Error())
	}

	res, err := json.Marshal(pj)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte(res))
}

func (t *ProjectChaincode) invokeChaincode(name, channel, function string, args ...string) pb.Response {
	q := toChaincodeArgs(append([]string{function}, args...)...)

	return t.stub.InvokeChaincode(name, q, channel)
}

func main() {
	err := shim.Start(new(ProjectChaincode))
	if err != nil {
		fmt.Printf("Error starting UserTest chaincode: %s", err)
	}
}
