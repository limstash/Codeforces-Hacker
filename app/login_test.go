package app

import (
	"testing"

	. "github.com/limstash/Codeforces-Hacker/common"
	"github.com/limstash/Codeforces-Hacker/module/token"
)

func testLogin01(t *testing.T, server string) bool {
	auth := Authentication{}
	err := token.GetCSRF(&auth, server)

	if err != nil {
		t.Skip("Fetch CSRF failed, skipped")
	}

	account := Account{}
	account.Username = ""
	account.Password = ""

	status, err := submitLogin(account, &auth, server)

	if status == true || err == nil {
		t.Error("Test Failed: (Case 01) submitLogin should throw an error here (Specify correct handle or email)")
		return false
	}

	if err.Error() != "Specify correct handle or email" {
		t.Error("Test Failed: (Case 01) submitLogin throw an unexpected error here (expected Specify correct handle or email)")
		t.Error(err)
		return false
	}

	return true
}

func testLogin02(t *testing.T, server string) bool {
	auth := Authentication{}
	err := token.GetCSRF(&auth, server)

	if err != nil {
		t.Skip("Fetch CSRF failed, skipped")
	}

	account := Account{}
	account.Username = "test02"
	account.Password = "test02"

	status, err := submitLogin(account, &auth, server)

	if status == true || err == nil {
		t.Error("Test Failed: (Case 02) submitLogin should throw an error here (Invalid handle/email or password)")
		return false
	}

	if err.Error() != "Invalid handle/email or password" {
		t.Error("Test Failed: (Case 01) submitLogin throw an unexpected error here (expected Invalid handle/email or password)")
		t.Error(err)
	}

	return true
}

func testLogin03(t *testing.T, server string) bool {
	auth := Authentication{}
	err := token.GetCSRF(&auth, server)

	if err != nil {
		t.Skip("Fetch CSRF failed, skipped")
	}

	account := Account{}
	account.Username = "test01"
	account.Password = "test01"

	status, err := submitLogin(account, &auth, server)

	if status == false || err != nil {
		t.Error(err)
		return false
	}

	return true
}

func Test_submitLogin(t *testing.T) {
	status := true

	if testLogin01(t, "https://codeforces.com") == false {
		status = false
	}

	if testLogin02(t, "https://codeforces.com") == false {
		status = false
	}

	if testLogin03(t, "https://codeforces.com") == false {
		status = false
	}

	if status == true {
		t.Log("Package app - submitLogin test passed")
	} else {
		t.Log("Package app - submitLogin test failed")
	}
}
