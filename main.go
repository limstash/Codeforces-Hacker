package main

import(
	"github.com/hytzongxuan/Codeforces-Hacker/app"
	"net/http"
	"fmt"
)

func version(){
	fmt.Println("Codeforces Hacker v1.0")
}

var Cookie []*http.Cookie

func main(){
	Contest, CSRF, e := app.Load(&Cookie)

	if e != nil {
		fmt.Println(e);
		return;
	}

	e = app.FindContest(Contest)

	if e != nil {
		fmt.Println(e);
		return;
	}

	fmt.Println(CSRF)
}