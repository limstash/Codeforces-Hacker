package app

import (
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/hytzongxuan/Codeforces-Hacker/common"
	"github.com/hytzongxuan/Codeforces-Hacker/module/problem"
	"github.com/hytzongxuan/Codeforces-Hacker/module/token"
)

var auth = Authentication{}

func fetchConfig(configFilePath string) Config {
	Config, err := LoadConfig(configFilePath)

	if err != nil {
		log(1, err.Error())
	}

	log(3, "Fetch config.json success")
	return Config
}

func fetchCSRF(auth *Authentication, server string) {
	err := token.GetCSRF(auth, server)

	if err != nil {
		log(1, err.Error())
	}

	log(3, "Fetch CSRF token success")
}

func login(config Config, auth *Authentication) {
	Login(config, auth)

	if config.IsAutoLogin {
		log(3, "Login success")
	}
}

func checkContest(config Config, auth *Authentication) Contest {
	contests, err := problem.GetContest(auth, config.Server)

	if err != nil {
		log(1, err.Error())
	}

	contest, err := FindContest(contests.Result, config.ContestID)

	if err != nil {
		log(1, err.Error())
	}

	return contest
}

func checkProblem(config Config, contest Contest, auth *Authentication) Problem {
	problems, err := problem.GetProblem(auth, config.Server)

	if err != nil {
		log(1, err.Error())
	}

	problem, err := FindProblem(problems.Result.Problems, contest, config.ProblemID)

	if err != nil {
		log(1, err.Error())
	}

	return problem
}
func getPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	rst := filepath.Dir(path)
	return rst
}

func Load(configFilePath string, remoteServerURL string) {
	log(3, "Codeforces Hacker Starting...")

	config := fetchConfig(configFilePath)
	config.Path = getPath()
	config.Server = remoteServerURL

	fetchCSRF(&auth, remoteServerURL)

	login(config, &auth)

	contest := checkContest(config, &auth)
	problem := checkProblem(config, contest, &auth)

	run(contest, problem, config, &auth)
}
