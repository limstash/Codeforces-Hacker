package judge

import (
	"bytes"
	"errors"
	"os/exec"
	"regexp"
	"strconv"

	. "github.com/limstash/Codeforces-Hacker/common"
)

func getVersion(command string) (string, error) {
	res := exec.Command(command, "--version")

	var out bytes.Buffer
	res.Stdout = &out

	err := res.Run()

	if err != nil {
		return "", errors.New("Not Available")
	}

	flysnowRegexp := regexp.MustCompile(`(\d+).(\d+).(\d+)`)
	params := flysnowRegexp.FindStringSubmatch(out.String())

	if params == nil {
		return "", err
	}

	return params[0], nil
}

func splitVersion(ver string) (int, int, int) {
	flysnowRegexp := regexp.MustCompile(`(\d+).(\d+).(\d+)`)
	params := flysnowRegexp.FindStringSubmatch(ver)

	first, _ := strconv.Atoi(params[1])
	second, _ := strconv.Atoi(params[2])
	third, _ := strconv.Atoi(params[3])

	return first, second, third
}

func getGCCVersion() (int, int, int, error) {
	res, err := getVersion("gcc")

	if err != nil {
		return 0, 0, 0, err
	}

	first, second, third := splitVersion(res)
	return first, second, third, nil
}

func getGPlusPlusVersion() (int, int, int, error) {
	res, err := getVersion("g++")

	if err != nil {
		return 0, 0, 0, err
	}

	first, second, third := splitVersion(res)
	return first, second, third, nil
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

func checkGCC(first int, second int, third int) int {
	if first >= 7 {
		return 3
	}

	if first >= 5 {
		return 2
	}

	if first >= 5 || (first == 4 && (second > 8 || (second == 8 && third >= 1))) {
		return 1
	}

	return 0
}

func GetSupport() Language {
	var support Language

	a, b, c, e := getGPlusPlusVersion()

	if e != nil {
		support.GNUCPP11 = false
		support.GNUCPP14 = false
		support.GNUCPP17 = false
	} else {
		support.GNUCPP11 = checkGCC(a, b, c) >= 1
		support.GNUCPP14 = checkGCC(a, b, c) >= 2
		support.GNUCPP17 = checkGCC(a, b, c) >= 3
	}

	a, b, c, e = getGCCVersion()

	if e != nil {
		support.GNUC11 = false
	} else {
		support.GNUC11 = checkGCC(a, b, c) >= 1
	}

	support.Go = checkGo()
	support.Python2 = checkPython2()
	support.Python3 = checkPython3()

	return support
}
