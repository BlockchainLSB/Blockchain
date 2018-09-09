package main

import (
	"errors"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const (
	prefix string = "pr_"
)

const (
	stateProjectID          string = "id"
	stateProjectAdmin              = "admin"
	stateProjectName               = "name"
	stateProjectLogo               = "logo"
	stateProjectCreatedAt          = "created_at"
	stateProjectDescription        = "description"
	stateProjectAttendees          = "attendees"
	stateProjectJobs               = "jobs"

	stateAttendeeName         = "name"
	stateAttendeeRole         = "role"
	stateAttendeeAffiliation  = "affiliation"
	stateAttendeeAccount      = "account"
	stateAttendeeContribution = "contribution"

	stateJobDone     = "done"
	stateJobAssignee = "assignee"
	stateJobSigners  = "signers"
)

type project struct {
	ID          string     `json:"id"`
	Admin       string     `json:"admin"`
	Logo        string     `json:"logo"`
	CreatedAt   int64      `json:"created_at"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Attendees   []Attendee `json:"attendees"`
	Jobs        []Job      `json:"jobs"`
}

func state(stub shim.ChaincodeStubInterface, key string, val interface{}) error {
	d, err := stub.GetState(key)

	if err != nil {
		return err
	}

	if string(d) == "" {
		return errors.New("Not found(" + key + ")")
	}

	switch v := val.(type) {
	case *int:
		*v, _ = strconv.Atoi(string(d))
	case *int64:
		*v, _ = strconv.ParseInt(string(d), 10, 64)
	case *string:
		*v = string(d)
	default:
		return errors.New("Unknown type")
	}

	return nil
}

func (p *project) load(stub shim.ChaincodeStubInterface, pid string) error {
	bk := prefix + pid + "_"
	v := []struct {
		key string
		val interface{}
	}{
		{
			key: stateProjectID,
			val: &p.ID,
		},
		{
			key: stateProjectName,
			val: &p.Name,
		},
		{
			key: stateProjectAdmin,
			val: &p.Admin,
		},
		{
			key: stateProjectLogo,
			val: &p.Logo,
		},
		{
			key: stateProjectCreatedAt,
			val: &p.CreatedAt,
		},
		{
			key: stateProjectDescription,
			val: &p.Description,
		},
	}

	for i := 0; i < len(v); i++ {
		if err := state(stub, bk+v[i].key, v[i].val); err != nil {
			return err
		}
	}

	atts, err := p.loadAttendees(stub, bk)

	if err != nil {
		return err
	}

	p.Attendees = atts

	jobs, err := p.loadJobs(stub, bk)

	if err != nil {
		return err
	}

	p.Jobs = jobs

	return nil
}

func (p *project) loadAttendees(stub shim.ChaincodeStubInterface, bk string) ([]Attendee, error) {
	var l int

	if err := state(stub, bk+stateProjectAttendees+"_len", &l); err != nil {
		return nil, err
	}

	res := make([]Attendee, l)

	bk = bk + stateProjectAttendees
	for i := 0; i < l; i++ {
		tk := bk + "_" + strconv.Itoa(i) + "_"
		state(stub, tk+stateAttendeeAccount, &res[i].Account)
		state(stub, tk+stateAttendeeRole, &res[i].Role)
		state(stub, tk+stateAttendeeAffiliation, &res[i].Affiliation)
		state(stub, tk+stateAttendeeContribution, &res[i].Contribution)
		state(stub, tk+stateAttendeeName, &res[i].Name)
	}

	return res, nil
}

func (p *project) loadJobs(stub shim.ChaincodeStubInterface, bk string) ([]Job, error) {
	var l int

	if err := state(stub, bk+stateProjectJobs+"_len", &l); err != nil {
		return nil, err
	}

	res := make([]Job, l)

	bk = bk + stateProjectJobs
	for i := 0; i < l; i++ {
		tk := bk + "_" + strconv.Itoa(i) + "_"
		state(stub, tk+stateJobDone, &res[i].Done)
		state(stub, tk+stateJobAssignee, &res[i].Assignee)

		var s int

		if err := state(stub, tk+stateJobSigners+"_len", &s); err != nil {
			return nil, err
		}

		ss := make([]string, s)

		for k := 0; k < s; k++ {
			stk := tk + stateJobSigners + "_" + strconv.Itoa(k)
			if err := state(stub, stk, &ss[k]); err != nil {
				return nil, err
			}
		}
		res[i].Signers = ss
	}

	return res, nil
}
