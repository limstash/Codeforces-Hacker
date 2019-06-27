package app

import (
	"errors"
	"strings"
	"time"

	. "github.com/hytzongxuan/Codeforces-Hacker/common"
)

// FindContest will find the latest contest in the contests array
func FindContest(contests []Contest, id int) (Contest, error) {
	contest := Contest{}

	contestsSiz := len(contests)
	contestIndex := -1

	for i := 0; i < contestsSiz; i++ {
		if contests[i].ID == id {
			contestIndex = i
			break
		}
	}

	if contestIndex == -1 {
		return contest, errors.New("No such contest")
	}

	contest = contests[contestIndex]
	log(3, "Switch to "+contest.Name)

	currentTime := time.Now().Unix()

	if !strings.Contains(contest.Name, "Educational") && !strings.Contains(contest.Name, "Div. 3") {
		return contest, errors.New("This contest is not available to hack")
	}

	timeDelta := currentTime - contest.StartTimeSeconds

	if timeDelta >= 12*3600 {
		return contest, errors.New("Contest finished")
	}

	return contest, nil
}

func FindProblem(problems []Problem, contest Contest, index string) (Problem, error) {
	problem := Problem{}

	problemSiz := len(problems)
	problemIndex := -1

	for i := 0; i < problemSiz; i++ {
		if problems[i].ContestID == contest.ID && problems[i].Index == index {
			problemIndex = i
			break
		}
	}

	if problemIndex == -1 {
		return problem, errors.New("No such problem")
	}

	problem = problems[problemIndex]

	log(3, "Switch to Problem "+problem.Index+" - "+problem.Name)

	return problem, nil
}
