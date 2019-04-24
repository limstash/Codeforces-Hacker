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
	version()

	Contest, CSRF, e := app.Load(&Cookie)

	if e != nil || CSRF == "" {
		fmt.Println(e);
		return;
	}

	contestID, e := app.FindContest(Contest)

	if e != nil{
		fmt.Println(e);
		return;
	}

	id, e := app.ChooseProblem(contestID, &Cookie)

	if e != nil {
		fmt.Println(e);
		return;
	}

	panic(id)
}