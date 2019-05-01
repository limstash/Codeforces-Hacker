package judge

import (
	"os/exec"
	"strconv"
)

func CompileUnixC11(SubmissionID int, lang string) bool {
	cmd := exec.Command("gcc", getPath()+"/src/"+strconv.Itoa(SubmissionID)+"/main.c", "-o", getPath()+"/src/"+strconv.Itoa(SubmissionID)+"/main", "--std=c++11", "-lm", "-DONLINE_JUDGE")
	err := cmd.Run()

	if err != nil {
		return false
	}

	return true
}

func CompileUnixCPP11(SubmissionID int, lang string) bool {
	cmd := exec.Command("g++", getPath()+"/src/"+strconv.Itoa(SubmissionID)+"/main.cpp", "-o", getPath()+"/src/"+strconv.Itoa(SubmissionID)+"/main", "--std=c++11", "-lm", "-DONLINE_JUDGE")
	err := cmd.Run()

	if err != nil {
		return false
	}

	return true
}

func CompileUnixCPP14(SubmissionID int, lang string) bool {
	cmd := exec.Command("g++", getPath()+"/src/"+strconv.Itoa(SubmissionID)+"/main.cpp", "-o", getPath()+"/src/"+strconv.Itoa(SubmissionID)+"/main", "--std=c++14", "-lm", "-DONLINE_JUDGE")
	err := cmd.Run()

	if err != nil {
		return false
	}

	return true
}

func CompileUnixCPP17(SubmissionID int, lang string) bool {
	cmd := exec.Command("g++", getPath()+"/src/"+strconv.Itoa(SubmissionID)+"/main.cpp", "-o", getPath()+"/src/"+strconv.Itoa(SubmissionID)+"/main", "--std=c++17", "-lm", "-DONLINE_JUDGE")
	err := cmd.Run()

	if err != nil {
		return false
	}

	return true
}

func CompileUnixGo(SubmissionID int, lang string) bool {
	cmd := exec.Command("go", "build", "-o", getPath()+"/src/"+strconv.Itoa(SubmissionID)+"/main", getPath()+"/src/"+strconv.Itoa(SubmissionID)+"/main.go")
	err := cmd.Run()

	if err != nil {
		return false
	}

	return true
}

func CompileUnixCode(SubmissionID int, lang string) bool {
	result := true

	switch lang {
	case "GNU C11":
		result = CompileUnixC11(SubmissionID, lang)
	case "GNU C++11":
		result = CompileUnixCPP11(SubmissionID, lang)
	case "GNU C++14":
		result = CompileUnixCPP14(SubmissionID, lang)
	case "GNU C++17":
		result = CompileUnixCPP17(SubmissionID, lang)
	case "Go":
		result = CompileUnixGo(SubmissionID, lang)
	default:
		result = true
	}

	return result
}
