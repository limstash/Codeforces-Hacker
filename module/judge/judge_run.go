package judge

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/process"

	. "github.com/hytzongxuan/Codeforces-Hacker/common"
)

func Runcode(submission Submission, problem Problem) (bool, int, int) {

	var command string
	var args []string

	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		command = GetUnixRunCommand(submission)
		args = GetUnixRunArgs(submission)
	} else {
		command = GetWindowsRunCommand(submission)
		args = GetWindowsRunArgs(submission)
	}

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

	go func(pid int32) {
		isRun, _ := process.PidExists(pid)
		processhandle, _ := process.NewProcess(pid)

		for isRun == true {
			meminfo, e := processhandle.MemoryInfo()

			if e == nil && meminfo.RSS > usedmemory {
				usedmemory = meminfo.RSS
			}

			if int(usedmemory) > memorylimitInt {
				syscall.Kill(int(pid), syscall.SIGKILL)
			}

			if int((time.Now().UnixNano()/1e6)-startStamp) >= timelimitInt+10 {
				syscall.Kill(int(pid), syscall.SIGKILL)
			}

			isRun, _ = process.PidExists(pid)
		}
	}(int32(subProcess.Process.Pid))

	err = subProcess.Wait()
	endStamp := time.Now().UnixNano() / 1e6

	if err != nil {
		return false, int(endStamp - startStamp), int(usedmemory)
	}

	stdout.Close()
	fi.Close()

	return true, int(endStamp - startStamp), int(usedmemory)
}
