package token

import (
	"github.com/opesun/goquery"
	"github.com/hytzongxuan/Codeforces-Hacker/module/http"
	"net/http"
	"strings"
	"errors"
)

func GetCSRF(cookie *[]*http.Cookie)(string, error){

	body, e := codeforces.HttpGet("https://codeforces.com/enter?back=%2F", cookie, nil)

	if e != nil{
		return "", e
	}
	
	html, e := goquery.Parse(strings.NewReader(body))

	if e != nil{
		return "", e
	}

	csrf := html.Find(".csrf-token").Eq(0).Attr("data-csrf")
	
	if csrf == ""{
		return "", errors.New("CSRF is an empty field")
	}
	
	return csrf, nil
}
