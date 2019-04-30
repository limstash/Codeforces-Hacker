package judge

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Diff(SubmissionID int) bool {
	SubmissionPath := getPath() + "/src/" + strconv.Itoa(SubmissionID)
	status := true

	outputFile, err := os.Open(SubmissionPath + "/data.out")

	if err != nil {
		return false
	}

	standFile, err := os.Open(SubmissionPath + "/data.ans")

	if err != nil {
		return false
	}

	outputBuff := bufio.NewReader(outputFile)
	standBuff := bufio.NewReader(standFile)

	end1 := false
	end2 := false

	for {
		var output string
		var stand string

		if end1 == false {
			a, _, c := outputBuff.ReadLine()

			if c == io.EOF {
				end1 = true
			}

			output = string(a)

			if end2 == true && output != "" && output != "\r" && output != "\n" && output != "\r\n" {
				status = false
			}
		}

		if end2 == false {
			b, _, d := standBuff.ReadLine()

			if d == io.EOF {
				end2 = true
			}

			stand = string(b)

			if end1 == true && stand != "" && stand != "\r" && stand != "\n" && stand != "\r\n" {
				status = false
			}
		}

		if end1 == true && end2 == true {
			break
		}

		for i := 0; i < min(len(output), len(stand)); i++ {
			if output[i] != stand[i] {
				status = false
				break
			}
		}

		if len(output) < len(stand) {
			for i := len(output); i < len(stand); i++ {
				if stand[i] != ' ' && stand[i] != '\r' && stand[i] != '\n' {
					status = false
					break
				}
			}
		} else {
			for i := len(stand); i < len(output); i++ {
				if output[i] != ' ' && output[i] != '\r' && output[i] != '\n' {
					status = false
					break
				}
			}
		}

		if status == false {
			return status
		}
	}

	return status
}
