package code

import (
	"bufio"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/bitly/go-simplejson"
	"github.com/hytzongxuan/Codeforces-Hacker/module/con"
)

func QueryCode(SubmissionID int, cookie *[]*http.Cookie, CSRF string) (string, error) {
	res, err := con.HTTPPostByte("https://codeforces.com/data/submitSource", cookie, map[string]string{"X-Csrf-Token": CSRF, "X-Requested-With": "XMLHttpRequest", "Origin": "https://codeforces.com", "Referer": "https://codeforces.com/problemset/status", "Host": "codeforces.com"}, map[string]string{"submissionId": strconv.Itoa(SubmissionID), "csrf_token": CSRF})

	if err != nil {
		return "", err
	}

	js, e := simplejson.NewJson(res)

	if e != nil {
		return "", e
	}

	source, e := js.Get("source").String()

	if e != nil {
		return "", e
	}

	return source, nil
}

func getPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	rst := filepath.Dir(path)
	return rst
}

func saveCodeOnUnix(SubmissionID int, Language string, code string) error {
	suffix := map[string]string{
		"GNU C11":   "c",
		"GNU C++11": "cpp",
		"GNU C++14": "cpp",
		"GNU C++17": "cpp",
		"Go":        "go",
		"Python 2":  "py",
		"Python 3":  "py",
	}

	SubmissionPath := getPath() + "/src/" + strconv.Itoa(SubmissionID)

	outputFile, err := os.OpenFile(SubmissionPath+"/main."+suffix[Language], os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return err
	}

	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)

	outputWriter.WriteString(code)
	outputWriter.Flush()

	return nil
}

func saveCodeOnWindows(SubmissionID int, Language string, code string) error {
	suffix := map[string]string{
		"GNU C11":   "c",
		"GNU C++11": "cpp",
		"GNU C++14": "cpp",
		"GNU C++17": "cpp",
		"Go":        "go",
		"Python 2":  "py",
		"Python 3":  "py",
	}

	SubmissionPath := getPath() + "\\src\\" + strconv.Itoa(SubmissionID)

	outputFile, err := os.OpenFile(SubmissionPath+"\\main."+suffix[Language], os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return err
	}

	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)

	outputWriter.WriteString(code)
	outputWriter.Flush()

	return nil
}

func SaveCode(SubmissionID int, Language string, code string) error {
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		e := saveCodeOnUnix(SubmissionID, Language, code)
		return e
	}

	e := saveCodeOnWindows(SubmissionID, Language, code)
	return e
}
