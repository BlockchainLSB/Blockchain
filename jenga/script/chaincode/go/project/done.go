package main

import (
	"errors"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Job struct {
	Done     string   `json:"done"`
	Assignee string   `json:"assignee"`
	Signers  []string `json:"signers"`
}

func (j *Job) save(stub shim.ChaincodeStubInterface, bk string) error {
	alb, err := stub.GetState(bk + "_" + stateProjectJobs + "_len")
	if err != nil {
		return err
	}
	als := string(alb)

	tk := bk + "_" + stateProjectJobs + "_" + als + "_"
	stub.PutState(tk+stateJobDone, []byte(j.Done))
	stub.PutState(tk+stateJobAssignee, []byte(j.Assignee))
	stub.PutState(tk+stateJobSigners+"_len", []byte(strconv.Itoa(len(j.Signers))))

	for i, v := range j.Signers {
		stub.PutState(tk+stateJobSigners+"_"+strconv.Itoa(i), []byte(v))
	}

	al, _ := strconv.Atoi(als)
	stub.PutState(bk+"_"+stateProjectJobs+"_len", []byte(strconv.Itoa(al+1)))

	return nil
}

func signJob(stub shim.ChaincodeStubInterface, bk, signer string) error {
	lb, err := stub.GetState(bk + "_" + stateJobSigners + "_len")
	if err != nil {
		return err
	}

	l, err := strconv.Atoi(string(lb))

	for i := 0; i < l; i++ {
		tk := bk + "_" + stateJobSigners + "_" + strconv.Itoa(i)
		s, err := stub.GetState(tk)
		if err != nil {
			return err
		}

		if string(s) == signer {
			return errors.New(signer + " was already signed to this job.")
		}
	}

	if err = stub.PutState(bk+"_"+stateJobSigners+"_"+string(lb), []byte(signer)); err != nil {
		return err
	}

	return stub.PutState(bk+"_"+stateJobSigners+"_len", []byte(strconv.Itoa(l+1)))
}

func (t *ProjectChaincode) addDone() pb.Response {
	args := t.args

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	token := args[0]
	pid := args[1]
	done := args[2]

	assignee, err := t.checkToken(token)

	if err != nil {
		return shim.Error(err.Error())
	}

	bk := keyType(prefix + pid)
	job := Job{
		Done:     done,
		Assignee: assignee,
		Signers:  make([]string, 0),
	}

	if err := job.save(t.stub, string(bk)); err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *ProjectChaincode) signDone() pb.Response {
	args := t.args

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	token := args[0]
	jid := args[1]
	pid := strings.Split(jid, "jobs")[0]

	signer, err := t.checkToken(token)

	if err != nil {
		return shim.Error(err.Error())
	}

	if err = signJob(t.stub, prefix+jid, signer); err != nil {
		return shim.Error(err.Error())
	}

	p := &project{}

	err = p.load(t.stub, prefix+pid)
	if err != nil {
		shim.Error(err.Error())
	}

	l := strings.Split(jid, "_")
	i, _ := strconv.Atoi(l[len(l)-1])
	cj := p.Jobs[i]
	q := (len(p.Attendees) - (len(p.Attendees) % 2)) / 2

	if q+1 == len(cj.Signers) {
		for i, v := range p.Attendees {
			if v.Account == cj.Assignee {
				if err := t.putState(
					keyType(prefix+pid+"_"+stateProjectAttendees+"_"+strconv.Itoa(i)),
					stateAttendeeContribution,
					[]byte(strconv.FormatInt(v.Contribution+1, 10))); err != nil {
					return shim.Error(err.Error())
				}
				break
			}
		}
	}

	return shim.Success(nil)
}
