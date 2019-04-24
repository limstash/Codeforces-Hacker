package app

import(
	"github.com/hytzongxuan/Codeforces-Hacker/module/contest"
	"github.com/hytzongxuan/Codeforces-Hacker/module/token"
	"net/http"
	"errors"
	"fmt"
)

func Load (cookie *[]*http.Cookie) ([]ReadContest.Contest, string, error){
	fmt.Println("[Info] Fetching contests info...")

	contests, err := ReadContest.GetContests(cookie)

	if err != nil {
		return nil, "", errors.New("[Error] Unable to fetching contest info")
	}

	fmt.Println("[Info] Fetching CSRF token...")
	CSRF, err := token.GetCSRF(cookie)

	if err != nil || CSRF == "" {
		return nil, "", errors.New("[Error] Unable to fetching CSRF token")
	}

	return contests, CSRF, nil
}