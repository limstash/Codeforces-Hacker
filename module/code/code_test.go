package code

import (
	"net/http"
	"testing"

	"github.com/hytzongxuan/Codeforces-Hacker/module/token"
)

func Test_QueryCode(t *testing.T) {
	var GlobalCookie []*http.Cookie

	CSRF, err := token.GetCSRF(&GlobalCookie)

	if err != nil {
		t.Skip("Fetched CSRF Failed, Skipped")
	}

	res, e := QueryCode(53127485, &GlobalCookie, CSRF)
	status := true

	if e != nil {
		if res != "" {
			t.Error("Module code return not null with error")
		}
		t.Error(e)
		status = false
	} else if res == "" {
		t.Error("No Response from codeforces.com")
		status = false
	} else if GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Error("Cookie is an empty field")
		status = false
	}

	if status == false {
		t.Log("Query Code Test Failed")
	} else {
		t.Log("Query Code Test Passed")
	}
}
