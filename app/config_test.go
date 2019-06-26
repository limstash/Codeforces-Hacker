package app

import (
	"io/ioutil"
	"os"
	"testing"
)

func prepare() error {

	data1 := "{\"contest\":1,\"problem\":\"A\",\"autoLogin\":true,\"autoHack\":false,\"testcase\":{\"inputFile\":\"./data/input\",\"outputFile\":\"./data/output\"},\"account\":{\"username\":\"user\",\"password\":\"pass\"}}"

	data2 := "{\"contest\":\"1\",\"problem\":\"A\",\"autoLogin\":true,\"autoHack\":false,\"testcase\":{\"inputFile\":\"./data/input\",\"outputFile\":\"./data/output\"},\"account\":{\"username\":\"user\",\"password\":\"pass\"}}"

	data3 := "{\"contest\":1,\"problem\":\"A\",\"autoLogin\":true,\"autoHack\":false,\"testcase\":{\"inputFile\":\"./data/input\",\"outputFile\":\"./data/output\"},\"account\":{\"username\":\"user\",\"password\":\"pass\"}q}"

	err := ioutil.WriteFile("config-test-01.json", []byte(data1), 0644)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile("config-test-02.json", []byte(data2), 0644)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile("config-test-03.json", []byte(data3), 0644)

	if err != nil {
		return err
	}

	return nil
}

func clean() {
	os.RemoveAll("config-test-01.json")
	os.RemoveAll("config-test-02.json")
	os.RemoveAll("config-test-03.json")
}

func Test_readConfig(t *testing.T) {
	err := prepare()

	if err != nil {
		t.Skip("Init Test Failed")
	}

	status := true
	_, err = readConfig("config-test-01.json")

	if err != nil {
		t.Error(err)
		status = false
	}

	_, err = readConfig("config-test.json")

	if err == nil {
		t.Error("Test Failed: Should Return Error Here")
		status = false
	}

	if status == true {
		t.Log("Package app - readConfig test Passed")
	} else {
		t.Log("Package app - readConfig test Failed")
	}

	clean()
}

func Test_loadConfig(t *testing.T) {
	err := prepare()

	if err != nil {
		t.Skip("Init Test Failed")
	}

	status := true
	_, err = LoadConfig("config-test.json")

	if err == nil {
		t.Error("Test Failed: (Case 01) LoadConfig should throw an error here (not found)")
		status = false
	}

	data1, err := LoadConfig("config-test-01.json")

	if data1.IsAutoHack != false || data1.IsAutoLogin != true || data1.ContestID != 1 || data1.ProblemID != "A" || data1.Testcase.InputFile != "./data/input" || data1.Testcase.OutputFile != "./data/output" || data1.Account.Username != "user" || data1.Account.Password != "pass" {
		t.Error("Test Failed: (Case 02) LoadConfig return wrong result")
		t.Error(data1)
		status = false
	}

	if err != nil {
		t.Error(err)
		status = false
	}

	_, err = LoadConfig("config-test-02.json")

	if err == nil {
		t.Error("Test Failed: (Case 03) LoadConfig should throw an error here (format error)")
		status = false
	}
	_, err = LoadConfig("config-test-03.json")

	if err == nil {
		t.Error("Test Failed: (Case 04) LoadConfig should throw an error here (format error)")
		status = false
	}

	if status == true {
		t.Log("Package app - LoadConfig test passed")
	} else {
		t.Log("Package app - LoadConfig test failed")
	}

	clean()
}
