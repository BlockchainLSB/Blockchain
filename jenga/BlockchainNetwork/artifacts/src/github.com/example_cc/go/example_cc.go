package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type keyType string

var project_num int = 0

const (
	tokenType       keyType = "tk_"
	temporalType    keyType = "tmp_"
	projectType     keyType = "prj_"
	projectlistType keyType = "prjl_"
	portfolioType   keyType = "por_"
)

// UserChaincode example simple Chaincode implementation
type UserChaincode struct {
	stub     shim.ChaincodeStubInterface
	function string
	args     []string
}

type UserPortfolio struct {
	Id            string
	Pw            string
	Toeic         string
	Topcit        string
	Toeicspeaking string
	School        string
	Major         string
}

/*************** list *******************/
type userProjectList struct {
	Username string
	Projects []userPLists
}

type userPLists struct {
	Pnum  int
	Pname string
	PDes  string
	POk   bool
}

/***************** send *******************/
type sendpList struct {
	Pnum  int
	Pname string
	PDes  string
	POk   string
}

/***************** user project *******************/

type UserProject struct {
	Pnum           int
	Pname          string
	PDes           string
	POkContributor []projectContributor
	Pappraise      []projectAppraise
}

type projectContributor struct {
	Cname string
	C_ok  bool
}

type projectAppraise struct {
	Pindex    int                // 평가 고유 값
	Pevalname string             // 평가 받는 사람 이름
	PDes      string             // Pevalname의 기여한 내용
	Pteamlist []projectAppraiser // project 참여한 멤버의 리스트
}
type projectAppraiser struct {
	PAname string
	PAok   bool
}

/************************************/

type UserTransaction struct {
	TransactionInfo []TransactionInfo
}

type TransactionInfo struct {
	TxId      string
	Value     string
	Timestamp string
}

func userProjectListInit(pname string, pnum []int) userProjectList {
	upl := userProjectList{}
	userlists := []userPLists{}

	upl.Username = pname

	for i := 0; i < len(pnum); i++ {
		userlist := userPLists{}
		userlist.Pnum = pnum[i]
		userlist.POk = false
		userlists = append(userlists, userlist)
	}
	upl.Projects = userlists
	return upl
}

func projectListInit(pname string, pdes string) sendpList {
	pl := sendpList{}
	pl.Pname = pname
	pl.PDes = pdes

	return pl
}

func UserProjectInit(pnum int, pname string, pdes string, pokcontributor []string) UserProject {

	// Create struct and append it to the slice.
	up := UserProject{}
	pCs := []projectContributor{}
	pAs := []projectAppraiser{}
	pAes := []projectAppraise{}
	//pindex := 0

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

	//projectappraiser 배열 초기화
	for j := 0; j < len(pokcontributor); j++ {
		pAs = []projectAppraiser{}
		pAe := projectAppraise{}
		pAe.Pindex = j
		pAe.Pevalname = pokcontributor[j]
		pAe.PDes = "-1"

		for k := 0; k < len(pokcontributor); k++ {
			pA := projectAppraiser{}
			pA.PAname = pokcontributor[k]

			if k == j {
				pA.PAok = true
			} else {
				pA.PAok = false
			}
			pAs = append(pAs, pA)
		}
		pAe.Pteamlist = pAs
		pAes = append(pAes, pAe)
	}

	up.Pappraise = pAes

	return up
}
func (t *UserChaincode) call() pb.Response {
	function := t.function

	callMap := map[string]func() pb.Response{
		"signup":             t.signup,
		"signin":             t.signin,
		"getToken":           t.getToken,
		"addProject":         t.addProject,
		"loadProject":        t.loadProject,
		"searchProject":      t.searchProject,
		"loadProjectdetail":  t.loadProjectdetail,
		"acceptProject":      t.acceptProject,
		"searchUser":         t.searchUser,
		"getUserTransaction": t.getUserTransaction,
		"addContribution":    t.addContribution,
		"acceptContribution": t.acceptContribution,
		"requestedConlist":   t.requestedConlist,
		"allacceptedConlist": t.allacceptedConlist,
		"searchPortfolio":    t.searchPortfolio,
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

	if len(args) != 14 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	up := UserPortfolio{}
	data := UserPortfolio{}

	existPortfolio, err := t.stub.GetState(string(portfolioType) + string(args[1]))

	if err != nil {
		return shim.Error(string(args[1]) + " has been resistered. ")
	}

	_ = json.Unmarshal(existPortfolio, &data)

	if data.Id == args[1] {
		return shim.Error(string(args[1]) + " has been resistered. ")
	}

	up.Id = args[1]
	up.Pw = args[3]
	up.Toeic = args[5]
	up.Topcit = args[7]
	up.Toeicspeaking = args[9]
	up.School = args[11]
	up.Major = args[13]

	doc, _ := json.MarshalIndent(up, "", "    ")
	fmt.Println(string(doc))

	err = t.stub.PutState(string(portfolioType)+string(args[1]), doc)

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

	up := UserPortfolio{}
	// id로 pw 찾을 수 있음
	portfolio, err := t.stub.GetState(string(portfolioType) + id)
	_ = json.Unmarshal(portfolio, &up)

	if err != nil {
		return shim.Error(id + " is not registered.")
	}

	if up.Pw != pw {
		return shim.Error("Incorrect password! ")

	}

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
	up := UserPortfolio{}
	portfolio, err := t.stub.GetState(string(portfolioType) + idVal)
	_ = json.Unmarshal(portfolio, &up)

	if err != nil {
		return shim.Error(idVal + "is not registered.")
	}

	if up.Pw == pwVal {
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
	// id로 pw 찾기
	up := UserPortfolio{}
	portfolio, err := t.stub.GetState(string(portfolioType) + id)
	_ = json.Unmarshal(portfolio, &up)

	if err != nil {
		fmt.Println("id is not registered")
		return shim.Error(id + "is not registered.")
	}

	if len(up.Pw) != 0 {
		fmt.Println(id + " is registered")
	} else if len(up.Pw) == 0 {
		return shim.Error(id + " is not registered")
	}
	return shim.Success(nil)

}

func (t *UserChaincode) addProject() pb.Response {

	/*

	               {
	       "Pnum": 2,
	       "Pname": "job3",
	       "PDes": "amazing!",
	       "POkContributor": [
	           {
	               "Cname": "xxx",
	               "C_ok": true
	           },
	           {
	               "Cname": "zzz",
	               "C_ok": false
	           }
	       ],
	       "Pappraise": [
	           {
	               "Pindex": 0,
	               "Pevalname": "xxx",
	               "PDes": "I am good programmer",
	               "Pteamlist": [
	                   {
	                       "PAname": "xxx",
	                       "PAok": true
	                   },
	                   {
	                       "PAname": "zzz",
	                       "PAok": false
	                   }
	               ]
	           },
	           {
	               "Pindex": 1,
	               "Pevalname": "zzz",
	               "PDes": "I am good programmer",
	               "Pteamlist": [
	                   {
	                       "PAname": "xxx",
	                       "PAok": false
	                   },
	                   {
	                       "PAname": "zzz",
	                       "PAok": true
	                   }
	               ]
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
	userid, err := t.getuserInfo(token)

	if err != nil {
		return shim.Error(err.Error())
	}

	//project 이름이 같은 지 확인
	for {
		var data UserProject
		existProject, perr := t.stub.GetState(string(projectType) + string(project_num))

		if perr != nil {
			return shim.Error(err.Error())
		}
		_ = json.Unmarshal(existProject, &data)

		if len(data.Pname) == 0 {
			break
		}
		if data.Pname == projectname {
			return shim.Error(projectname + " was already added")
		}

		project_num = project_num + 1
		fmt.Println("project_num :", project_num)
	}

	//기존에 존재했던 project들과 이름이 같지 않을 떄
	p := UserProjectInit(project_num, projectname, description, contributorArray)
	doc, jerr := json.MarshalIndent(p, "", "    ")

	if jerr != nil {
		return shim.Error("json Marshal Error")
	}

	fmt.Println(string(doc))

	err = t.stub.PutState(string(projectType)+string(project_num), doc)

	if err != nil {
		return shim.Error(err.Error())
	}

	/****************** user project list ******************/

	for i := 0; i < len(contributorArray); i++ {
		upl := userProjectList{}
		existProjectList, eperr := t.stub.GetState(string(projectlistType) + string(contributorArray[i]))
		if eperr != nil {
			return shim.Error(err.Error())
		}
		userlist := userPLists{}
		_ = json.Unmarshal(existProjectList, &upl)

		userlist.Pnum = project_num
		userlist.Pname = projectname
		userlist.PDes = description

		if contributorArray[i] == string(userid) {
			userlist.POk = true
		} else {
			userlist.POk = false
		}

		upl.Username = contributorArray[i]
		upl.Projects = append(upl.Projects, userlist)

		ldoc, _ := json.MarshalIndent(upl, "", "    ")
		fmt.Println(contributorArray[i] + "'s project list :" + string(ldoc))

		uerr := t.stub.PutState(string(projectlistType)+string(contributorArray[i]), ldoc)
		if uerr != nil {
			return shim.Error(err.Error())
		}
	}

	return shim.Success(nil)

}

func (t *UserChaincode) loadProject() pb.Response {

	/*
	            user project list  :  {
	       "Username": "xxx",
	       "Projects": [
	           {
	               "Pnum": 0,
	               "Pname": "job1",
	               "PDes": "wow!",
	               "POk": false
	           },
	           {
	               "Pnum": 1,f
	               "Pname": "job2",
	               "PDes": "amazing!",
	               "POk": false
	           },
	           {
	               "Pnum": 2,
	               "Pname": "job3",
	               "PDes": "amazing!",
	               "POk": true
	           }
	       ]
	   }
	*/

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
	upl := userProjectList{}

	fmt.Println("user id : ", string(userid))
	existProjectList, _ := t.stub.GetState(string(projectlistType) + string(userid))
	_ = json.Unmarshal(existProjectList, &upl)

	ldoc, _ := json.MarshalIndent(upl, "", "    ")

	fmt.Println(string(ldoc))

	return shim.Success([]byte(ldoc))
}

func (t *UserChaincode) searchProject() pb.Response {
	fmt.Println("========================= searchProject =========================")
	args := t.args

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	userid := args[1]

	upl := userProjectList{}
	existProjectList, _ := t.stub.GetState(string(projectlistType) + string(userid))
	_ = json.Unmarshal(existProjectList, &upl)

	doc, _ := json.MarshalIndent(upl, "", "    ")
	fmt.Println(string(doc))

	return shim.Success([]byte(doc))

}
func (t *UserChaincode) loadProjectdetail() pb.Response {
	fmt.Println("========================= loadProjectdetail =========================")
	args := t.args
	// pnum

	pnum := args[1]

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	i, _ := strconv.Atoi(pnum)

	up := UserProject{}

	existProject, _ := t.stub.GetState(string(projectType) + string(i))
	_ = json.Unmarshal(existProject, &up)

	doc, _ := json.MarshalIndent(up, "", "    ")

	fmt.Println(string(doc))

	return shim.Success([]byte(doc))

}
func (t *UserChaincode) acceptProject() pb.Response {
	fmt.Println("========================= acceptProject =========================")
	args := t.args
	//웹 -> cc : token, pnum

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	token := args[1]
	pnum := args[3]

	//token으로 id 찾기
	userid, err := t.stub.GetState(string(temporalType) + string(token))

	if err != nil {
		return shim.Error(err.Error())
	}
	var Pstruct, newPstruct UserProject

	i, _ := strconv.Atoi(pnum)
	doc, derr := t.stub.GetState(string(projectType) + string(i))
	_ = json.Unmarshal(doc, &Pstruct)
	fmt.Println("doc : ", string(doc))

	if derr != nil {
		return shim.Error(err.Error())
	}

	newPstruct.Pnum = Pstruct.Pnum
	newPstruct.PDes = Pstruct.PDes
	newPstruct.Pname = Pstruct.Pname

	flag := 0
	tmps := []projectContributor{}
	tmp := projectContributor{}

	for _, v := range Pstruct.POkContributor {
		if flag == 0 {
			//master는 무조건 그대로
			tmp.C_ok = true
			tmp.Cname = v.Cname
		} else {
			if v.Cname == string(userid) {
				tmp.C_ok = true
				tmp.Cname = string(userid)
			} else {
				if v.C_ok {
					tmp.C_ok = true
				} else {
					tmp.C_ok = false
				}
				tmp.Cname = v.Cname
			}
		}
		tmps = append(tmps, tmp)
		flag = 1
	}
	newPstruct.POkContributor = tmps
	newPstruct.Pappraise = Pstruct.Pappraise
	doc, _ = json.MarshalIndent(newPstruct, "", "    ")

	err = t.stub.PutState(string(projectType)+string(i), doc)

	if err != nil {
		return shim.Error(err.Error())
	}

	/********* accept 후 userProjectList의 Projects의 POk 바꾸기 ***********/
	//웹 -> cc : token, pnum
	doc, _ = t.stub.GetState(string(projectlistType) + string(userid))

	prjlist := userProjectList{}
	newprjlist := userProjectList{}
	prjdetail := []userPLists{}
	_ = json.Unmarshal(doc, &prjlist)

	newprjlist.Username = prjlist.Username

	for _, v := range prjlist.Projects {
		tmp := userPLists{}

		tmp.Pnum = v.Pnum
		tmp.Pname = v.Pname
		tmp.PDes = v.PDes
		if v.Pnum == i {
			tmp.POk = true
		} else {
			if v.POk {
				tmp.POk = true
			} else {
				tmp.POk = false
			}
		}
		prjdetail = append(prjdetail, tmp)
	}

	newprjlist.Projects = prjdetail

	listdataJson, _ := json.MarshalIndent(newprjlist, "", "    ")
	err = t.stub.PutState(string(projectlistType)+string(userid), listdataJson)
	fmt.Println(string(listdataJson))

	return shim.Success([]byte("{ \"is_ok\": true }"))

}

func (t *UserChaincode) getuserInfo(key string) ([]byte, error) {
	return t.stub.GetState(string(temporalType) + key)
}

func (t *UserChaincode) getUserTransaction() pb.Response {
	fmt.Println("=======getTransaction============")

	if len(t.args) != 2 {
		return shim.Error("Incorrect # of arguments. Expecting 0")
	}

	token := t.args[1]
	id, err := t.getuserInfo(token)
	if err != nil {
		return shim.Error("Not Invalid token")
	}
	fmt.Println("id: " + string(id))
	testresult, terr := t.stub.GetState(string(projectlistType) + string(id))
	if terr != nil {
		return shim.Error("get state error")
	}
	fmt.Println("testresult: " + string(testresult))
	resultsIterator, err := t.stub.GetHistoryForKey(string(projectlistType) + string(id))
	if err != nil {
		return shim.Error("get history for key error")
	}
	var transactionInfos []TransactionInfo
	defer resultsIterator.Close()
	fmt.Println("Before results Iterator")
	for resultsIterator.HasNext() {
		fmt.Println("In results Iterator")
		var transactionInfo TransactionInfo
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		transactionInfo.TxId = response.TxId
		fmt.Println("transactionInfo.TxId: " + response.TxId)
		if response.IsDelete {
			transactionInfo.Value = "null"
		} else {
			transactionInfo.Value = string(response.Value)
		}
		fmt.Println("transactionInfo.Value: " + string(response.Value))
		transactionInfo.Timestamp = time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String()
		fmt.Println("transactionInfo.Timestamp: " + transactionInfo.Timestamp)
		transactionInfos = append(transactionInfos, transactionInfo)
	}
	var userTransaction UserTransaction
	userTransaction.TransactionInfo = transactionInfos

	jResult, jerr := json.Marshal(userTransaction)
	if jerr != nil {
		return shim.Error("json Marshal Error")
	}
	fmt.Println("jResult string: " + string(jResult))

	return shim.Success([]byte(jResult))
}

func (t *UserChaincode) addContribution() pb.Response {
	// 웹 -> cc : token, Pnum, PDes

	fmt.Println("========================= addContribution =========================")

	args := t.args

	token := args[1]
	pnum := args[3]
	pdes := args[5]

	//token 값으로 id 값 받아오기
	userid, err := t.getuserInfo(token)

	if err != nil {
		return shim.Error(err.Error())
	}

	i, _ := strconv.Atoi(pnum)

	var Pstruct, newPstruct UserProject
	Pappstruct := []projectAppraise{}

	doc, _ := t.stub.GetState(string(projectType) + string(i))
	_ = json.Unmarshal(doc, &Pstruct)

	newPstruct.Pnum = Pstruct.Pnum
	newPstruct.Pname = Pstruct.Pname
	newPstruct.PDes = Pstruct.PDes
	newPstruct.POkContributor = Pstruct.POkContributor

	for _, v := range Pstruct.Pappraise {
		tmp := projectAppraise{}

		tmp.Pindex = v.Pindex
		tmp.Pevalname = v.Pevalname

		if v.Pevalname == string(userid) {
			tmp.PDes = pdes
		} else {
			tmp.PDes = v.PDes
		}
		tmp.Pteamlist = v.Pteamlist
		Pappstruct = append(Pappstruct, tmp)
	}
	newPstruct.Pappraise = Pappstruct

	jdoc, _ := json.MarshalIndent(newPstruct, "", "    ")
	err = t.stub.PutState(string(projectType)+string(i), jdoc)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println(string(jdoc))

	return shim.Success(nil)

}

func (t *UserChaincode) acceptContribution() pb.Response {
	fmt.Println("========================= acceptContribution =========================")
	args := t.args

	// - 웹 -> cc : token, pid, index(eval)
	// cc -> 웹 : true 로 바꾸기

	token := args[1]
	pnum := args[3]
	pindex := args[5]

	//token으로 id 찾기
	userid, err := t.stub.GetState(string(temporalType) + string(token))

	var Pstruct, newPstruct UserProject
	tmps := []projectAppraise{}

	i, _ := strconv.Atoi(pnum)
	j, _ := strconv.Atoi(pindex)
	doc, derr := t.stub.GetState(string(projectType) + string(i))
	_ = json.Unmarshal(doc, &Pstruct)
	fmt.Println("doc : ", string(doc))

	if derr != nil {
		return shim.Error(err.Error())
	}

	newPstruct.Pnum = Pstruct.Pnum
	newPstruct.Pname = Pstruct.Pname
	newPstruct.PDes = Pstruct.PDes
	newPstruct.POkContributor = Pstruct.POkContributor

	for _, v := range Pstruct.Pappraise {
		tmp := projectAppraise{}
		ts := []projectAppraiser{}
		tmp.Pindex = v.Pindex
		tmp.Pevalname = v.Pevalname
		tmp.PDes = v.PDes

		if v.Pindex == j {
			for _, j := range v.Pteamlist {
				t := projectAppraiser{}
				if j.PAname == string(userid) {
					t.PAok = true
				} else {
					if j.PAok {
						t.PAok = true
					} else {
						t.PAok = false
					}

				}
				t.PAname = j.PAname
				ts = append(ts, t)
			}
			tmp.Pteamlist = ts

		} else {
			tmp.Pteamlist = v.Pteamlist
		}

		tmps = append(tmps, tmp)
	}

	newPstruct.Pappraise = tmps

	doc, _ = json.MarshalIndent(newPstruct, "", "    ")
	err = t.stub.PutState(string(projectType)+string(i), doc)

	fmt.Println(string(doc))

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *UserChaincode) requestedConlist() pb.Response {
	fmt.Println("========================= requestedConList =========================")

	args := t.args

	// token, pnum

	// appraiser 중 pevalname 이 나의 아이디와 다를 때,
	// 즉 Pteamlist에 내가 있는데 PAok가 false일때의 리스트

	token := args[1]
	pnum := args[3]
	//token으로 id 찾기
	userid, err := t.stub.GetState(string(temporalType) + string(token))

	if err != nil {
		return shim.Error(err.Error())
	}

	var Pstruct UserProject
	Pappraises := []projectAppraise{}

	i, _ := strconv.Atoi(pnum)
	doc, derr := t.stub.GetState(string(projectType) + string(i))
	_ = json.Unmarshal(doc, &Pstruct)
	//fmt.Println("doc : ", string(doc))

	if derr != nil {
		return shim.Error(err.Error())
	}

	flag := 0
	for _, v := range Pstruct.Pappraise {
		flag = 0
		Pappraise := projectAppraise{}
		appraisers := []projectAppraiser{}
		if v.PDes != "-1" {
			for _, j := range v.Pteamlist {
				appraiser := projectAppraiser{}
				if j.PAname == string(userid) {
					if !j.PAok {
						flag = 1
						appraiser.PAname = string(userid)
						appraiser.PAok = false
						appraisers = append(appraisers, appraiser)
					}
				}
			}
			if flag == 1 {
				Pappraise.Pindex = v.Pindex
				Pappraise.Pevalname = v.Pevalname
				Pappraise.PDes = v.PDes
				Pappraise.Pteamlist = appraisers

				Pappraises = append(Pappraises, Pappraise)
			}
		}

	}

	doc, _ = json.MarshalIndent(Pappraises, "", "    ")

	fmt.Println(string(doc))

	return shim.Success([]byte(doc))

}

func (t *UserChaincode) allacceptedConlist() pb.Response {
	fmt.Println("========================= allacceptedConlist =========================")

	args := t.args

	// pnum
	pnum := args[1]

	var Pstruct UserProject
	Pappraises := []projectAppraise{}

	i, _ := strconv.Atoi(pnum)
	doc, derr := t.stub.GetState(string(projectType) + string(i))
	_ = json.Unmarshal(doc, &Pstruct)

	if derr != nil {
		return shim.Error(derr.Error())
	}

	flag := 0
	for _, v := range Pstruct.Pappraise {
		flag = 0
		Pappraise := projectAppraise{}

		for _, j := range v.Pteamlist {
			if !j.PAok {
				flag = 1
				break
			}
		}
		if flag == 0 {
			Pappraise.Pindex = v.Pindex
			Pappraise.PDes = v.PDes
			Pappraise.Pevalname = v.Pevalname
			Pappraise.Pteamlist = v.Pteamlist
			Pappraises = append(Pappraises, Pappraise)
		}

	}

	doc, _ = json.MarshalIndent(Pappraises, "", "    ")

	fmt.Println(string(doc))

	return shim.Success([]byte(doc))

}
func (t *UserChaincode) searchPortfolio() pb.Response {
	fmt.Println("========================= searchPortfolio =========================")
	args := t.args

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	userid := args[1]

	up := UserPortfolio{}
	existPortfolio, _ := t.stub.GetState(string(portfolioType) + string(userid))
	_ = json.Unmarshal(existPortfolio, &up)

	doc, _ := json.MarshalIndent(up, "", "    ")
	fmt.Println(string(doc))

	return shim.Success([]byte(doc))

}
func main() {
	err := shim.Start(new(UserChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
