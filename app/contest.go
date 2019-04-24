package app

import (
	"github.com/hytzongxuan/Codeforces-Hacker/module/contest"
	"fmt"
	"time"
	"strings"
	"errors"
	"net/http"
	"strconv"
)

func FindContest(contests []contest.Contest)(int, error){
	contests_siz := len(contests)
		
	var lastContestTime int64 = 100000000
	lastContestIndex := 0

	currentTime := time.Now().Unix() 

	for i := 0; i < contests_siz; i++{
		if strings.Contains(contests[i].Name, "Educational") || strings.Contains(contests[i].Name, "Div. 3"){
			timeDelta := currentTime - contests[i].StartTimeSeconds

			if timeDelta > 0 && timeDelta < lastContestTime {
				lastContestTime = timeDelta
				lastContestIndex = i
			}

		}
	}

	fmt.Println("[Info] The current contest is "+contests[lastContestIndex].Name)

	openHackingPhase := currentTime - contests[lastContestIndex].StartTimeSeconds - contests[lastContestIndex].DurationSeconds

	if openHackingPhase > 12 * 24 * 3600 {
		return 0, errors.New("[Info] Open hacking phase finished")
	}else{
		fmt.Println("[Info] Open hacking phase running")
	}

	return contests[lastContestIndex].ID, nil;
}

func ChooseProblem(ContestID int, cookie *[]*http.Cookie) (int, error){
	problems, e := contest.GetProblems(ContestID, cookie)

	if e != nil {
		return 0, e;
	}

	fmt.Println("")
	fmt.Println("Please choose the problem you want to hack")
	fmt.Println("")

	for i := 0; i < len(problems); i++ {
		fmt.Println("[" + strconv.Itoa(i) + "] " + problems[i].Index + ". " + problems[i].Name)
	}

	var choose int
	fmt.Println("")

	fmt.Printf("Your Choose is [0-" + strconv.Itoa(len(problems)-1) + "] : ")
	_, e = fmt.Scanf("%d", &choose)
	fmt.Println("")

	for ; e != nil || choose < 0 || choose >= len(problems); {
		fmt.Printf("Your Choose is [0-" + strconv.Itoa(len(problems)-1) + "] : ")
		_, e = fmt.Scanf("%d", &choose)
		fmt.Println("")
	}

	return choose, nil
}