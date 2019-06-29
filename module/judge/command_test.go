package judge

import (
	"testing"

	. "github.com/hytzongxuan/Codeforces-Hacker/common"
)

func SubmissionCPP11() Submission {
	submission := Submission{}
	submission.Language = "GNU C++11"
	submission.Path = "./src/56192359"
	submission.SubmissionID = 56192359

	return submission
}

func SubmissionCPP14() Submission {
	submission := Submission{}
	submission.Language = "GNU C++14"
	submission.Path = "./src/56192351"
	submission.SubmissionID = 56192351

	return submission
}

func SubmissionCPP17() Submission {
	submission := Submission{}
	submission.Language = "GNU C++17"
	submission.Path = "./src/56192374"
	submission.SubmissionID = 56192374

	return submission
}

func SubmissionC11() Submission {
	submission := Submission{}
	submission.Language = "GNU C11"
	submission.Path = "./src/56192830"
	submission.SubmissionID = 56192830

	return submission
}

func SubmissionPython2() Submission {
	submission := Submission{}
	submission.Language = "Python 2"
	submission.Path = "./src/55525058"
	submission.SubmissionID = 55525058

	return submission
}

func SubmissionPython3() Submission {
	submission := Submission{}
	submission.Language = "Python 3"
	submission.Path = "./src/56192367"
	submission.SubmissionID = 56192367

	return submission
}

func SubmissionGo() Submission {
	submission := Submission{}
	submission.Language = "Go"
	submission.Path = "./src/40507899"
	submission.SubmissionID = 40507899

	return submission
}

func SubmissionUnknown() Submission {
	submission := Submission{}
	submission.Language = "Unknown"
	submission.Path = "./src/12345678"
	submission.SubmissionID = 12345678

	return submission
}

func Test_GetCompileCommand(t *testing.T) {
	status := true

	case01 := GetCompileCommand(SubmissionCPP11())
	case02 := GetCompileCommand(SubmissionCPP14())
	case03 := GetCompileCommand(SubmissionCPP17())
	case04 := GetCompileCommand(SubmissionC11())
	case05 := GetCompileCommand(SubmissionGo())
	case06 := GetCompileCommand(SubmissionUnknown())

	if case01 != "g++" {
		t.Error("Test Failed: (Case 01) GetCompileCommand return wrong compile command (GNU C++11) ")
		t.Error(case01)
		status = false
	}

	if case02 != "g++" {
		t.Error("Test Failed: (Case 02) GetCompileCommand return wrong compile command (GNU C++14) ")
		t.Error(case02)
		status = false
	}

	if case03 != "g++" {
		t.Error("Test Failed: (Case 03) GetCompileCommand return wrong compile command (GNU C++17) ")
		t.Error(case03)
		status = false
	}

	if case04 != "gcc" {
		t.Error("Test Failed: (Case 04) GetCompileCommand return wrong compile command (GNU C11) ")
		t.Error(case04)
		status = false
	}

	if case05 != "go" {
		t.Error("Test Failed: (Case 05) GetCompileCommand return wrong compile command (Go) ")
		t.Error(case05)
		status = false
	}

	if case06 != "" {
		t.Error("Test Failed: (Case 06) GetCompileCommand should return a empty string (Unknown) ")
		t.Error(case06)
		status = false
	}

	if status == true {
		t.Log("Package judge - GetCompileCommand test passed")
	} else {
		t.Log("Package judge - GetCompileCommand test failed")
	}
}

func Test_GetUnixCompileArgs(t *testing.T) {
	status := true

	case01 := GetUnixCompileArgs(SubmissionCPP11())
	case02 := GetUnixCompileArgs(SubmissionCPP14())
	case03 := GetUnixCompileArgs(SubmissionCPP17())
	case04 := GetUnixCompileArgs(SubmissionC11())
	case05 := GetUnixCompileArgs(SubmissionGo())
	case06 := GetUnixCompileArgs(SubmissionUnknown())

	if len(case01) != 6 || case01[0] != "./src/56192359/main.cpp" || case01[2] != "./src/56192359/main" {
		t.Error("Test Failed: (Case 01) GetUnixCompileArgs return wrong compile args (GNU C++11) ")
		t.Error(case01)
		status = false
	}

	if len(case02) != 6 || case02[0] != "./src/56192351/main.cpp" || case02[2] != "./src/56192351/main" {
		t.Error("Test Failed: (Case 02) GetUnixCompileArgs return wrong compile args (GNU C++14) ")
		t.Error(case02)
		status = false
	}

	if len(case03) != 6 || case03[0] != "./src/56192374/main.cpp" || case03[2] != "./src/56192374/main" {
		t.Error("Test Failed: (Case 03) GetUnixCompileArgs return wrong compile args (GNU C++17) ")
		t.Error(case03)
		status = false
	}

	if len(case04) != 6 || case04[0] != "./src/56192830/main.c" || case04[2] != "./src/56192830/main" {
		t.Error("Test Failed: (Case 04) GetUnixCompileArgs return wrong compile args (GNU C11) ")
		t.Error(case04)
		status = false
	}

	if len(case05) != 3 || case05[1] != "./src/40507899/main" || case05[2] != "./src/40507899/main.go" {
		t.Error("Test Failed: (Case 05) GetUnixCompileArgs return wrong compile args (Go) ")
		t.Error(case05)
		status = false
	}

	if len(case06) != 0 {
		t.Error("Test Failed: (Case 06) GetCompileCommand should return a empty string array (Unknown) ")
		t.Error(case06)
		status = false
	}

	if status == true {
		t.Log("Package judge - GetUnixCompileArgs test passed")
	} else {
		t.Log("Package judge - GetUnixCompileArgs test failed")
	}
}

func Test_GetWindowsCompileArgs(t *testing.T) {
	status := true

	case01 := GetWindowsCompileArgs(SubmissionCPP11())
	case02 := GetWindowsCompileArgs(SubmissionCPP14())
	case03 := GetWindowsCompileArgs(SubmissionCPP17())
	case04 := GetWindowsCompileArgs(SubmissionC11())
	case05 := GetWindowsCompileArgs(SubmissionGo())
	case06 := GetWindowsCompileArgs(SubmissionUnknown())

	if len(case01) != 7 || case01[0] != "./src/56192359/main.cpp" || case01[2] != "./src/56192359/main.exe" {
		t.Error("Test Failed: (Case 01) GetWindowsCompileArgs return wrong compile args (GNU C++11) ")
		t.Error(case01)
		status = false
	}

	if len(case02) != 7 || case02[0] != "./src/56192351/main.cpp" || case02[2] != "./src/56192351/main.exe" {
		t.Error("Test Failed: (Case 02) GetWindowsCompileArgs return wrong compile args (GNU C++14) ")
		t.Error(case02)
		status = false
	}

	if len(case03) != 7 || case03[0] != "./src/56192374/main.cpp" || case03[2] != "./src/56192374/main.exe" {
		t.Error("Test Failed: (Case 03) GetWindowsCompileArgs return wrong compile args (GNU C++17) ")
		t.Error(case03)
		status = false
	}

	if len(case04) != 7 || case04[0] != "./src/56192830/main.c" || case04[2] != "./src/56192830/main.exe" {
		t.Error("Test Failed: (Case 04) GetWindowsCompileArgs return wrong compile args (GNU C11) ")
		t.Error(case04)
		status = false
	}

	if len(case05) != 3 || case05[1] != "./src/40507899/main.exe" || case05[2] != "./src/40507899/main.go" {
		t.Error("Test Failed: (Case 05) GetWindowsCompileArgs return wrong compile args (Go) ")
		t.Error(case05)
		status = false
	}

	if len(case06) != 0 {
		t.Error("Test Failed: (Case 06) GetWindowsCompileArgs should return a empty string array (Unknown) ")
		t.Error(case06)
		status = false
	}

	if status == true {
		t.Log("Package judge - GetWindowsCompileArgs test passed")
	} else {
		t.Log("Package judge - GetWindowsCompileArgs test failed")
	}
}

func Test_GetUnixRunCommand(t *testing.T) {
	status := true

	case01 := GetUnixRunCommand(SubmissionCPP11())
	case02 := GetUnixRunCommand(SubmissionCPP14())
	case03 := GetUnixRunCommand(SubmissionCPP17())
	case04 := GetUnixRunCommand(SubmissionC11())
	case05 := GetUnixRunCommand(SubmissionPython2())
	case06 := GetUnixRunCommand(SubmissionPython3())
	case07 := GetUnixRunCommand(SubmissionGo())
	case08 := GetUnixRunCommand(SubmissionUnknown())

	if case01 != "./src/56192359/main" {
		t.Error("Test Failed: (Case 01) GetUnixRunCommand return wrong run command (GNU C++11) ")
		t.Error(case01)
		status = false
	}

	if case02 != "./src/56192351/main" {
		t.Error("Test Failed: (Case 02) GetUnixRunCommand return wrong run command (GNU C++14) ")
		t.Error(case02)
		status = false
	}

	if case03 != "./src/56192374/main" {
		t.Error("Test Failed: (Case 03) GetUnixRunCommand return wrong run command (GNU C++17) ")
		t.Error(case03)
		status = false
	}

	if case04 != "./src/56192830/main" {
		t.Error("Test Failed: (Case 04) GetUnixRunCommand return wrong run command (GNU C11) ")
		t.Error(case04)
		status = false
	}

	if case05 != "python2" {
		t.Error("Test Failed: (Case 05) GetUnixRunCommand return wrong run command (Python2) ")
		t.Error(case05)
		status = false
	}

	if case06 != "python3" {
		t.Error("Test Failed: (Case 06) GetUnixRunCommand return wrong run command (Python3) ")
		t.Error(case06)
		status = false
	}

	if case07 != "./src/40507899/main" {
		t.Error("Test Failed: (Case 07) GetUnixRunCommand return wrong run command (Go) ")
		t.Error(case07)
		status = false
	}

	if case08 != "" {
		t.Error("Test Failed: (Case 08) GetUnixRunCommand should return a empty string (Unknown) ")
		t.Error(case08)
		status = false
	}

	if status == true {
		t.Log("Package judge - GetUnixRunCommand test passed")
	} else {
		t.Log("Package judge - GetUnixRunCommand test failed")
	}
}

func Test_GetWindowsRunCommand(t *testing.T) {
	status := true

	case01 := GetWindowsRunCommand(SubmissionCPP11())
	case02 := GetWindowsRunCommand(SubmissionCPP14())
	case03 := GetWindowsRunCommand(SubmissionCPP17())
	case04 := GetWindowsRunCommand(SubmissionC11())
	case05 := GetWindowsRunCommand(SubmissionPython2())
	case06 := GetWindowsRunCommand(SubmissionPython3())
	case07 := GetWindowsRunCommand(SubmissionGo())
	case08 := GetWindowsRunCommand(SubmissionUnknown())

	if len(case01) == 0 || case01 != "./src/56192359/main.exe" {
		t.Error("Test Failed: (Case 01) GetWindowsRunCommand return wrong run command (GNU C++11) ")
		t.Error(case01)
		status = false
	}

	if len(case02) == 0 || case02 != "./src/56192351/main.exe" {
		t.Error("Test Failed: (Case 02) GetWindowsRunCommand return wrong run command (GNU C++14) ")
		t.Error(case02)
		status = false
	}

	if len(case03) == 0 || case03 != "./src/56192374/main.exe" {
		t.Error("Test Failed: (Case 03) GetWindowsRunCommand return wrong run command (GNU C++17) ")
		t.Error(case03)
		status = false
	}

	if len(case04) == 0 || case04 != "./src/56192830/main.exe" {
		t.Error("Test Failed: (Case 04) GetWindowsRunCommand return wrong run command (GNU C11) ")
		t.Error(case04)
		status = false
	}

	if len(case05) == 0 || case05 != "python2" {
		t.Error("Test Failed: (Case 05) GetWindowsRunCommand return wrong run command (Python2) ")
		t.Error(case05)
		status = false
	}

	if len(case06) == 0 || case06 != "python3" {
		t.Error("Test Failed: (Case 06) GetWindowsRunCommand return wrong run command (Python3) ")
		t.Error(case06)
		status = false
	}

	if len(case07) == 0 || case07 != "./src/40507899/main.exe" {
		t.Error("Test Failed: (Case 07) GetWindowsRunCommand return wrong run command (Go) ")
		t.Error(case07)
		status = false
	}

	if case08 != "" {
		t.Error("Test Failed: (Case 08) GetWindowsRunCommand should return a empty string (Unknown) ")
		t.Error(case08)
		status = false
	}

	if status == true {
		t.Log("Package judge - GetWindowsRunCommand test passed")
	} else {
		t.Log("Package judge - GetWindowsRunCommand test failed")
	}
}

func Test_GetRunArgs(t *testing.T) {
	status := true

	case01 := GetRunArgs(SubmissionCPP11())
	case02 := GetRunArgs(SubmissionCPP14())
	case03 := GetRunArgs(SubmissionCPP17())
	case04 := GetRunArgs(SubmissionC11())
	case05 := GetRunArgs(SubmissionPython2())
	case06 := GetRunArgs(SubmissionPython3())
	case07 := GetRunArgs(SubmissionGo())
	case08 := GetRunArgs(SubmissionUnknown())

	if len(case01) != 0 {
		t.Error("Test Failed: (Case 01) GetRunArgs return wrong run args (GNU C++11) ")
		t.Error(case01)
		status = false
	}

	if len(case02) != 0 {
		t.Error("Test Failed: (Case 02) GetRunArgs return wrong run args (GNU C++14) ")
		t.Error(case02)
		status = false
	}

	if len(case03) != 0 {
		t.Error("Test Failed: (Case 03) GetRunArgs return wrong run args (GNU C++17) ")
		t.Error(case03)
		status = false
	}

	if len(case04) != 0 {
		t.Error("Test Failed: (Case 04) GetRunArgs return wrong run args (GNU C11) ")
		t.Error(case04)
		status = false
	}

	if len(case05) == 0 || case05[0] != "./src/55525058/main.py" {
		t.Error("Test Failed: (Case 05) GetRunArgs return wrong run args (Python2) ")
		t.Error(case05)
		status = false
	}

	if len(case06) == 0 || case06[0] != "./src/56192367/main.py" {
		t.Error("Test Failed: (Case 06) GetRunArgs return wrong run args (Python3) ")
		t.Error(case06)
		status = false
	}

	if len(case07) != 0 {
		t.Error("Test Failed: (Case 07) GetRunArgs return wrong run args (Go) ")
		t.Error(case07)
		status = false
	}

	if len(case08) != 0 {
		t.Error("Test Failed: (Case 08) GetRunArgs should return a empty string array (Unknown) ")
		t.Error(case08)
		status = false
	}

	if status == true {
		t.Log("Package judge - GetRunArgs test passed")
	} else {
		t.Log("Package judge - GetRunArgs test failed")
	}
}
