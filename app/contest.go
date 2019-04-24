package app

import (
	"github.com/hytzongxuan/Codeforces-Hacker/module/contest"
	"fmt"
	"time"
	"strings"
	"errors"
)

func FindContest(contests []ReadContest.Contest)(error){
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

	if openHackingPhase < 12 * 24 * 3600 {
		return errors.New("[Info] Open hacking phase finished")
	}else{
		fmt.Println("[Info] Open hacking phase running")
	}

	return nil
}