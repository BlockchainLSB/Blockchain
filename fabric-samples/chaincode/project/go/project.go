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

func (t *ProjectChaincode) call() pb.Response {
	function := t.function

	callMap := map[string]func() pb.Response{
		"create":	t.create,
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

func (t *ProjectChaincode) create() pb.Response {

	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(ProjectChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
