package app

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"testing"
)

var Cookie []*http.Cookie

func Test_TestCode(t *testing.T) {
	fmt.Println(getPath())
	os.MkdirAll(getPath()+"/src", 0777)

	inputFile, e := os.OpenFile(getPath()+"/src/data.in", os.O_WRONLY|os.O_CREATE, 0666)

	if e != nil {
		t.Skip("Create data.in Failed, Skipped")
	}

	defer inputFile.Close()
	inputWriter := bufio.NewWriter(inputFile)
	inputWriter.WriteString("9\n1 3 3 6 7 6 8 8 9")
	inputWriter.Flush()

	outputFile, e := os.OpenFile(getPath()+"/src/data.ans", os.O_WRONLY|os.O_CREATE, 0666)

	if e != nil {
		t.Skip("Create data.ans Failed, Skipped")
	}

	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputWriter.WriteString("4 \n")
	outputWriter.Flush()

	_, CSRF, e := Load(&Cookie)

	if e != nil {
		t.Skip("Fetch CSRF Failed, Skipped")
	}

	status, e := TestCode(51737890, "Python 3", false, &Cookie, CSRF)

	if status == false || (e != nil && e.Error() != "Not Support") {
		t.Error("App judge failed at test Submission 51737890 with Python3")
		if e != nil {
			t.Error(e)
		}
	} else {
		t.Log("App judge passed at test Submission 51737890 with Python3")
	}

	status, e = TestCode(51837213, "Python 2", false, &Cookie, CSRF)

	if status == false || (e != nil && e.Error() != "Not Support") {
		t.Error("App judge failed at test Submission 51837213 with Python2")
		if e != nil {
			t.Error(e)
		}
	} else {
		t.Log("App judge passed at test Submission 51837213 with Python2")
	}

	status, e = TestCode(51812034, "GNU C++17", false, &Cookie, CSRF)

	if status == false || (e != nil && e.Error() != "Not Support") {
		t.Error("App judge failed at test Submission 51812034 with GNU C++17")
		if e != nil {
			t.Error(e)
		}
	} else {
		t.Log("App judge passed at test Submission 51812034 with GNU C++17")
	}

	status, e = TestCode(51856295, "GNU C++14", false, &Cookie, CSRF)

	if status == false || (e != nil && e.Error() != "Not Support") {
		t.Error("App judge failed at test Submission 51856295 with GNU C++14")
		if e != nil {
			t.Error(e)
		}
	} else {
		t.Log("App judge passed at test Submission 51856295 with GNU C++14")
	}

	status, e = TestCode(51726415, "GNU C++11", false, &Cookie, CSRF)

	if status == false || (e != nil && e.Error() != "Not Support") {
		t.Error("App judge failed at test Submission 51726415 with GNU C++11")
		if e != nil {
			t.Error(e)
		}
	} else {
		t.Log("App judge passed at test Submission 51726415 with GNU C++11")
	}

	status, e = TestCode(51731138, "Go", false, &Cookie, CSRF)

	if status == false || (e != nil && e.Error() != "Not Support") {
		t.Error("App judge failed at test Submission 51731138 with Go")
		if e != nil {
			t.Error(e)
		}
	} else {
		t.Log("App judge passed at test Submission 51731138 with Go")
	}

	status, e = TestCode(52556322, "GNU C11", false, &Cookie, CSRF)

	if status == false || (e != nil && e.Error() != "Not Support") {
		t.Error("App judge failed at test Submission 52556322 with GNU C11")
		if e != nil {
			t.Error(e)
		}
	} else {
		t.Log("App judge passed at test Submission 52556322 with GNU C11")
	}

}
