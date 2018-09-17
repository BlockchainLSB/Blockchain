/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type keyType string

const (
	passwordType keyType = "pw_"
	tokenType    keyType = "tk_"
	temporalType keyType = "tmp_"
)

// UserChaincode example simple Chaincode implementation
type UserChaincode struct {
	stub     shim.ChaincodeStubInterface
	function string
	args     []string
}

type User struct {
	Id     string `json:"id"`
	Passwd string `json:"passwd"`
	Email  string `json:"email"`
}

func (t *UserChaincode) call() pb.Response {
	function := t.function

	callMap := map[string]func() pb.Response{
		"signup": t.signup,
		"signin": t.signin,
		"token":  t.token,
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

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	id := args[0]
	pw := args[1]

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

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	id := args[0]
	pw := args[1]

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

func (t *UserChaincode) token() pb.Response {

	fmt.Println("========================= token =========================")

	args := t.args

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	id := args[0]
	pw := args[1]

	pwd, err := t.stub.GetState(string(passwordType) + id)

	if err != nil {
		return shim.Error(id + "is not registered.")
	}

	if pw == string(pwd) {
		token, err := t.stub.GetState(string(tokenType) + id)

		if err != nil {
			return shim.Error(err.Error())
		}

		return shim.Success(token)
	}
	return shim.Error("Incorrect password.")
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
