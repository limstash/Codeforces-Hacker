package submission

import (
	"testing"

	. "github.com/hytzongxuan/Codeforces-Hacker/common"
	"github.com/hytzongxuan/Codeforces-Hacker/module/token"
)

func Test_GetSubmission(t *testing.T) {
	status := true

	auth := Authentication{}
	err := token.GetCSRF(&auth, "https://codeforces.com")

	contest := Contest{}
	problem := Problem{}

	contest.ID = 1175
	contest.StartTimeSeconds = 1559745300
	contest.DurationSeconds = 7200

	problem.Index = "A"

	res, err := GetSubmission(contest, problem, &auth, "https://codeforces.com")

	if err != nil {
		t.Error(err)
		status = false
	}

	if len(res) <= 0 {
		t.Error("Test Failed: GetSubmission should return a non-empty submissions array")
		status = false
	}

	if status == true {
		t.Log("Package submission - GetSubmission test passed")
	} else {
		t.Log("Package submission - GetSubmission test failed")
	}
}
