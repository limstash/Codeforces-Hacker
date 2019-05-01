package app

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hytzongxuan/Codeforces-Hacker/module/contest"
)

// FindContest will find the latest contest in the contests array
func FindContest(contests []contest.Contest) (contest.Contest, error) {
	contestsSiz := len(contests)

	if contestsSiz <= 0 {
		return contest.Contest{}, errors.New("[Info] Contests is an empty field")
	}

	var lastContestTime int64 = 100000000
	lastContestIndex := -1

	currentTime := time.Now().Unix()

	for i := 0; i < contestsSiz; i++ {
		if strings.Contains(contests[i].Name, "Educational") || strings.Contains(contests[i].Name, "Div. 3") {
			timeDelta := currentTime - contests[i].StartTimeSeconds

			if timeDelta > 0 && timeDelta < lastContestTime {
				lastContestTime = timeDelta
				lastContestIndex = i
			}

		}
	}

	if lastContestIndex == -1 {
		return contest.Contest{}, errors.New("[Info] No vaild contest")
	}

	fmt.Println("[Info] The current contest is " + contests[lastContestIndex].Name)

	openHackingPhase := currentTime - contests[lastContestIndex].StartTimeSeconds - contests[lastContestIndex].DurationSeconds

	if openHackingPhase < 12*3600 {
		return contest.Contest{}, errors.New("[Info] Open hacking phase finished")
	}

	fmt.Println("[Info] Open hacking phase running")

	return contests[lastContestIndex], nil
}

// ChooseProblem will fetch the problem in the contest and read user's input from stdin
func ChooseProblem(ContestID int, cookie *[]*http.Cookie) (string, error) {
	fmt.Println("[Info] Fetching Problems...")

	problems, e := contest.GetProblems(ContestID, cookie)

	if e != nil {
		return "", e
	}

	fmt.Println("")
	fmt.Println("Please choose the problem you want to hack")
	fmt.Println("")

	for i := 0; i < len(problems); i++ {
		fmt.Println("[" + strconv.Itoa(i+1) + "] " + problems[i].Index + ". " + problems[i].Name)
	}

	var choose int

	fmt.Println("")
	fmt.Printf("Your Choose is [1-" + strconv.Itoa(len(problems)) + "] : ")

	myReader := bufio.NewReader(nil)
	myReader.Reset(os.Stdin)
	content, e := myReader.ReadString('\n')

	if e != nil {
		return "", e
	}

	fields := strings.Fields(content)

	if fields == nil || len(fields) == 0 {
		choose = 0
	} else {
		choose, e = strconv.Atoi(fields[0])
	}

	for e != nil || choose < 1 || choose > len(problems) {
		fmt.Printf("Your Choose is [1-" + strconv.Itoa(len(problems)) + "] : ")

		myReader = bufio.NewReader(nil)
		myReader.Reset(os.Stdin)
		content, e = myReader.ReadString('\n')

		if e != nil {
			return "", e
		}

		fields = strings.Fields(content)

		if fields == nil || len(fields) == 0 {
			choose = 0
			continue
		}

		choose, e = strconv.Atoi(fields[0])
	}

	return problems[choose-1].Index, nil
}
