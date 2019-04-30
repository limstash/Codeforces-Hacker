package main

import (
	"fmt"
	"net/http"
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
}

func main() {

	version()
	Contest, CSRF, e := app.Load(&Cookie)

	if e != nil || CSRF == "" {
		fmt.Println(e)
		finish()
	}

	contestID, e := app.FindContest(Contest)

	if e != nil {
		fmt.Println(e)
		finish()
	}

	id, e := app.ChooseProblem(contestID, &Cookie)

	if e != nil {
		fmt.Println(e)
		finish()
	}

	loginChoose, e := app.QueryLoginChoose()

	if e != nil {
		fmt.Println(e)
		finish()
	}

	if loginChoose == true {
		hackChoose, e := app.QueryHackChoose()

		if e != nil {
			fmt.Println(e)
			finish()
		}

		app.Login(&Cookie, CSRF)
		fmt.Println(hackChoose)
	}

	app.SaveData()
	fmt.Println(id)

}
