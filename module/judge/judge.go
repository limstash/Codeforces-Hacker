package judge

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func getPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	rst := filepath.Dir(path)
	return rst
}

func testOnLinux(SubmissionID int, Language string, customDiff bool) (bool, error) {
	status := CompileUnixCode(SubmissionID, Language)

	if status == false {
		return false, errors.New("Compile Error")
	}

	status = RunUnixCode(SubmissionID, Language)

	if status == false {
		return false, errors.New("Runtime Error")
	}

	if customDiff == false {
		status = Diff(SubmissionID)

		if status == false {
			return false, errors.New("Wrong Answer")
		}

	} else {
		//todo
		panic("todo")
	}

	return true, nil
}

func judgeUnix(SubmissionID int, Language string, customDiff bool) (bool, error) {
	res, e := testOnLinux(SubmissionID, Language, customDiff)
	return res, e
}

func Judge(SubmissionID int, Language string, customDiff bool) (bool, error) {
	AvailableLanguage := GetAvailableLanguage()

	switch Language {
	case "GNU C++11":
		if AvailableLanguage.GNUCPP11 != true {
			return false, errors.New("Not Support")
		}
	case "GNU C++14":
		if AvailableLanguage.GNUCPP14 != true {
			return false, errors.New("Not Support")
		}
	case "GNU C++17":
		if AvailableLanguage.GNUCPP17 != true {
			return false, errors.New("Not Support")
		}
	case "GNU C11":
		if AvailableLanguage.GNUC11 != true {
			return false, errors.New("Not Support")
		}
	case "Go":
		if AvailableLanguage.Go != true {
			return false, errors.New("Not Support")
		}
	case "Python2":
		if AvailableLanguage.Python2 != true {
			return false, errors.New("Not Support")
		}
	case "Python3":
		if AvailableLanguage.Python3 != true {
			return false, errors.New("Not Support")
		}
	default:
		return false, errors.New("Not Support")
	}

	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		res, e := judgeUnix(SubmissionID, Language, customDiff)
		return res, e
	} else {
		//todo
		panic("todo")
	}

	return false, nil
}
