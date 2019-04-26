package code

import (
	"net/http"
	"os"
	"runtime"
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

func Test_Save(t *testing.T) {

	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		os.MkdirAll(getPath()+"/src/53319035", 0777)
	} else {
		os.MkdirAll(getPath()+"\\src\\53319035", 0777)
	}

	status := true

	e := SaveCode(53319035, "GNU C++11", "test")

	if e != nil {
		status = false
		t.Error("Module code fail at saving GNU C++11 Code")
		t.Error(e)
	}

	e = SaveCode(53319035, "GNU C++14", "test")

	if e != nil {
		status = false
		t.Error("Module code fail at saving GNU C++14 Code")
		t.Error(e)
	}

	e = SaveCode(53319035, "GNU C++17", "test")

	if e != nil {
		status = false
		t.Error("Module code fail at saving GNU C++17 Code")
		t.Error(e)
	}

	e = SaveCode(53319035, "GNU C11", "test")

	if e != nil {
		status = false
		t.Error("Module code fail at saving GNU C11 Code")
		t.Error(e)
	}

	e = SaveCode(53319035, "Python2", "test")

	if e != nil {
		status = false
		t.Error("Module code fail at saving Python2 Code")
		t.Error(e)
	}

	e = SaveCode(53319035, "Python3", "test")

	if e != nil {
		status = false
		t.Error("Module code fail at saving Python3 Code")
		t.Error(e)
	}

	e = SaveCode(53319035, "Go", "test")

	if e != nil {
		status = false
		t.Error("Module code fail at saving Go Code")
		t.Error(e)
	}

	if status == true {
		t.Log("Save Code Test Passed")
	} else {
		t.Log("Save Code Test Failed")
	}
}
