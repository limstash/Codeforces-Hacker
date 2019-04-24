package con

import (
	"testing"
	"net/http"
)

func Test_httpGET(t *testing.T){
	var GlobalCookie []*http.Cookie
	http_res, http_err := HttpGet("http://www.baidu.com", &GlobalCookie, map[string]string{"Host" : "www.baidu.com"})

	status := true

	if http_err != nil {
		t.Error(http_err)
		status = false
	}else if http_res == "" {
		t.Error("No Response from http://www.baidu.com")
		status = false
	}else if GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Error("Cookie is an empty field")
		status = false
	}

	if status == false {
		t.Log("HTTP GET Query Test Failed")
	}else{
		t.Log("HTTP GET Query Test Passed")
	}
}

func Test_httpsGET(t *testing.T){
	var GlobalCookie []*http.Cookie
	https_res, https_err := HttpGet("http://www.baidu.com", &GlobalCookie, map[string]string{"Host" : "www.baidu.com"})

	status := true

	if https_err != nil {
		t.Error(https_err)
		status = false
	}else if https_res == "" {
		t.Error("No Response from http://www.baidu.com")
		status = false
	}else if GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Error("Cookie is an empty field")
		status = false
	}

	if status == false {
		t.Log("HTTPS GET Query Test Failed")
	}else{
		t.Log("HTTPS GET Query Test Passed")
	}
}

func Test_httpPOST(t *testing.T){
	var GlobalCookie []*http.Cookie
	http_res, http_err := HttpPost("http://www.baidu.com", &GlobalCookie, map[string]string{"Host" : "www.baidu.com"}, map[string]string{"hi" : "hi"})

	status := true

	if http_err != nil {
		t.Error(http_err)
		status = false
	}else if http_res == "" {
		t.Error("No Response from http://www.baidu.com")
		status = false
	}else if GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Error("Cookie is an empty field")
		status = false
	}

	if status == false {
		t.Log("HTTP GET Query Test Failed")
	}else{
		t.Log("HTTP GET Query Test Passed")
	}
}

func Test_httpsPOST(t *testing.T){
	var GlobalCookie []*http.Cookie
	https_res, https_err := HttpPost("http://www.baidu.com", &GlobalCookie, map[string]string{"Host" : "www.baidu.com"}, map[string]string{"hi" : "hi"})

	status := true

	if https_err != nil {
		t.Error(https_err)
		status = false
	}else if https_res == "" {
		t.Error("No Response from http://www.baidu.com")
		status = false
	}else if GlobalCookie == nil || len(GlobalCookie) == 0 {
		t.Error("Cookie is an empty field")
		status = false
	}

	if 	status == false{
		t.Log("HTTPS GET Query Test Failed")
	}else{
		t.Log("HTTPS GET Query Test Passed")
	}
}