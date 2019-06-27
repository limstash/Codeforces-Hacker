package code

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	. "github.com/hytzongxuan/Codeforces-Hacker/common"
	"github.com/hytzongxuan/Codeforces-Hacker/module/token"
)

func Test_QueryCode(t *testing.T) {
	status := true

	auth := Authentication{}
	err := token.GetCSRF(&auth, "https://codeforces.com")

	res, err := QueryCode(56065447, &auth, "https://codeforces.com")

	if err != nil {
		t.Error(err)
		status = false
	}

	if len(res) <= 10 {
		t.Error("Test Failed: the string return by QueryCode is too short")
		status = false
	}

	if status == true {
		t.Log("Package code - QueryCode test passed")
	} else {
		t.Log("Package code - QueryCode test failed")
	}
}

func getPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	rst := filepath.Dir(path)
	return rst
}

func Test_SaveCode(t *testing.T) {
	status := true

	os.MkdirAll(getPath()+"/src/53319035", 0777)

	auth := Authentication{}
	err := token.GetCSRF(&auth, "https://codeforces.com")

	if err != nil {
		t.Skip("Fetch CSRF failed, skipped")
	}

	config := Config{}
	config.Settings.Path = getPath()

	submission := Submission{}
	submission.SubmissionID = 53319035

	code, err := QueryCode(53319035, &auth, "https://codeforces.com")
	submission.Code = code

	submission.Language = "GNU C++11"
	e := SaveCode(submission, config)

	if e != nil {
		status = false
		t.Error("Test Failed: (Case 01) SaveCode failed at saving GNU C++11 Code")
		t.Error(e)
	}

	submission.Language = "GNU C++11"
	e = SaveCode(submission, config)

	if e != nil {
		status = false
		t.Error("Test Failed: (Case 02) SaveCode failed at saving GNU C++14 Code")
		t.Error(e)
	}

	submission.Language = "GNU C++17"
	e = SaveCode(submission, config)

	if e != nil {
		status = false
		t.Error("Test Failed: (Case 03) SaveCode failed at saving GNU C++17 Code")
		t.Error(e)
	}

	submission.Language = "GNU C11"
	e = SaveCode(submission, config)

	if e != nil {
		status = false
		t.Error("Test Failed: (Case 04) SaveCode failed at saving GNU C11 Code")
		t.Error(e)
	}

	submission.Language = "Python2"
	e = SaveCode(submission, config)

	if e != nil {
		status = false
		t.Error("Test Failed: (Case 05) SaveCode failed at saving Python2 Code")
		t.Error(e)
	}

	submission.Language = "Python3"
	e = SaveCode(submission, config)

	if e != nil {
		status = false
		t.Error("Test Failed: (Case 06) SaveCode failed at saving Python3 Code")
		t.Error(e)
	}

	submission.Language = "Go"
	e = SaveCode(submission, config)

	if e != nil {
		status = false
		t.Error("Test Failed: (Case 07) SaveCode failed at saving Golang Code")
		t.Error(e)
	}

	if status == true {
		t.Log("Package code - SaveCode test passed")
	} else {
		t.Log("Package code - SaveCode test failed")
	}
}
