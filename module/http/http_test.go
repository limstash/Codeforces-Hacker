package con

import (
	"testing"
	"net/http"
)

func Test_httpGET(t *testing.T){
	var GlobalCookie []*http.Cookie
	http_res, http_err := HttpGet("http://www.baidu.com", &GlobalCookie, nil)

	if http_err != nil {
		t.Error(http_err)
	}else if http_res == "" {
		t.Error("No Response from http://www.baidu.com")
	}else if GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Error("Cookie is an empty field")
	}

	if http_err != nil || http_res == "" || GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Log("HTTP GET Query Test Failed")
	}else{
		t.Log("HTTP GET Query Test Passed")
	}
}

func Test_httpsGET(t *testing.T){
	var GlobalCookie []*http.Cookie
	https_res, https_err := HttpGet("http://www.baidu.com", &GlobalCookie, nil)

	if https_err != nil {
		t.Error(https_err)
	}else if https_res == "" {
		t.Error("No Response from http://www.baidu.com")
	}else if GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Error("Cookie is an empty field")
	}

	if https_err != nil || https_res == "" || GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Log("HTTPS GET Query Test Failed")
	}else{
		t.Log("HTTPS GET Query Test Passed")
	}
}

func Test_httpPOST(t *testing.T){
	var GlobalCookie []*http.Cookie
	http_res, http_err := HttpPost("http://www.baidu.com", &GlobalCookie, nil, nil)

	if http_err != nil {
		t.Error(http_err)
	}else if http_res == "" {
		t.Error("No Response from http://www.baidu.com")
	}else if GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Error("Cookie is an empty field")
	}

	if http_err != nil || http_res == "" || GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Log("HTTP GET Query Test Failed")
	}else{
		t.Log("HTTP GET Query Test Passed")
	}
}

func Test_httpsPOST(t *testing.T){
	var GlobalCookie []*http.Cookie
	https_res, https_err := HttpPost("http://www.baidu.com", &GlobalCookie, nil, nil)

	if https_err != nil {
		t.Error(https_err)
	}else if https_res == "" {
		t.Error("No Response from http://www.baidu.com")
	}else if GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Error("Cookie is an empty field")
	}

	if https_err != nil || https_res == "" || GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Log("HTTPS GET Query Test Failed")
	}else{
		t.Log("HTTPS GET Query Test Passed")
	}
}