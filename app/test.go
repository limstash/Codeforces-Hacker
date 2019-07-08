package app

import (
	"os"
	"strconv"

	"github.com/limstash/Codeforces-Hacker/module/judge"

	"github.com/limstash/Codeforces-Hacker/module/code"

	. "github.com/limstash/Codeforces-Hacker/common"
	"github.com/termie/go-shutil"
)

func checkLanguage(lang string, AvailableLanguage Language) bool {
	switch lang {
	case "GNU C++11":
		if AvailableLanguage.GNUCPP11 != true {
			return false
		}
	case "GNU C++14":
		if AvailableLanguage.GNUCPP14 != true {
			return false
		}
	case "GNU C++17":
		if AvailableLanguage.GNUCPP17 != true {
			return false
		}
	case "GNU C11":
		if AvailableLanguage.GNUC11 != true {
			return false
		}
	case "Go":
		if AvailableLanguage.Go != true {
			return false
		}
	case "Python 2":
		if AvailableLanguage.Python2 != true {
			return false
		}
	case "Python 3":
		if AvailableLanguage.Python3 != true {
			return false
		}
	default:
		return false
	}

	return true
}

func testResult(submissionID int, result string, err bool) {
	if err {
		log(2, "Submission "+strconv.Itoa(submissionID)+": "+result)
	} else {
		log(3, "Submission "+strconv.Itoa(submissionID)+": "+result)
	}
}

func needCompile(lang string) bool {
	if lang == "GNU C++11" || lang == "GNU C++14" || lang == "GNU C++17" || lang == "Go" || lang == "GNU C11" {
		return true
	}

	return false
}

func test(submission Submission, problem Problem, config Config, auth *Authentication, support Language, server string) {
	res := checkLanguage(submission.Language, support)

	if res == false {
		testResult(submission.SubmissionID, "Not Support "+submission.Language, true)
		return
	}

	err := os.MkdirAll(submission.Path, 0777)

	if err != nil {
		testResult(submission.SubmissionID, "System Error", true)
		testResult(submission.SubmissionID, err.Error(), true)
		return
	}

	err = shutil.CopyFile(config.Testcase.InputFile, submission.Path+"/data.in", false)

	if err != nil {
		testResult(submission.SubmissionID, "System Error", true)
		testResult(submission.SubmissionID, err.Error(), true)
		return
	}

	err = shutil.CopyFile(config.Testcase.OutputFile, submission.Path+"/data.ans", false)

	if err != nil {
		testResult(submission.SubmissionID, "System Error", true)
		testResult(submission.SubmissionID, err.Error(), true)
		return
	}

	source, err := code.QueryCode(submission.SubmissionID, auth, server)

	if err != nil {
		testResult(submission.SubmissionID, "System Error", true)
		testResult(submission.SubmissionID, err.Error(), true)
		return
	}

	submission.Code = source

	err = code.SaveCode(submission)

	if err != nil {
		testResult(submission.SubmissionID, "System Error", true)
		testResult(submission.SubmissionID, err.Error(), true)
		return
	}

	if needCompile(submission.Language) {
		res, err := judge.Compile(submission)

		if err != nil && err.Error() == "Compile Timeout" {
			testResult(submission.SubmissionID, "Compile Timeout", true)
			return
		}

		if res == false {
			testResult(submission.SubmissionID, "Compile Error", true)
			return
		}
	}

	res, time, memory := judge.Runcode(submission, problem)

	if time > problem.Timelimit {
		testResult(submission.SubmissionID, "Time Limit Exceeded", true)
		return
	}

	if memory > problem.Memorylimit*1024 {
		testResult(submission.SubmissionID, "Memory Limit Exceeded", true)
		return
	}

	if res == false {
		testResult(submission.SubmissionID, "Runtime Error", true)
		return
	}

	res = judge.Diff(submission)

	if res == false {
		testResult(submission.SubmissionID, "Wrong Answer", true)
		return
	}

	testResult(submission.SubmissionID, "Accepted", false)
	return
}
