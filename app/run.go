package app

import (
	"errors"
	"os"
	"regexp"
	"strconv"

	. "github.com/limstash/Codeforces-Hacker/common"
	"github.com/limstash/Codeforces-Hacker/module/conn"
	"github.com/limstash/Codeforces-Hacker/module/judge"
	"github.com/limstash/Codeforces-Hacker/module/submission"
)

func apiQueryLimit(contestID string, problemIndex string, auth *Authentication, server string) (Response, error) {
	request := Request{}

	request.URL = server + "/problemset/problem/" + contestID + "/" + problemIndex
	request.Method = "GET"
	request.Authentication = auth
	request.Header = map[string]string{"Referer": "http://codeforces.com/problemset", "Host": "codeforces.com"}

	response, err := conn.HTTPRequest(request)
	return response, err
}

func getLimit(problem *Problem, auth *Authentication, server string) error {
	res, err := apiQueryLimit(strconv.Itoa(problem.ContestID), problem.Index, auth, server)

	if err != nil {
		return err
	}

	flysnowRegexp := regexp.MustCompile(`(\d+) second`)
	params := flysnowRegexp.FindStringSubmatch(string(res.ResponseBody))

	if len(params) == 0 {
		return errors.New("Failed to fetch time limit")
	}

	timelimit, err := strconv.Atoi(params[1])
	problem.Timelimit = timelimit * 1000

	flysnowRegexp = regexp.MustCompile(`(\d+) megabyte`)
	params = flysnowRegexp.FindStringSubmatch(string(res.ResponseBody))

	if len(params) == 0 {
		return errors.New("Failed to fetch memory limit")
	}

	memorylimit, err := strconv.Atoi(params[1])
	problem.Memorylimit = memorylimit * 1024

	return nil
}

func run(contest Contest, problem Problem, config Config, auth *Authentication, server string) {
	err := getLimit(&problem, auth, server)

	if err != nil {
		log(1, err.Error())
	}

	log(3, "Time limit: "+strconv.Itoa(problem.Timelimit)+" ms")
	log(3, "Memory limit: "+strconv.Itoa(problem.Memorylimit)+" KB")

	submissions, err := submission.GetSubmission(contest, problem, auth, server)

	if err != nil {
		log(1, err.Error())
	}

	log(3, "Number of submissions: "+strconv.Itoa(len(submissions)))

	language := judge.GetSupport()

	for i := 0; i < len(submissions); i++ {
		test(submissions[i], problem, config, auth, language, server)
		os.RemoveAll(submissions[i].Path)
	}
}
