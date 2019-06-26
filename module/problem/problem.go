package problem

import (
	"encoding/json"
	"errors"

	"github.com/hytzongxuan/Codeforces-Hacker/module/conn"

	. "github.com/hytzongxuan/Codeforces-Hacker/common"
)

func apiQueryContest(auth *Authentication, server string) (Response, error) {
	request := Request{}

	request.URL = server + "/api/contest.list?gym=false"
	request.Method = "GET"
	request.NotRedirect = false
	request.Authentication = auth
	request.Header = map[string]string{"Host": "codeforces.com"}

	response, err := conn.HTTPRequest(request)

	return response, err
}

func apiQueryProblem(auth *Authentication, server string) (Response, error) {
	request := Request{}

	request.URL = server + "/api/problemset.problems"
	request.Method = "GET"
	request.NotRedirect = false
	request.Authentication = auth
	request.Header = map[string]string{"Host": "codeforces.com"}

	response, err := conn.HTTPRequest(request)

	return response, err
}

func GetContest(auth *Authentication, server string) (Contests, error) {
	contests := Contests{}

	response, err := apiQueryContest(auth, server)

	if err != nil {
		return contests, err
	}

	err = json.Unmarshal(response.ResponseBody, &contests)

	if err != nil {
		return contests, err
	}

	if contests.Status != "OK" {
		return contests, errors.New("Codeforces API return unknown error")
	}

	return contests, nil
}

func GetProblem(auth *Authentication, server string) (Problems, error) {
	problems := Problems{}

	response, err := apiQueryProblem(auth, server)

	if err != nil {
		return problems, err
	}

	err = json.Unmarshal(response.ResponseBody, &problems)

	if err != nil {
		return problems, err
	}

	if problems.Status != "OK" {
		return problems, errors.New("Codeforces API return unknown error")
	}

	return problems, nil
}
