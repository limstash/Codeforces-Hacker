package token

import (
	"testing"
	"net/http"
)

func Test_GetCSRF(t *testing.T) {
	var GlobalCookie []*http.Cookie

	csrf, e := GetCSRF(&GlobalCookie)
	status := true

	if e != nil {
		if csrf != "" {
			t.Error("Module token return not null with error")
		}

		t.Error(e);
		status = false
	}

	if csrf == "" && (e == nil || e.Error() != "CSRF is an empty field") {
		t.Error("Module token failed in checking empty CSRF")
		status = false
	}

	if status == false {
		t.Log("Get CSRF Test Failed");
	}else{
		t.Log("Get CSRF Test Passed")
	}
}