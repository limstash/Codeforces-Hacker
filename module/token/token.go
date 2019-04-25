package token

import (
	"errors"
	"net/http"
	"strings"

	"github.com/hytzongxuan/Codeforces-Hacker/module/con"
	"github.com/opesun/goquery"
)

// GetCSRF will fetch CSRF token from codeforces's login page
func GetCSRF(cookie *[]*http.Cookie) (string, error) {

	body, e := con.HTTPGet("https://codeforces.com/enter?back=%2F", cookie, nil)

	if e != nil {
		return "", e
	}

	html, e := goquery.Parse(strings.NewReader(body))

	if e != nil {
		return "", e
	}

	csrf := html.Find(".csrf-token").Eq(0).Attr("data-csrf")

	if csrf == "" {
		return "", errors.New("CSRF is an empty field")
	}

	return csrf, nil
}
