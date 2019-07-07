package judge

import (
	"testing"

	. "github.com/limstash/Codeforces-Hacker/common"
	"github.com/limstash/Codeforces-Hacker/module/code"
	"github.com/limstash/Codeforces-Hacker/module/token"
)

func CompileCase01(t *testing.T, auth Authentication) bool {
	a, b, c, err := getGPlusPlusVersion()

	if checkGCC(a, b, c) < 1 {
		t.Log("Test Skipped: (Case 01) Not Supported (C++11)")
	}

	submission := SubmissionCPP11()

	submitcode, err := code.QueryCode(submission.SubmissionID, &auth, "https://codeforces.com")

	if err != nil {
		t.Log(err)
		t.Skip("Test Failed: (Case 01) Query code failed (C++11)")
	}

	submission.Code = submitcode

	err = code.SaveCode(submission)

	if err != nil {
		t.Log(err)
		t.Skip("Test Failed: (Case 01) Save code failed (C++11)")
	}

	res, err := Compile(submission)

	if err != nil || res == false {
		t.Error("Test Failed: (Case 01) Compile failed (C++11) ")
		t.Error(err)
		return false
	}

	return true
}

func CompileCase02(t *testing.T, auth Authentication) bool {
	a, b, c, err := getGPlusPlusVersion()

	if checkGCC(a, b, c) < 2 {
		t.Log("Test Skipped: (Case 02) Not Supported (C++14)")
	}

	submission := SubmissionCPP14()

	submitcode, err := code.QueryCode(submission.SubmissionID, &auth, "https://codeforces.com")

	if err != nil {
		t.Log(err)
		t.Skip("Test Failed: (Case 02) Query code failed (C++14)")
	}

	submission.Code = submitcode

	err = code.SaveCode(submission)

	if err != nil {
		t.Log(err)
		t.Skip("Test Failed: (Case 02) Save code failed (C++14)")
	}

	res, err := Compile(submission)

	if err != nil || res == false {
		t.Error("Test Failed: (Case 02) Compile failed (C++14) ")
		t.Error(err)
		return false
	}

	return true
}

func CompileCase03(t *testing.T, auth Authentication) bool {
	a, b, c, err := getGPlusPlusVersion()

	if checkGCC(a, b, c) < 3 {
		t.Log("Test Skipped: (Case 03) Not Supported (C++17)")
	}

	submission := SubmissionCPP17()

	submitcode, err := code.QueryCode(submission.SubmissionID, &auth, "https://codeforces.com")

	if err != nil {
		t.Log(err)
		t.Skip("Test Failed: (Case 03) Query code failed (C++17)")
	}

	submission.Code = submitcode

	err = code.SaveCode(submission)

	if err != nil {
		t.Log(err)
		t.Skip("Test Failed: (Case 03) Save code failed (C++17)")
	}

	res, err := Compile(submission)

	if err != nil || res == false {
		t.Error("Test Failed: (Case 03) Compile failed (C++17) ")
		t.Error(err)
		return false
	}

	return true
}

func CompileCase04(t *testing.T, auth Authentication) bool {
	a, b, c, err := getGCCVersion()

	if checkGCC(a, b, c) < 1 {
		t.Log("Test Skipped: (Case 04) Not Supported (C11)")
	}

	submission := SubmissionC11()

	submitcode, err := code.QueryCode(submission.SubmissionID, &auth, "https://codeforces.com")

	if err != nil {
		t.Log(err)
		t.Skip("Test Failed: (Case 04) Query code failed (C11)")
	}

	submission.Code = submitcode

	err = code.SaveCode(submission)

	if err != nil {
		t.Log(err)
		t.Skip("Test Failed: (Case 04) Save code failed (C11)")
	}

	res, err := Compile(submission)

	if err != nil || res == false {
		t.Error("Test Failed: (Case 04) Compile failed (C11) ")
		t.Error(err)
		return false
	}

	return true
}

func CompileCase05(t *testing.T, auth Authentication) bool {

	if checkGo() == false {
		t.Log("Test Skipped: (Case 05) Not Supported (Go)")
	}

	submission := SubmissionGo()

	submitcode, err := code.QueryCode(submission.SubmissionID, &auth, "https://codeforces.com")

	if err != nil {
		t.Log(err)
		t.Skip("Test Failed: (Case 05) Query code failed (Go)")
	}

	submission.Code = submitcode

	err = code.SaveCode(submission)

	if err != nil {
		t.Log(err)
		t.Skip("Test Failed: (Case 05) Save code failed (Go)")
	}

	res, err := Compile(submission)

	if err != nil || res == false {
		t.Error("Test Failed: (Case 05) Compile failed (Go) ")
		t.Error(err)
		return false
	}

	return true
}

func Test_Compile(t *testing.T) {
	auth := Authentication{}
	err := token.GetCSRF(&auth, "https://codeforces.com")

	if err != nil {
		t.Skip("Fetch CSRF failed, skipped")
	}

	status := true

	if CompileCase01(t, auth) == false {
		status = false
	}

	if CompileCase02(t, auth) == false {
		status = false
	}

	if CompileCase03(t, auth) == false {
		status = false
	}

	if CompileCase04(t, auth) == false {
		status = false
	}

	if CompileCase05(t, auth) == false {
		status = false
	}

	if status == true {
		t.Log("Package judge - Compile test passed")
	} else {
		t.Log("Package judge - Compile test failed")
	}
}
