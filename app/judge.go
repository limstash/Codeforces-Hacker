package app

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/hytzongxuan/Codeforces-Hacker/module/code"
	"github.com/hytzongxuan/Codeforces-Hacker/module/judge"
)

func runCode(SubmissionID int, Language string, customDiff bool) (bool, error) {
	_, e := judge.Judge(SubmissionID, Language, customDiff)

	if e == nil {
		return true, nil
	}

	if e.Error() == "Not Support" {
		return true, nil
	}

	if e.Error() == "Compile Error" {
		fmt.Println("[Info] Code " + strconv.Itoa(SubmissionID) + " Compile Failed")
	}

	if e.Error() == "Runtime Error" {
		fmt.Println("[Info] Code " + strconv.Itoa(SubmissionID) + " Runtime Error")
	}

	if e.Error() == "Wrong Answer" {
		fmt.Println("[Info] Code " + strconv.Itoa(SubmissionID) + " Wrong Answer")
	}

	return false, e
}

func copyfile(origin string, remote string) {
	from, err := os.Open(origin)
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile(remote, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}
}

func saveCode(SubmissionID int, Language string, Cookie *[]*http.Cookie, CSRF string) error {
	text, e := code.QueryCode(SubmissionID, Cookie, CSRF)

	if e != nil {
		return e
	}

	os.MkdirAll(getPath()+"/src/"+strconv.Itoa(SubmissionID), 0777)
	copyfile(getPath()+"/src/data.in", getPath()+"/src/"+strconv.Itoa(SubmissionID)+"/data.in")
	copyfile(getPath()+"/src/data.ans", getPath()+"/src/"+strconv.Itoa(SubmissionID)+"/data.ans")

	e = code.SaveCode(SubmissionID, Language, text)

	if e != nil {
		return e
	}

	return nil
}

func TestCode(SubmissionID int, Language string, customDiff bool, Cookie *[]*http.Cookie, CSRF string) (bool, error) {
	e := saveCode(SubmissionID, Language, Cookie, CSRF)

	if e != nil {
		return false, e
	}

	return runCode(SubmissionID, Language, customDiff)
}
