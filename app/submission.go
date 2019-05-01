package app

import (
	"net/http"

	"github.com/hytzongxuan/Codeforces-Hacker/module/contest"
	"github.com/hytzongxuan/Codeforces-Hacker/module/submission"
)

func GetSubmission(cookie *[]*http.Cookie, contestInfo contest.Contest, problem string) []submission.Submission {
	data, _ := submission.GetSubmissionArray(contestInfo.ID, cookie, contestInfo.StartTimeSeconds, contestInfo.DurationSeconds, problem)
	return data
}
