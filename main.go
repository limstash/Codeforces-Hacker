package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/hytzongxuan/Codeforces-Hacker/app"
)

func version() {
	fmt.Println("Codeforces Hacker v1.0")
}

// Cookie are the HTTP Cookies which saved from https://codeforces.com
var Cookie []*http.Cookie

func finish() {
	time.Sleep(time.Millisecond * 5000)
	os.Exit(0)
}

func main() {

	version()
	Contest, CSRF, e := app.Load(&Cookie)

	if e != nil || CSRF == "" {
		fmt.Println(e)
		finish()
	}

	contestInfo, e := app.FindContest(Contest)

	if e != nil {
		fmt.Println(e)
		finish()
	}

	choose, e := app.ChooseProblem(contestInfo.ID, &Cookie)

	if e != nil {
		fmt.Println(e)
		finish()
	}

	loginChoose, e := app.QueryLoginChoose()

	if e != nil {
		fmt.Println(e)
		finish()
	}

	var hackChoose bool

	if loginChoose == true {
		hackChoose, e = app.QueryHackChoose()

		if e != nil {
			fmt.Println(e)
			finish()
		}

		app.Login(&Cookie, CSRF)
	}

	app.SaveData()

	fmt.Println("[Info] Fetching submissions...")
	submit := app.GetSubmission(&Cookie, contestInfo, choose)

	fmt.Println(" ")

	if hackChoose == true {
		fmt.Println("[Info] Not Support auto hack")
	}

	for i := 0; i < len(submit); i++ {
		flag, _ := app.TestCode(submit[i].SubmissionID, submit[i].Language, false, &Cookie, CSRF)

		if flag == true {
			fmt.Println("[Info] Submission " + strconv.Itoa(submit[i].SubmissionID) + " Accepted")
		}

		if e != nil {
			if hackChoose == true {

			} else {

			}
		}
	}
}
