package app

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/hytzongxuan/Codeforces-Hacker/module/conn"
)

func QueryLoginChoose() (bool, error) {
	return GetUserResponseYN("\n", "Do you want to log in to your codeforces account (yes/no): ", "")
}

func QueryHackChoose() (bool, error) {
	return GetUserResponseYN("", "Do you want to enable automatic hacking (yes/no): ", "")
}

func queryUsername() (string, error) {
	return GetUserResponseString("\n", "Username: ", "")
}

func queryPassword() (string, error) {
	return GetUserResponseString("", "Password: ", "")
}

func submitLogin(username string, password string, Cookie *[]*http.Cookie, CSRF string) (bool, string, error) {
	resp, redirect, e := conn.HTTPPostNR("https://codeforces.com/enter?back=%2F", Cookie, map[string]string{"X-Csrf-Token": CSRF, "X-Requested-With": "XMLHttpRequest", "Origin": "https://codeforces.com", "Referer": "https://codeforces.com/problemset/status", "Host": "codeforces.com"}, map[string]string{"action": "enter", "csrf_token": CSRF, "handleOrEmail": username, "password": password})

	if e != nil {
		return false, "", e
	}

	flysnowRegexp := regexp.MustCompile(`Specify correct handle or email`)
	params := flysnowRegexp.FindStringSubmatch(resp)

	if len(params) > 0 {
		return false, "Specify correct handle or email", nil
	}

	flysnowRegexp = regexp.MustCompile(`Invalid handle/email or password`)
	params = flysnowRegexp.FindStringSubmatch(resp)

	if len(params) > 0 {
		return false, "Invalid handle/email or password", nil
	}

	if redirect {
		return true, "", nil
	}

	return false, "Unknown", nil
}

func Login(Cookie *[]*http.Cookie, CSRF string) error {
	for true {
		username, _ := queryUsername()
		password, _ := queryPassword()

		status, msg, e := submitLogin(username, password, Cookie, CSRF)

		if e != nil {
			return e
		}

		fmt.Println("")

		if status == true {
			fmt.Println("[Info] Login Success")
			break
		}

		if msg != "" {
			fmt.Println("[Info] " + msg + ", please retry")
			continue
		}

		fmt.Println("[Info] Unknown Error")
	}

	return nil
}
