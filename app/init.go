package app

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hytzongxuan/Codeforces-Hacker/module/contest"
	"github.com/hytzongxuan/Codeforces-Hacker/module/token"
)

func getPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	rst := filepath.Dir(path)
	return rst
}

// GetCSRF will fetch CSRF token from https://codeforces.com
func GetCSRF(cookie *[]*http.Cookie) (string, error) {
	fmt.Println("[Info] Fetching CSRF token...")

	CSRF, err := token.GetCSRF(cookie)

	if err != nil || CSRF == "" {
		return "", errors.New("[Error] Unable to fetching CSRF token")
	}

	return CSRF, nil
}

// Load will fetch contests info and CSRF token from https://codeforces.com
func Load(cookie *[]*http.Cookie) ([]contest.Contest, string, error) {
	fmt.Println("[Info] Fetching contests info...")

	contests, err := contest.GetContests(cookie)

	if err != nil {
		return nil, "", errors.New("[Error] Unable to fetching contest info")
	}

	CSRF, err := GetCSRF(cookie)

	if err != nil {
		return nil, "", err
	}

	return contests, CSRF, nil
}
