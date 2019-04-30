package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserResponseYN(before string, output string, after string) (bool, error) {
	fmt.Printf(before)
	fmt.Printf(output)

	var choose string

	myReader := bufio.NewReader(nil)
	myReader.Reset(os.Stdin)
	content, e := myReader.ReadString('\n')

	if e != nil {
		return false, e
	}

	fields := strings.Fields(content)

	if fields == nil || len(fields) == 0 {
		choose = ""
	} else {
		choose = fields[0]
	}

	for e != nil || (choose != "yes" && choose != "no") {
		fmt.Printf(output)

		myReader = bufio.NewReader(nil)
		myReader.Reset(os.Stdin)
		content, e = myReader.ReadString('\n')

		if e != nil {
			return false, e
		}

		fields = strings.Fields(content)

		if fields == nil || len(fields) == 0 {
			choose = ""
		} else {
			choose = fields[0]
		}
	}

	fmt.Printf(after)

	if choose == "yes" {
		return true, e
	}

	return false, e
}

func GetUserResponseString(before string, output string, after string) (string, error) {
	fmt.Printf(before)
	fmt.Printf(output)

	var data string

	myReader := bufio.NewReader(nil)
	myReader.Reset(os.Stdin)
	content, e := myReader.ReadString('\n')

	fields := strings.Fields(content)

	if fields == nil || len(fields) == 0 {
		data = ""
	} else {
		data = fields[0]
	}

	for e != nil || data == "" {
		fmt.Printf(output)

		myReader = bufio.NewReader(nil)
		myReader.Reset(os.Stdin)
		content, e = myReader.ReadString('\n')

		if e != nil {
			return "", e
		}

		fields = strings.Fields(content)

		if fields == nil || len(fields) == 0 {
			data = ""
		} else {
			data = fields[0]
		}
	}

	fmt.Printf(after)
	return data, e
}
