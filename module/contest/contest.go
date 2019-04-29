package contest

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/bitly/go-simplejson"
	"github.com/hytzongxuan/Codeforces-Hacker/module/conn"
)

//Contest type contains contest info
type Contest struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Types               string `json:"type"`
	Phase               string `json:"phase"`
	Frozen              bool   `json:"frozen"`
	DurationSeconds     int64  `json:"durationSeconds"`
	StartTimeSeconds    int64  `json:"startTimeSeconds"`
	RelativeTimeSeconds int64  `json:"relativeTimeSeconds"`
}

//Contests type contain contests info
type Contests struct {
	Status string    `json:"status"`
	Result []Contest `json:"result"`
}

//Problem type contain problem info
type Problem struct {
	Index string
	Name  string
}

func queryContests(cookie *[]*http.Cookie) (string, error) {
	res, err := conn.HTTPGet("https://codeforces.com/api/contest.list?gym=false", cookie, map[string]string{"HOST": "codeforces.com"})

	if err != nil {
		return "", err
	}

	return res, nil
}

//GetContests will fetch contests list
func GetContests(cookie *[]*http.Cookie) ([]Contest, error) {
	data, e := queryContests(cookie)

	if e != nil {
		return nil, e
	}

	res := Contests{}
	e = json.Unmarshal([]byte(data), &res)

	if e != nil {
		return nil, e
	}

	if res.Status != "OK" {
		return nil, errors.New("Codeforces Return Error Response")
	}

	return res.Result, nil
}

func queryProblems(contestID int, cookie *[]*http.Cookie) ([]byte, error) {
	res, err := conn.HTTPGetByte("https://codeforces.com/api/contest.standings?contestId="+strconv.Itoa(contestID)+"&from=1&count=1", cookie, map[string]string{"HOST": "codeforces.com"})

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetProblems will fetch problems list from a contest
func GetProblems(contestID int, cookie *[]*http.Cookie) ([]Problem, error) {
	data, e := queryProblems(contestID, cookie)

	if e != nil {
		return nil, e
	}

	res := []Problem{}

	js, e := simplejson.NewJson(data)

	if e != nil {
		return nil, e
	}

	status, e := js.Get("status").String()

	if status != "OK" {
		return nil, errors.New("Codeforces Return Error Response")
	}

	problem, e := js.Get("result").Get("problems").Array()

	if e != nil {
		return nil, e
	}

	for i := 0; i < len(problem); i++ {
		index := problem[i].(map[string]interface{})["index"].(string)
		name := problem[i].(map[string]interface{})["name"].(string)

		res = append(res, Problem{index, name})
	}

	return res, nil
}
