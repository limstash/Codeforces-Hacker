package token

import (
	"errors"
	"strings"

	"github.com/limstash/Codeforces-Hacker/module/conn"
	"github.com/opesun/goquery"

	. "github.com/limstash/Codeforces-Hacker/common"
)

// GetCSRF will fetch CSRF token from codeforces's login page
func GetCSRF(authentication *Authentication, server string) error {
	request := Request{}

	request.URL = server + "/enter?back=%2F"
	request.Method = "GET"
	request.NotRedirect = false
	request.Authentication = authentication

	response, err := conn.HTTPRequest(request)

	if err != nil {
		return err
	}

	html, err := goquery.Parse(strings.NewReader(string(response.ResponseBody)))

	if err != nil {
		return err
	}

	csrf := html.Find(".csrf-token").Eq(0).Attr("data-csrf")

	if csrf == "" {
		return errors.New("CSRF not exists")
	}

	authentication.CSRF = csrf
	return nil
}
