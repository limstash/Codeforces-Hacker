package main

import (
	"fmt"
	"net/http"

	"github.com/hytzongxuan/Codeforces-Hacker/app"
)

func version() {
	fmt.Println("Codeforces Hacker v1.0")
}

// Cookie are the HTTP Cookies which saved from https://codeforces.com
var Cookie []*http.Cookie

func main() {
	version()

	Contest, CSRF, e := app.Load(&Cookie)

	if e != nil || CSRF == "" {
		fmt.Println(e)
		return
	}

	contestID, e := app.FindContest(Contest)

	if e != nil {
		fmt.Println(e)
		return
	}

	id, e := app.ChooseProblem(contestID, &Cookie)

	if e != nil {
		fmt.Println(e)
		return
	}

	panic(id)
}
