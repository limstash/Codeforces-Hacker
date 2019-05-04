package judge

import (
	"os/exec"
	"strconv"
)

func CompileWindowsC11(SubmissionID int, lang string) bool {
	cmd := exec.Command("gcc", getPath()+"\\src\\"+strconv.Itoa(SubmissionID)+"\\main.c", "-o", getPath()+"\\src\\"+strconv.Itoa(SubmissionID)+"\\main", "--std=c++11", "-lm", "-Wl,--stack=536870912", "-DONLINE_JUDGE")
	err := cmd.Run()

	if err != nil {
		return false
	}

	return true
}

func CompileWindowsCPP11(SubmissionID int, lang string) bool {
	cmd := exec.Command("g++", getPath()+"\\src\\"+strconv.Itoa(SubmissionID)+"\\main.cpp", "-o", getPath()+"\\src\\"+strconv.Itoa(SubmissionID)+"\\main", "--std=c++11", "-lm", "-Wl,--stack=536870912", "-DONLINE_JUDGE")
	err := cmd.Run()

	if err != nil {
		return false
	}

	return true
}

func CompileWindowsCPP14(SubmissionID int, lang string) bool {
	cmd := exec.Command("g++", getPath()+"\\src\\"+strconv.Itoa(SubmissionID)+"\\main.cpp", "-o", getPath()+"\\src\\"+strconv.Itoa(SubmissionID)+"\\main", "--std=c++14", "-lm", "-Wl,--stack=536870912", "-DONLINE_JUDGE")
	err := cmd.Run()

	if err != nil {
		return false
	}

	return true
}

func CompileWindowsCPP17(SubmissionID int, lang string) bool {
	cmd := exec.Command("g++", getPath()+"\\src\\"+strconv.Itoa(SubmissionID)+"\\main.cpp", "-o", getPath()+"\\src\\"+strconv.Itoa(SubmissionID)+"\\main", "--std=c++17", "-lm", "-Wl,--stack=536870912", "-DONLINE_JUDGE")
	err := cmd.Run()

	if err != nil {
		return false
	}

	return true
}

func CompileWindowsGo(SubmissionID int, lang string) bool {
	cmd := exec.Command("go", "build", "-o", getPath()+"\\src\\"+strconv.Itoa(SubmissionID)+"\\main", getPath()+"\\src\\"+strconv.Itoa(SubmissionID)+"\\main.go")
	err := cmd.Run()

	if err != nil {
		return false
	}

	return true
}

func CompileWindowsCode(SubmissionID int, lang string) bool {
	result := true

	switch lang {
	case "GNU C11":
		result = CompileWindowsC11(SubmissionID, lang)
	case "GNU C++11":
		result = CompileWindowsCPP11(SubmissionID, lang)
	case "GNU C++14":
		result = CompileWindowsCPP14(SubmissionID, lang)
	case "GNU C++17":
		result = CompileWindowsCPP17(SubmissionID, lang)
	case "Go":
		result = CompileWindowsGo(SubmissionID, lang)
	default:
		result = true
	}

	return result
}
