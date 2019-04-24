package code

import (
	"testing"
	"net/http"
	"github.com/hytzongxuan/Codeforces-Hacker/module/token"
)


func Test_QueryCode(t *testing.T){
	var GlobalCookie []*http.Cookie
	
	CSRF, err := token.GetCSRF(&GlobalCookie)

	if err != nil {
		t.Skip("Fetched CSRF Failed, Skipped")
	}
	
	res, e := QueryCode(53127485, &GlobalCookie, CSRF)

	if e != nil {
		t.Error(e)
	}else if res == "" {
		t.Error("No Response from codeforces.com")
	}else if(GlobalCookie == nil || len(GlobalCookie) == 0){
		t.Error("Cookie is an empty field")
	}

	if e != nil || res == "" {
		t.Log("Query Code Test Failed")
	}else{
		t.Log("Query Code Test Passed")
	}
}