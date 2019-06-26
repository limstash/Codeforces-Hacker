package conn

import (
	"testing"

	"github.com/bitly/go-simplejson"
	. "github.com/hytzongxuan/Codeforces-Hacker/common"
)

var authentication Authentication

func testPOST01(t *testing.T) bool {
	request := Request{}

	request.URL = "https://api.yellowfish.top/Codeforces-Hacker/post-test.php"
	request.Method = "POST"
	request.NotRedirect = false
	request.Authentication = &authentication
	request.Data = map[string]string{"post-data01": "test01", "post-data02": "test02"}

	response, err := httpRequest(request)

	if err != nil {
		t.Error(err)
		return false
	}

	js, err := simplejson.NewJson(response.ResponseBody)

	if err != nil {
		t.Skip("Something goes wrong, Skipped")
		return true
	}

	postdata01 := js.Get("post-data01").MustString()
	postdata02 := js.Get("post-data02").MustString()

	if postdata01 != "test01" || postdata02 != "test02" {
		t.Error("Test Failed: (Case 01) POST body not exists or format error")
		return false
	}

	if response.RedirectStatus == true {
		t.Error("Test Failed: (Case 01) Redirect should not be found here")
		return false
	}

	return true
}

func testPOST02(t *testing.T) bool {
	request := Request{}

	request.URL = "https://api.yellowfish.top/Codeforces-Hacker/not-found.php"
	request.Method = "POST"
	request.NotRedirect = false
	request.Authentication = &authentication

	_, err := httpRequest(request)

	if err == nil || err.Error() != "Remote server return with code 404" {
		t.Error("Test Failed: (Case 02) httpRequest should throw an error here (http 404)")
		return false
	}

	return true
}

func testPOST03(t *testing.T) bool {
	request := Request{}

	request.URL = "https://api.yellowfish.top/Codeforces-Hacker/"
	request.Method = "POST"
	request.NotRedirect = false
	request.Authentication = &authentication

	_, err := httpRequest(request)

	if err == nil || err.Error() != "Remote server return with code 403" {
		t.Error("Test Failed: (Case 03) httpRequest should throw an error here (http 403)")
		return false
	}

	return true
}

func testPOST04(t *testing.T) bool {
	request := Request{}

	request.URL = "https://api.yellowfish.top/Codeforces-Hacker/cookie-test.php"
	request.Method = "POST"
	request.NotRedirect = false
	request.Authentication = &authentication

	_, err := httpRequest(request)

	if err != nil {
		t.Error(err)
		return false
	}

	response, err := httpRequest(request)

	js, err := simplejson.NewJson(response.ResponseBody)

	if err != nil {
		t.Skip("Something goes wrong, Skipped")
		return true
	}

	cookiedata01 := js.Get("cookie-test01").MustString()
	cookiedata02 := js.Get("cookie-test02").MustString()

	if cookiedata01 != "test01" || cookiedata02 != "test02" {
		t.Error("Test Failed: (Case 04) Cookie not exists or format error")
		return false
	}

	if response.RedirectStatus == true {
		t.Error("Test Failed: (Case 04) Redirect should not be found here")
		return false
	}

	return true
}

func testGet05(t *testing.T) bool {
	request := Request{}

	request.URL = "https://api.yellowfish.top/Codeforces-Hacker/header-test.php"
	request.Method = "GET"
	request.NotRedirect = false
	request.Authentication = &authentication
	request.Header = map[string]string{"Header-Data01": "test01", "Header-Data02": "test02"}

	response, err := httpRequest(request)

	if err != nil {
		t.Error(err)
		return false
	}

	js, err := simplejson.NewJson(response.ResponseBody)

	if err != nil {
		t.Skip("Something goes wrong, Skipped")
		return true
	}

	headerdata01 := js.Get("Header-Data01").MustString()
	headerdata02 := js.Get("Header-Data02").MustString()

	if headerdata01 != "test01" || headerdata02 != "test02" {
		t.Error("Test Failed: (Case 05) Header not exists or format error")
		return false
	}

	if response.RedirectStatus == true {
		t.Error("Test Failed: (Case 05) Redirect should not be found here")
		return false
	}
	return true
}

func testGet06(t *testing.T) bool {
	request := Request{}

	request.URL = "https://api.yellowfish.top"
	request.Method = "GET"
	request.NotRedirect = true
	request.Authentication = &authentication

	response, err := httpRequest(request)

	if err != nil {
		t.Error(err)
		return false
	}

	if response.RedirectStatus == false {
		t.Error("Test Failed: (Case 06) Redirect should be found here")
		return false
	}

	return true
}

func Test_httpRequest(t *testing.T) {
	status := true
	authentication = Authentication{}

	if testPOST01(t) == false {
		status = false
	}

	if testPOST02(t) == false {
		status = false
	}

	if testPOST03(t) == false {
		status = false
	}

	if testPOST04(t) == false {
		status = false
	}

	if testGet05(t) == false {
		status = false
	}

	if testGet06(t) == false {
		status = false
	}

	if status == true {
		t.Log("Package conn - httpRequest test passed")
	} else {
		t.Log("Package conn - httpRequest test failed")
	}
}

func testQuery01(t *testing.T) bool {
	request := Request{}

	request.URL = "https://www.baidu.com"
	request.Method = "GET"
	request.NotRedirect = false
	request.Authentication = &authentication

	_, err := HTTPRequest(request)

	if err != nil {
		t.Error(err)
		return false
	}

	return true
}

func testQuery02(t *testing.T) bool {
	request := Request{}

	request.URL = "https://www.baidu.as-i/ppp"
	request.Method = "POST"
	request.NotRedirect = false
	request.Authentication = &authentication

	_, err := HTTPRequest(request)

	if err == nil {
		t.Error("Test Failed: (Case 03) HttpRequest should throw an error here")
		return false
	}

	return true
}

func Test_HTTPRequest(t *testing.T) {
	status := true
	authentication = Authentication{}

	if testQuery01(t) == false {
		status = false
	}

	if testQuery02(t) == false {
		status = false
	}

	if status == true {
		t.Log("Package conn - HTTPRequest test passed")
	} else {
		t.Log("Package conn - HTTPRequest test failed")
	}
}
