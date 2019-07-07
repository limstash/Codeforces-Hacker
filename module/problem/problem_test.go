package problem

import (
	"testing"

	. "github.com/limstash/Codeforces-Hacker/common"
)

func Test_GetContest(t *testing.T) {
	status := true
	auth := Authentication{}

	contest, err := GetContest(&auth, "https://codeforces.com")

	if err != nil {
		t.Error(err)
		status = false
	}

	if len(contest.Result) <= 0 {
		t.Error("Test Failed: GetContest should return a non-empty contests array")
		status = false
	}

	if status == true {
		t.Log("Package contest - GetContest test passed")
	} else {
		t.Log("Package contest - GetContest test failed")
	}
}

func Test_GetProblem(t *testing.T) {
	status := true
	auth := Authentication{}

	problem, err := GetProblem(&auth, "https://codeforces.com")

	if err != nil {
		t.Error(err)
		status = false
	}

	if len(problem.Result.Problems) <= 0 {
		t.Error("Test Failed: GetProblem should return a non-empty problems array")
		status = false
	}

	if status == true {
		t.Log("Package contest - GetProblem test passed")
	} else {
		t.Log("Package contest - GetProblem test failed")
	}
}
