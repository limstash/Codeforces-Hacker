package submission

import (
	"errors"
	"strconv"

	"github.com/bitly/go-simplejson"
	"github.com/limstash/Codeforces-Hacker/module/conn"

	. "github.com/limstash/Codeforces-Hacker/common"
)

func apiQuerySubmission(contestID int, auth *Authentication, server string) (Response, error) {
	request := Request{}

	request.URL = server + "/api/contest.status?contestId=" + strconv.Itoa(contestID) + "&from=1&count=100000000"
	request.Method = "GET"
	request.NotRedirect = false
	request.Authentication = auth
	request.Header = map[string]string{"Host": "codeforces.com"}

	response, err := conn.HTTPRequest(request)

	return response, err
}

func GetSubmission(contest Contest, problem Problem, auth *Authentication, server string) ([]Submission, error) {
	submission := []Submission{}

	response, err := apiQuerySubmission(contest.ID, auth, server)

	if err != nil {
		return submission, err
	}

	js, err := simplejson.NewJson(response.ResponseBody)

	if err != nil {
		return submission, err
	}

	status, err := js.Get("status").String()

	if err != nil {
		return submission, err
	}

	if status != "OK" {
		return submission, errors.New("Codeforces API return unknown error")
	}

	submissions, err := js.Get("result").Array()

	if err != nil {
		return submission, err
	}

	for i := 0; i < len(submissions); i++ {
		submissionInfo := js.Get("result").GetIndex(i)

		verdict := submissionInfo.Get("verdict").MustString()
		submissionID := submissionInfo.Get("id").MustInt()
		language := submissionInfo.Get("programmingLanguage").MustString()
		problemIndex := submissionInfo.Get("problem").Get("index").MustString()
		submitTime := submissionInfo.Get("creationTimeSeconds").MustInt64()

		if verdict == "OK" && problemIndex == problem.Index && submitTime > contest.StartTimeSeconds && submitTime <= contest.StartTimeSeconds+contest.DurationSeconds {
			submission = append(submission, Submission{SubmissionID: submissionID, Language: language, Code: "", Path: "./src/" + strconv.Itoa(submissionID)})
		}
	}

	return submission, nil
}
