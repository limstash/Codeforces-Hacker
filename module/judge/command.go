package judge

import (
	. "github.com/hytzongxuan/Codeforces-Hacker/common"
)

func GetCompileCommand(submission Submission) string {
	switch submission.Language {

	case "GNU C11":
		return "gcc"
	case "GNU C++11":
		return "g++"
	case "GNU C++14":
		return "g++"
	case "GNU C++17":
		return "g++"
	case "Go":
		return "go"
	default:
		return ""
	}
}

func GetUnixCompileArgs(submission Submission) []string {
	path := submission.Path
	src := path + "/main"

	switch submission.Language {

	case "GNU C11":
		return []string{src + ".c", "-o", src, "-lm", "-std=c11", "-DONLINE_JUDGE"}
	case "GNU C++11":
		return []string{src + ".cpp", "-o", src, "-lm", "-std=c++11", "-DONLINE_JUDGE"}
	case "GNU C++14":
		return []string{src + ".cpp", "-o", src, "-lm", "-std=c++14", "-DONLINE_JUDGE"}
	case "GNU C++17":
		return []string{src + ".cpp", "-o", src, "-lm", "-std=c++17", "-DONLINE_JUDGE"}
	case "Go":
		return []string{"-o", src, src + ".go"}
	default:
		return []string{}
	}
}

func GetWindowsCompileArgs(submission Submission) []string {
	path := submission.Path
	src := path + "/main"
	target := src + ".exe"

	switch submission.Language {

	case "GNU C11":
		return []string{src + ".c", "-o", target, "-lm", "-std=c11", "-DONLINE_JUDGE", "-Wl,--stack=536870912"}
	case "GNU C++11":
		return []string{src + ".cpp", "-o", target, "-lm", "-std=c++11", "-DONLINE_JUDGE", "-Wl,--stack=536870912"}
	case "GNU C++14":
		return []string{src + ".cpp", "-o", target, "-lm", "-std=c++14", "-DONLINE_JUDGE", "-Wl,--stack=536870912"}
	case "GNU C++17":
		return []string{src + ".cpp", "-o", target, "-lm", "-std=c++17", "-DONLINE_JUDGE", "-Wl,--stack=536870912"}
	case "Go":
		return []string{"-o", target, src + ".go"}
	default:
		return []string{}
	}
}

func GetUnixRunCommand(submission Submission) string {
	path := submission.Path
	target := path + "/main"

	switch submission.Language {

	case "GNU C11":
		return target
	case "GNU C++11":
		return target
	case "GNU C++14":
		return target
	case "GNU C++17":
		return target
	case "Go":
		return target
	case "Python 2":
		return "python2"
	case "Python 3":
		return "python3"
	default:
		return ""
	}
}

func GetRunArgs(submission Submission) []string {
	path := submission.Path
	src := path + "/main"

	switch submission.Language {
	case "GNU C11":
		return []string{}
	case "GNU C++11":
		return []string{}
	case "GNU C++14":
		return []string{}
	case "GNU C++17":
		return []string{}
	case "Go":
		return []string{}
	case "Python 2":
		return []string{src + ".py"}
	case "Python 3":
		return []string{src + ".py"}
	default:
		return []string{}
	}
}

func GetWindowsRunCommand(submission Submission) string {
	path := submission.Path
	target := path + "/main.exe"

	switch submission.Language {

	case "GNU C11":
		return target
	case "GNU C++11":
		return target
	case "GNU C++14":
		return target
	case "GNU C++17":
		return target
	case "Go":
		return target
	case "Python 2":
		return "python2"
	case "Python 3":
		return "python3"
	default:
		return ""
	}
}
