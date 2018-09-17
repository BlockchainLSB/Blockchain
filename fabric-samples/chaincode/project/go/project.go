package main

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// ProjectChaincode example simple Chaincode implementation
type ProjectChaincode struct {
	stub shim.ChaincodeStubInterface
	function string
	args []string
}

type Project struct {
	pId string 'json:"id"'
	pName string 'json:"name"'
	pDes string 'json:"des"'
	pOkContributors []string 'json:"okcontributor"'
}

// Project struct Constructor
func newProject (pId string, pName string, pDes string, pOkContributor string) *Project {
	p := Project{}
	p.pId = pId
	p.pName = pName
	p.pDes = pDes
	p.pOkContributors = []string{pOkContributor}
	return &p
}

func (t *ProjectChaincode) call() pb.Response {
	function := t.function

	callMap := map[string]func() pb.Response{
		"getProjectDetail":	t.getProjectDetail, // query
		"addTestProject": t.addTestProject, // for test
	}

	h := callMap[function]
	if h != nil {
		return callMap[function]()
	}

	res := make([]string, 0)
	for k := range callMap {
		res = append(res, k)
	}

	return shim.Error("Invalid invoke function name. Expecting " + strings.Join(res, ", "))
}

func (t *ProjectChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ProjectChaincode Init")
	_, args := stub.GetFunctionAndParameters()

	if len(args) != 0 {
		fmt.Println("len(args) != 0")
		return shim.Error("Incorrect # of arguments. Expecting 0")
	}
	fmt.Println("ProjectChaincode Succes")
	return shim.Success(nil)
}

func (t *ProjectChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	t.function = function
	t.args = args
	t.stub = stub

	return t.call()
}

func (t *ProjectChaincode) getProjectDetail() pb.Response {
	fmt.Println("getProjectDetail")

	return shim.Success(nil)
}

// pid, pname, pdes, pok
func (t *ProjectChaincode) addTestProject() pb.Response {
	fmt.Println("addTestProject")
	fmt.Println("args[0]: " + t.args[0]) // pid
	fmt.Println("args[1]: " + t.args[1]) // pname
	fmt.Println("args[2]: " + t.args[2]) // pDes
	fmt.Println("args[3]: " + t.args[3]) // pok

	p := newProject(t.args[0], t.args[1], t.args[2], t.args[3])
	fmt.Println("OK for mapping p structure and value")

	doc, _ := json.Marshal(p)
	fmt.Println("Transform JSON " + doc)

	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(ProjectChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
