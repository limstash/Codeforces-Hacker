package token

import (
	"testing"
	"net/http"
)

func Test_GetCSRF(t *testing.T) {
	var GlobalCookie []*http.Cookie

	csrf, e := GetCSRF(&GlobalCookie)

	if e != nil {
		t.Error(e);
	}

	if e != nil || csrf == "" {
		t.Log("Get CSRF Test Failed");
	}else{
		t.Log("Get CSRF Test Passed")
	}
}