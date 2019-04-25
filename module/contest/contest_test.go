package contest

import (
	"net/http"
	"testing"
)

func Test_QueryContests(t *testing.T) {
	var GlobalCookie []*http.Cookie
	i, e := queryContests(&GlobalCookie)

	status := true

	if e != nil {
		t.Error(e)
		status = false
	} else if i == "" {
		t.Error("Empty Response from codeforces.com")
		status = false
	} else if GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Error("Cookie is en empty field")
		status = false
	}

	if status == false {
		t.Log("Query Contests Test Failed")
	} else {
		t.Log("Query Contests Test Passed")
	}
}

func Test_GetContests(t *testing.T) {
	var GlobalCookie []*http.Cookie
	i, e := GetContests(&GlobalCookie)

	status := true

	if e != nil {
		t.Error(e)
		status = false
	} else if i == nil {
		t.Error("Result is an empty field")
		status = false
	} else if GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Error("Cookie is en empty field")
		status = false
	}

	if status == false {
		t.Log("Get Contests Test Failed")
	} else {
		t.Log("Get Contests Test Passed")
	}
}

func Test_QueryProblems(t *testing.T) {
	var GlobalCookie []*http.Cookie

	i, e := queryProblems(556, &GlobalCookie)

	status := true

	if e != nil {
		t.Error(e)
		status = false
	} else if i == nil || len(i) == 0 {
		t.Error("Empty Response from codeforces.com")
		status = false
	} else if GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Error("Cookie is en empty field")
		status = false
	}

	if status == false {
		t.Log("Query Problems Test Failed")
	} else {
		t.Log("Query Problems Test Passed")
	}
}

func Test_GetProblems(t *testing.T) {
	var GlobalCookie []*http.Cookie
	i, e := GetProblems(556, &GlobalCookie)

	status := true

	if e != nil {
		t.Error(e)
		status = false
	} else if i == nil || len(i) == 0 {
		t.Error("Problems is an empty field")
		status = false
	} else if GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Error("Cookie is en empty field")
		status = false
	}

	if status == false {
		t.Log("Get Problems Test Failed")
	} else {
		t.Log("Get Problems Test Passed")
	}
}
