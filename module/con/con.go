package con

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// HTTPGetByte will send a GET query to remote server and return in []byte
func HTTPGetByte(uri string, cookie *[]*http.Cookie, header map[string]string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", uri, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	req.Header.Set("Accept-Encoding", "none")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("User-Agent", "Mozilla/5.s0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36")

	for headerName, headerValue := range header {
		req.Header.Set(headerName, headerValue)
	}

	for _, v := range *cookie {
		req.AddCookie(v)
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, errors.New("Return HTTP Code " + strconv.Itoa(resp.StatusCode))
	}

	newCookie := resp.Cookies()

	for i := 0; i < len(newCookie); i++ {
		*cookie = append(*cookie, newCookie[i])
	}

	return body, nil
}

// HTTPPostByte will send a POST query to remote server and return in []byte
func HTTPPostByte(uri string, cookie *[]*http.Cookie, header map[string]string, data map[string]string) ([]byte, error) {
	client := &http.Client{}

	var postdata http.Request
	postdata.ParseForm()

	for dataName, dataValue := range data {
		postdata.Form.Add(dataName, dataValue)
	}

	bodystr := strings.TrimSpace(postdata.Form.Encode())
	req, err := http.NewRequest("POST", uri, strings.NewReader(bodystr))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Encoding", "none")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.s0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36")

	for headerName, headerValue := range header {
		req.Header.Set(headerName, headerValue)
	}

	for _, v := range *cookie {
		req.AddCookie(v)
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("Return HTTP Code " + strconv.Itoa(resp.StatusCode))
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	newCookie := resp.Cookies()

	for i := 0; i < len(newCookie); i++ {
		*cookie = append(*cookie, newCookie[i])
	}

	return body, nil
}

// HTTPGet will send a Get query to remote server and return in string
func HTTPGet(uri string, cookie *[]*http.Cookie, header map[string]string) (string, error) {
	a, b := HTTPGetByte(uri, cookie, header)
	return string(a), b
}

// HTTPPost will send a POST query to remote server and return in string
func HTTPPost(uri string, cookie *[]*http.Cookie, header map[string]string, data map[string]string) (string, error) {
	a, b := HTTPPostByte(uri, cookie, header, data)
	return string(a), b
}
