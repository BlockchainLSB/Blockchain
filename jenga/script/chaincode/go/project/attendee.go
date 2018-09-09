package main

import (
	"encoding/json"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Attendee struct {
	Name         string `json:"name"`
	Account      string `json:"account"`
	Affiliation  string `json:"affiliation"`
	Role         string `json:"role"`
	Contribution int64  `json:"contribution"`
}

func (a *Attendee) save(stub shim.ChaincodeStubInterface, bk string) error {
	alb, err := stub.GetState(bk + "_" + stateProjectAttendees + "_len")
	if err != nil {
		return err
	}
	als := string(alb)

	tk := bk + "_" + stateProjectAttendees + "_" + als + "_"
	stub.PutState(tk+stateAttendeeAccount, []byte(a.Account))
	stub.PutState(tk+stateAttendeeRole, []byte(a.Role))
	stub.PutState(tk+stateAttendeeAffiliation, []byte(a.Affiliation))
	stub.PutState(tk+stateAttendeeContribution, []byte(strconv.FormatInt(a.Contribution, 10)))
	stub.PutState(tk+stateAttendeeName, []byte(a.Name))

	al, _ := strconv.Atoi(als)
	stub.PutState(bk+"_"+stateProjectAttendees+"_len", []byte(strconv.Itoa(al+1)))

	return nil
}

func (t *ProjectChaincode) addAttendee() pb.Response {
	args := t.args

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	token := args[0]
	pid := args[1]

	var attendee Attendee

	err := json.Unmarshal([]byte(args[2]), &attendee)
	if err != nil {
		return shim.Error(err.Error())
	}

	admin, err := t.checkToken(token)

	if err != nil {
		return shim.Error(err.Error())
	}

	bk := keyType(prefix + pid)
	pj := &project{}

	if err = pj.load(t.stub, pid); err != nil {
		return shim.Error(err.Error())
	}

	if err != nil {
		return shim.Error(err.Error())
	}

	if admin != pj.Admin {
		return shim.Error("You don't have a permission to add an attendee into the project.(" + admin + "," + pj.Admin + ")")
	}

	for _, v := range pj.Attendees {
		if v.Account == attendee.Account {
			return shim.Error(attendee.Name + " was alreadt attended in this project.")
		}
	}

	err = attendee.save(t.stub, string(bk))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}
