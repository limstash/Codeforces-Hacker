package app

import (
	"testing"
	"net/http"
)

func Test_Load(t *testing.T){
	var Cookie []*http.Cookie
	contest, CSRF, e := Load(&Cookie)

	if e != nil {
		t.Error(e);
	}else if CSRF == "" {
		t.Error("CSRF is an empty field")
	}else if contest == nil || len(contest) == 0 {
		t.Error("Contest is an empty field")
	}

	if e != nil || CSRF == "" || contest == nil || len(contest) == 0 {
		t.Error("App Loading Test Failed")
	}else{
		t.Log("App Loading Test Passed")
	}
}