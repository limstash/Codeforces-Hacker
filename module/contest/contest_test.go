package contest
	
import (
	"testing"
	"net/http"
)

func Test_QueryContests(t *testing.T) {
	var GlobalCookie []*http.Cookie
	i, e := queryContests(&GlobalCookie)
	
	if e != nil {
		t.Error(e)
	}else if i == "" {
		t.Error("Empty Response from codeforces.com")
	}else if GlobalCookie == nil || len(GlobalCookie) == 0{
		t.Error("Cookie is en empty field")
	}

	if e != nil || i == "" || GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Log("Query Contests Test Failed")
	}else{
		t.Log("Query Contests Test Passed")
	}
}

func Test_GetContests(t *testing.T){
	var GlobalCookie []*http.Cookie
	i, e := GetContests(&GlobalCookie)

	if e != nil {
		t.Error(e)
	}else if i == nil {
		t.Error("Result is an empty field")
	}else if GlobalCookie == nil || len(GlobalCookie) == 0{
		t.Error("Cookie is en empty field")
	}

	if e != nil || i == nil || GlobalCookie == nil || len(GlobalCookie) == 0{
		t.Log("Get Contests Test Failed")
	}else{
		t.Log("Get Contests Test Passed")
	}
}

func Test_QueryProblems(t *testing.T){
	var GlobalCookie []*http.Cookie

	i, e := queryProblems(556, &GlobalCookie)

	if e != nil {
		t.Error(e)
	}else if i == nil || len(i) == 0{
		t.Error("Empty Response from codeforces.com")
	}else if GlobalCookie == nil || len(GlobalCookie) == 0{
		t.Error("Cookie is en empty field")
	}

	if e != nil || i == nil || len(i) == 0 || GlobalCookie == nil || len(GlobalCookie) == 0{
		t.Log("Query Problems Test Failed")
	}else{
		t.Log("Query Problems Test Passed")
	}
}

func Test_GetProblems(t *testing.T){
	var GlobalCookie []*http.Cookie
	i, e := GetProblems(556, &GlobalCookie)

	if e != nil {
		t.Error(e)
	}else if i == nil || len(i) == 0{
		t.Error("Problems is an empty field")
	}else if GlobalCookie == nil || len(GlobalCookie) == 0{
		t.Error("Cookie is en empty field")
	}

	if e != nil || i == nil || len(i) == 0 || GlobalCookie == nil || len(GlobalCookie) == 0{
		t.Log("Get Problems Test Failed")
	}else{
		t.Log("Get Problems Test Passed")
	}
}