package code

import (
	"bufio"
	"errors"
	"os"
	"strconv"

	"github.com/bitly/go-simplejson"
	"github.com/limstash/Codeforces-Hacker/module/conn"

	. "github.com/limstash/Codeforces-Hacker/common"
)

func apiQueryCode(SubmissionID int, auth *Authentication, server string) (Response, error) {
	request := Request{}

	request.URL = server + "/data/submitSource"
	request.Method = "POST"
	request.NotRedirect = false
	request.Authentication = auth
	request.Header = map[string]string{"X-Csrf-Token": auth.CSRF, "X-Requested-With": "XMLHttpRequest", "Origin": "https://codeforces.com", "Referer": "https://codeforces.com/problemset/status", "Host": "codeforces.com"}
	request.Data = map[string]string{"submissionId": strconv.Itoa(SubmissionID), "csrf_token": auth.CSRF}

	response, err := conn.HTTPRequest(request)

	return response, err
}

func QueryCode(SubmissionID int, auth *Authentication, server string) (string, error) {
	response, err := apiQueryCode(SubmissionID, auth, server)

	if err != nil {
		return "", err
	}

	js, e := simplejson.NewJson(response.ResponseBody)

	if e != nil {
		return "", e
	}

	source, e := js.Get("source").String()

	if e != nil {
		return "", e
	}

	return source, nil
}

func SaveCode(submission Submission) error {
	suffix := map[string]string{
		"GNU C11":   "c",
		"GNU C++11": "cpp",
		"GNU C++14": "cpp",
		"GNU C++17": "cpp",
		"Go":        "go",
		"Python 2":  "py",
		"Python 3":  "py",
	}

	SubmissionPath := submission.Path

	err := os.MkdirAll(SubmissionPath, 0777)

	if err != nil {
		return err
	}

	extensionSuffix := suffix[submission.Language]

	if len(extensionSuffix) == 0 {
		return errors.New("No such programming language")
	}

	outputFile, err := os.OpenFile(SubmissionPath+"/main."+extensionSuffix, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return err
	}

	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)

	outputWriter.WriteString(submission.Code)
	outputWriter.Flush()

	return nil
}
