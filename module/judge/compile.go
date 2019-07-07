package judge

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/process"

	. "github.com/limstash/Codeforces-Hacker/common"
)

func Compile(submission Submission) (bool, error) {

	var args []string

	command := GetCompileCommand(submission)

	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		args = GetUnixCompileArgs(submission)
	} else {
		args = GetWindowsCompileArgs(submission)
	}

	cmd := exec.Command(command, args...)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return false, err
	}

	if err := cmd.Start(); err != nil {
		return false, err
	}

	startStamp := time.Now().UnixNano() / 1e6
	compileStatus := true

	go func(processHandle *os.Process) {
		pid := int32(processHandle.Pid)
		isRun, _ := process.PidExists(pid)

		for isRun == true {
			if int((time.Now().UnixNano()/1e6)-startStamp) >= 10000 {
				processHandle.Kill()
				compileStatus = false
			}

			isRun, _ = process.PidExists(pid)
		}
	}(cmd.Process)

	msg, _ := ioutil.ReadAll(stderr)

	if err := cmd.Wait(); err != nil {
		if compileStatus == false {
			return false, errors.New("Compile Timeout")
		}
		return false, errors.New(string(msg))
	}

	return true, nil
}
