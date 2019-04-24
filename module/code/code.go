package code

import(
	"github.com/hytzongxuan/Codeforces-Hacker/module/http"
	"net/http"
	"strconv"
)

func QueryCode(submission int, cookie *[]*http.Cookie, CSRF string)(string, error){
	res, err := con.HttpPost("https://codeforces.com/data/submitSource", cookie, map[string]string{"X-Csrf-Token": CSRF, "X-Requested-With": "XMLHttpRequest", "Origin" : "https://codeforces.com", "Referer" : "https://codeforces.com/problemset/status", "Host" : "codeforces.com"}, map[string]string{"submissionId": strconv.Itoa(submission), "csrf_token": CSRF})

	if(err != nil){
		return "", err
	}

	return res, nil
}