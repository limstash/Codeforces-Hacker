package judge

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"strconv"
)

func RunUnixBin(SubmissionID int, SubmissionPath string) bool {
	subProcess := exec.Command(SubmissionPath + "/main")

	stdin, err := subProcess.StdinPipe()

	if err != nil {
		return false
	}

	stdout, err := os.OpenFile(SubmissionPath+"/data.out", os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		return false
	}

	subProcess.Stdout = stdout

	stderr, err := os.OpenFile(SubmissionPath+"/data.err", os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		return false
	}

	subProcess.Stderr = stderr

	if err = subProcess.Start(); err != nil {
		return false
	}

	fi, err := os.Open(SubmissionPath + "/data.in")

	if err != nil {
		return false
	}

	inputBuff := bufio.NewReader(fi)

	for {
		a, _, c := inputBuff.ReadLine()
		if c == io.EOF {
			stdin.Close()
			break
		}
		io.WriteString(stdin, string(a)+"\n")
	}

	subProcess.Wait()

	stdout.Close()
	fi.Close()

	return true
}

func RunPython2(SubmissionID int, SubmissionPath string) bool {
	subProcess := exec.Command("python2", SubmissionPath+"/main.py")

	stdin, err := subProcess.StdinPipe()

	if err != nil {
		return false
	}

	stdout, err := os.OpenFile(SubmissionPath+"/data.out", os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		return false
	}

	subProcess.Stdout = stdout

	stderr, err := os.OpenFile(SubmissionPath+"/data.err", os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		return false
	}

	subProcess.Stderr = stderr

	if err = subProcess.Start(); err != nil {
		return false
	}

	fi, err := os.Open(SubmissionPath + "/data.in")

	if err != nil {
		return false
	}

	inputBuff := bufio.NewReader(fi)

	for {
		a, _, c := inputBuff.ReadLine()
		if c == io.EOF {
			stdin.Close()
			break
		}
		io.WriteString(stdin, string(a)+"\n")
	}

	subProcess.Wait()

	stdout.Close()
	fi.Close()

	return true
}

func RunPython3(SubmissionID int, SubmissionPath string) bool {
	subProcess := exec.Command("python3", SubmissionPath+"/main.py")

	stdin, err := subProcess.StdinPipe()

	if err != nil {
		return false
	}

	stdout, err := os.OpenFile(SubmissionPath+"/data.out", os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		return false
	}

	subProcess.Stdout = stdout

	stderr, err := os.OpenFile(SubmissionPath+"/data.err", os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		return false
	}

	subProcess.Stderr = stderr

	if err = subProcess.Start(); err != nil {
		return false
	}

	fi, err := os.Open(SubmissionPath + "/data.in")

	if err != nil {
		return false
	}

	inputBuff := bufio.NewReader(fi)

	for {
		a, _, c := inputBuff.ReadLine()
		if c == io.EOF {
			break
		}
		io.WriteString(stdin, string(a)+"\n")
	}

	subProcess.Wait()

	stdout.Close()
	fi.Close()

	return true
}

func RunUnixCode(SubmissionID int, lang string) bool {
	SubmissionPath := getPath() + "/src/" + strconv.Itoa(SubmissionID)

	result := true

	switch lang {
	case "GNU C11":
		result = RunUnixBin(SubmissionID, SubmissionPath)
	case "GNU C++11":
		result = RunUnixBin(SubmissionID, SubmissionPath)
	case "GNU C++14":
		result = RunUnixBin(SubmissionID, SubmissionPath)
	case "GNU C++17":
		result = RunUnixBin(SubmissionID, SubmissionPath)
	case "Go":
		result = RunUnixBin(SubmissionID, SubmissionPath)
	case "Python 2":
		result = RunPython2(SubmissionID, SubmissionPath)
	case "Python 3":
		result = RunPython3(SubmissionID, SubmissionPath)
	default:
		result = true
	}

	return result
}
