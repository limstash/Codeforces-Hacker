package submission

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/bitly/go-simplejson"
	"github.com/hytzongxuan/Codeforces-Hacker/module/conn"
)

func querySubmission(contestID int, cookie *[]*http.Cookie) ([]byte, error) {
	res, err := conn.HTTPGetByte("http://codeforces.com/api/contest.status?contestId="+strconv.Itoa(contestID)+"&from=1&count=100000000", cookie, map[string]string{"HOST": "codeforces.com"})

	if err != nil {
		return nil, err
	}

	return res, nil
}

type Submission struct {
	SubmissionID int
	Language     string
}

func GetSubmissionArray(contestID int, cookie *[]*http.Cookie, startTime int64, duringTime int64, problem string) ([]Submission, error) {
	resp, e := querySubmission(contestID, cookie)

	if e != nil {
		return nil, e
	}

	js, e := simplejson.NewJson(resp)

	status, e := js.Get("status").String()

	if status != "OK" {
		return nil, errors.New("Codeforces Return Error Response")
	}

	submissions, e := js.Get("result").Array()

	data := []Submission{}

	for i := 0; i < len(submissions); i++ {
		submission := js.Get("result").GetIndex(i)

		verdict := submission.Get("verdict").MustString()
		submissionID := submission.Get("id").MustInt()
		language := submission.Get("programmingLanguage").MustString()
		problemIndex := submission.Get("problem").Get("index").MustString()
		submitTime := submission.Get("creationTimeSeconds").MustInt64()

		if verdict == "OK" && problemIndex == problem && submitTime > startTime && submitTime <= startTime+duringTime {
			data = append(data, Submission{submissionID, language})
		}
	}

	return data, nil
}
