package judge

import (
	"bytes"
	"errors"
	"os/exec"
	"regexp"
	"strconv"
)

type Language struct {
	GNUC11   bool
	GNUCPP11 bool
	GNUCPP14 bool
	GNUCPP17 bool
	Python2  bool
	Python3  bool
	Go       bool
}

func getGPlusPlusVersion() (string, error) {
	res := exec.Command("g++", "--version")

	var out bytes.Buffer
	res.Stdout = &out

	err := res.Run()

	if err != nil {
		return "", errors.New("Not Available")
	}

	flysnowRegexp := regexp.MustCompile(`[\d+].[\d+].[\d+]`)
	params := flysnowRegexp.FindStringSubmatch(out.String())

	if params == nil {
		return "", errors.New("Not Available")
	}

	return params[0], nil
}

func getGCCVersion() (string, error) {
	res := exec.Command("gcc", "--version")

	var out bytes.Buffer
	res.Stdout = &out

	err := res.Run()

	if err != nil {
		return "", errors.New("Not Available")
	}

	flysnowRegexp := regexp.MustCompile(`[\d+].[\d+].[\d+]`)
	params := flysnowRegexp.FindStringSubmatch(out.String())

	if params == nil {
		return "", errors.New("Not Available")
	}

	return params[0], nil
}

func checkPython2() bool {
	res := exec.Command("python2", "--version")

	var out bytes.Buffer
	res.Stdout = &out

	err := res.Run()

	if err != nil {
		return false
	}

	return true
}

func checkPython3() bool {
	res := exec.Command("python3", "--version")

	var out bytes.Buffer
	res.Stdout = &out

	err := res.Run()

	if err != nil {
		return false
	}

	return true
}

func checkGo() bool {
	res := exec.Command("go", "version")

	var out bytes.Buffer
	res.Stdout = &out

	err := res.Run()

	if err != nil {
		return false
	}

	return true
}

func checkC11(ver string) bool {
	flysnowRegexp := regexp.MustCompile(`([\d+]).([\d+]).([\d+])`)
	params := flysnowRegexp.FindStringSubmatch(ver)

	first, _ := strconv.Atoi(params[1])
	second, _ := strconv.Atoi(params[2])
	third, _ := strconv.Atoi(params[3])

	if first >= 5 || (first == 4 && (second > 8 || (second == 8 && third >= 1))) {
		return true
	}

	return false
}

func checkCPP11(ver string) bool {
	flysnowRegexp := regexp.MustCompile(`([\d+]).([\d+]).([\d+])`)
	params := flysnowRegexp.FindStringSubmatch(ver)

	first, _ := strconv.Atoi(params[1])
	second, _ := strconv.Atoi(params[2])
	third, _ := strconv.Atoi(params[3])

	if first >= 5 || (first == 4 && (second > 8 || (second == 8 && third >= 1))) {
		return true
	}

	return false
}

func checkCPP14(ver string) bool {
	flysnowRegexp := regexp.MustCompile(`([\d+]).([\d+]).([\d+])`)
	params := flysnowRegexp.FindStringSubmatch(ver)

	first, _ := strconv.Atoi(params[1])

	if first >= 5 {
		return true
	}

	return false
}

func checkCPP17(ver string) bool {
	flysnowRegexp := regexp.MustCompile(`([\d+]).([\d+]).([\d+])`)
	params := flysnowRegexp.FindStringSubmatch(ver)

	first, _ := strconv.Atoi(params[1])

	if first >= 7 {
		return true
	}

	return false
}

func GetAvailableLanguage() Language {
	var support Language

	ver, e := getGPlusPlusVersion()

	if e != nil {
		support.GNUCPP11 = false
		support.GNUCPP14 = false
		support.GNUCPP17 = false
	} else {
		support.GNUCPP11 = checkCPP11(ver)
		support.GNUCPP14 = checkCPP14(ver)
		support.GNUCPP17 = checkCPP17(ver)
	}

	ver, e = getGCCVersion()

	if e != nil {
		support.GNUC11 = false
	} else {
		support.GNUC11 = checkC11(ver)
	}

	support.Go = checkGo()
	support.Python2 = checkPython2()
	support.Python3 = checkPython3()

	return support
}
