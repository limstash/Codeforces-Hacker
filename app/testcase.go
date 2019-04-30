package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

func openLinux(address string) error {
	subProcess := exec.Command("xdg-open", address)

	if err := subProcess.Start(); err != nil {
		return err
	}

	subProcess.Wait()

	return nil
}

func openWindows(address string) error {
	subProcess := exec.Command("notepad.exe", address)

	if err := subProcess.Start(); err != nil {
		return err
	}

	subProcess.Wait()

	return nil
}

func openMac(address string) error {
	subProcess := exec.Command("open", address)

	if err := subProcess.Start(); err != nil {
		return err
	}

	subProcess.Wait()

	return nil
}

func createFile(address string) error {
	f, err := os.Create(address)
	defer f.Close()

	if err != nil {
		return err
	}

	return nil
}

func saveInput() error {
	confirm := false

	for confirm == false {
		fmt.Println("")
		fmt.Println("[Info] Please enter your input data")

		var e error

		if runtime.GOOS == "linux" {
			e = openLinux(getPath() + "/src/data.in")
		} else if runtime.GOOS == "darwin" {
			e = openMac(getPath() + "/src/data.in")
		} else {
			e = openWindows(getPath() + "/src/data.in")
		}

		if e != nil {
			return e
		}

		confirm, e = GetUserResponseYN("", "[Info] type yes to continue or no to retry: ", "")

		if e != nil {
			return e
		}
	}

	return nil
}

func saveOutput() error {
	confirm := false

	for confirm == false {
		fmt.Println("")
		fmt.Println("[Info] Please enter your output data")

		var e error

		if runtime.GOOS == "linux" {
			e = openLinux(getPath() + "/src/data.ans")
		} else if runtime.GOOS == "darwin" {
			e = openMac(getPath() + "/src/data.ans")
		} else {
			e = openWindows(getPath() + "/src/data.ans")
		}

		if e != nil {
			return e
		}

		confirm, e = GetUserResponseYN("", "[Info] type yes to continue or no to retry: ", "")

		if e != nil {
			return e
		}
	}

	return nil
}

func printInput() error {
	b, err := ioutil.ReadFile(getPath() + "/src/data.in")

	if err != nil {
		return err
	}
	str := string(b)
	fmt.Println(str)

	return nil
}

func printOutput() error {
	b, err := ioutil.ReadFile(getPath() + "/src/data.ans")

	if err != nil {
		return err
	}

	str := string(b)
	fmt.Println(str)

	return nil
}

func SaveData() error {
	e := os.MkdirAll(getPath()+"/src", 0777)

	if e != nil {
		return e
	}

	e = createFile(getPath() + "/src/data.in")

	if e != nil {
		return e
	}

	e = createFile(getPath() + "/src/data.ans")

	if e != nil {
		return e
	}

	choose := false

	for choose == false {
		e := saveInput()

		if e != nil {
			return e
		}

		e = saveOutput()

		if e != nil {
			return e
		}

		fmt.Println("")
		fmt.Println("--- data.in ---")
		e = printInput()

		if e != nil {
			return e
		}

		fmt.Println("--- data.ans ---")
		e = printOutput()

		if e != nil {
			return e
		}

		chooseInput, e := GetUserResponseYN("", "[Info] type yes to continue or no to retry: ", "")

		if e != nil {
			return e
		}

		choose = chooseInput
	}

	return nil
}
