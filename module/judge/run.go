package judge

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/process"

	. "github.com/limstash/Codeforces-Hacker/common"
)

func Runcode(submission Submission, problem Problem) (bool, int, int) {

	var command string
	var args []string

	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		command = GetUnixRunCommand(submission)
	} else {
		command = GetWindowsRunCommand(submission)
	}

	args = GetRunArgs(submission)

	subProcess := exec.Command(command, args...)

	stdin, err := subProcess.StdinPipe()

	if err != nil {
		return false, 0, 0
	}

	stdout, err := os.OpenFile(submission.Path+"/data.out", os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		return false, 0, 0
	}

	subProcess.Stdout = stdout

	stderr, err := os.OpenFile(submission.Path+"/data.err", os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		return false, 0, 0
	}

	subProcess.Stderr = stderr

	startStamp := time.Now().UnixNano() / 1e6
	if err = subProcess.Start(); err != nil {
		return false, 0, 0
	}

	fi, err := os.Open(submission.Path + "/data.in")

	if err != nil {
		return false, 0, 0
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

	var usedmemory uint64
	usedmemory = 0

	memorylimitInt := problem.Memorylimit
	memorylimitInt = memorylimitInt * 1024

	timelimitInt := problem.Timelimit

	go func(processHandle *os.Process) {
		pid := int32(processHandle.Pid)

		isRun, _ := process.PidExists(pid)
		processhandle, _ := process.NewProcess(pid)

		for isRun == true {
			meminfo, e := processhandle.MemoryInfo()

			if e == nil && meminfo.RSS > usedmemory {
				usedmemory = meminfo.RSS
			}

			if int(usedmemory) > memorylimitInt {
				processHandle.Kill()
			}

			if int((time.Now().UnixNano()/1e6)-startStamp) >= timelimitInt+10 {
				processHandle.Kill()
			}

			isRun, _ = process.PidExists(pid)
		}
	}(subProcess.Process)

	err = subProcess.Wait()
	endStamp := time.Now().UnixNano() / 1e6

	if err != nil {
		return false, int(endStamp - startStamp), int(usedmemory)
	}

	stdout.Close()
	fi.Close()

	return true, int(endStamp - startStamp), int(usedmemory)
}
