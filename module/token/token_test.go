package token

import (
	"testing"

	. "github.com/hytzongxuan/Codeforces-Hacker/common"
)

func Test_GetCSRF(t *testing.T) {
	status := true
	authentication := Authentication{}

	err := GetCSRF(&authentication, "https://codeforces.com")

	if err != nil {
		if authentication.CSRF != "" {
			t.Error("Test failed: GetCSRF throw an error even if CSRF exists")
		}

		t.Error(err)
		status = false
	}

	if authentication.CSRF == "" && (err == nil || err.Error() != "CSRF not exists") {
		t.Error("Test failed: GetCSRF should throw an error here (CSRF not exists)")
		status = false
	}

	if status == true {
		t.Log("Package token - GetCSRF test passed")
	} else {
		t.Log("Package token - GetCSRF test failed")
	}
}
