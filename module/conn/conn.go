package conn

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	. "github.com/hytzongxuan/Codeforces-Hacker/common"
)

func redirectRules(req *http.Request, via []*http.Request) error {
	if len(via) >= 0 {
		return errors.New("No Redirect")
	}
	return nil
}

func addCommonHeader(req *http.Request) {
	req.Header.Set("Accept-Encoding", "none")
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")
}

func addExtraHeader(header map[string]string, req *http.Request) {
	for headerName, headerValue := range header {
		req.Header.Set(headerName, headerValue)
	}
}

func addExtraCookie(authentication *Authentication, req *http.Request) {
	for _, v := range authentication.Cookie {
		req.AddCookie(v)
	}
}

func addPostBody(data map[string]string) string {
	var postdata http.Request

	postdata.ParseForm()

	for dataName, dataValue := range data {
		postdata.Form.Add(dataName, dataValue)
	}

	bodystr := strings.TrimSpace(postdata.Form.Encode())
	return bodystr
}

func httpRequest(request Request) (Response, error) {
	var response = Response{}
	var client *http.Client

	if request.NotRedirect {
		client = &http.Client{
			CheckRedirect: redirectRules,
		}
	} else {
		client = &http.Client{}
	}

	var bodystr string

	if request.Method == "POST" {
		bodystr = addPostBody(request.Data)
	}

	req, err := http.NewRequest(request.Method, request.URL, strings.NewReader(bodystr))

	if err != nil {
		return response, err
	}

	addCommonHeader(req)
	addExtraHeader(request.Header, req)
	addExtraCookie(request.Authentication, req)

	if request.Method == "POST" {
		req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	}

	resp, err := client.Do(req)

	if err != nil {
		flysnowRegexp := regexp.MustCompile(`No Redirect`)
		params := flysnowRegexp.FindStringSubmatch(err.Error())

		if len(params) == 0 {
			return response, err
		} else {
			response.RedirectStatus = true
		}
	}

	if resp.StatusCode >= 400 {
		return response, errors.New("Remote server return with code " + strconv.Itoa(resp.StatusCode))
	}

	if response.RedirectStatus {
		return response, nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return response, err
	}

	newCookie := resp.Cookies()

	for i := 0; i < len(newCookie); i++ {
		request.Authentication.Cookie = append(request.Authentication.Cookie, newCookie[i])
	}

	response.RedirectStatus = false
	response.ResponseBody = body

	return response, nil
}

func HTTPRequest(request Request) (Response, error) {
	var response Response
	var err error

	for i := 1; i <= 3; i++ {
		response, err = httpRequest(request)

		if err == nil {
			return response, nil
		}
	}

	return response, err
}
