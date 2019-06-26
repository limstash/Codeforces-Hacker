package app

import (
	"errors"
	"regexp"

	"github.com/hytzongxuan/Codeforces-Hacker/module/conn"
	"github.com/hytzongxuan/Codeforces-Hacker/module/token"

	. "github.com/hytzongxuan/Codeforces-Hacker/common"
)

func apiSubmitLogin(account Account, auth *Authentication, server string) (Response, error) {
	request := Request{}

	request.URL = server + "/enter?back=%2F"
	request.Method = "POST"
	request.NotRedirect = true
	request.Authentication = auth
	request.Data = map[string]string{"action": "enter", "csrf_token": auth.CSRF, "handleOrEmail": account.Username, "password": account.Password}
	request.Header = map[string]string{"X-Csrf-Token": auth.CSRF, "X-Requested-With": "XMLHttpRequest", "Origin": "https://codeforces.com", "Referer": "https://codeforces.com/problemset/status", "Host": "codeforces.com"}

	response, err := conn.HTTPRequest(request)
	return response, err
}

func submitLogin(account Account, auth *Authentication, server string) (bool, error) {
	response, err := apiSubmitLogin(account, auth, server)

	if err != nil {
		return false, err
	}

	flysnowRegexp := regexp.MustCompile(`Specify correct handle or email`)
	params := flysnowRegexp.FindStringSubmatch(string(response.ResponseBody))

	if len(params) > 0 {
		return false, errors.New("Specify correct handle or email")
	}

	flysnowRegexp = regexp.MustCompile(`Invalid handle/email or password`)
	params = flysnowRegexp.FindStringSubmatch(string(response.ResponseBody))

	if len(params) > 0 {
		return false, errors.New("Invalid handle/email or password")
	}

	if response.RedirectStatus {
		return true, nil
	}

	return false, errors.New("Unknown Error")
}

func Login(config Config, auth *Authentication) {
	if config.IsAutoLogin {
		status, err := submitLogin(config.Account, auth, config.Server)

		if status == false {
			log(1, err.Error())
		}

		token.GetCSRF(auth, config.Server)
	}
}
