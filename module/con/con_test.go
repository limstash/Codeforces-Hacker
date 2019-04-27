package con

import (
	"net/http"
	"testing"
)

func Test_httpGET(t *testing.T) {
	var GlobalCookie []*http.Cookie

	res, e := HTTPGet("http://www.baidu.com", &GlobalCookie, map[string]string{"Host": "www.baidu.com"})

	status := true
	if e != nil {
		t.Error(e)
		status = false
	} else if res == "" {
		t.Error("No Response from http://www.baidu.com")
		status = false
	} else if GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Error("Cookie is an empty field")
		status = false
	}

	if status == false {
		t.Log("HTTP GET Query Test Failed")
	} else {
		t.Log("HTTP GET Query Test Passed")
	}
}

func Test_httpsGET(t *testing.T) {
	var GlobalCookie []*http.Cookie
	res, e := HTTPGet("http://www.baidu.com", &GlobalCookie, map[string]string{"Host": "www.baidu.com"})

	status := true

	if e != nil {
		t.Error(e)
		status = false
	} else if res == "" {
		t.Error("No Response from http://www.baidu.com")
		status = false
	} else if GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Error("Cookie is an empty field")
		status = false
	}

	if status == false {
		t.Log("HTTPS GET Query Test Failed")
	} else {
		t.Log("HTTPS GET Query Test Passed")
	}
}

func Test_httpPOST(t *testing.T) {
	var GlobalCookie []*http.Cookie
	res, e := HTTPPost("http://www.baidu.com", &GlobalCookie, map[string]string{"Host": "www.baidu.com"}, map[string]string{"hi": "hi"})

	status := true

	if e != nil {
		t.Error(e)
		status = false
	} else if res == "" {
		t.Error("No Response from http://www.baidu.com")
		status = false
	} else if GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Error("Cookie is an empty field")
		status = false
	}

	if status == false {
		t.Log("HTTP POST Query Test Failed")
	} else {
		t.Log("HTTP POST Query Test Passed")
	}
}

func Test_httpsPOST(t *testing.T) {
	var GlobalCookie []*http.Cookie
	res, e := HTTPPost("http://www.baidu.com", &GlobalCookie, map[string]string{"Host": "www.baidu.com"}, map[string]string{"hi": "hi"})

	status := true

	if e != nil {
		t.Error(e)
		status = false
	} else if res == "" {
		t.Error("No Response from http://www.baidu.com")
		status = false
	} else if GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Error("Cookie is an empty field")
		status = false
	}

	if status == false {
		t.Log("HTTPS POST Query Test Failed")
	} else {
		t.Log("HTTPS POST Query Test Passed")
	}
}
