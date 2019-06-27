package common

import "net/http"

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Testcase struct {
	InputFile  string `json:"inputFile"`
	OutputFile string `json:"outputFile"`
}

type Settings struct {
	Path   string
	Server string

	IsAutoLogin bool `json:"autoLogin"`
	IsAutoHack  bool `json:"autoHack"`
}

type Target struct {
	ContestID int    `json:"contest"`
	ProblemID string `json:"index"`
}

type Config struct {
	Target   Target   `json:"target"`
	Settings Settings `json:"settings"`

	Account  Account  `json:"account"`
	Testcase Testcase `json:"testcase"`
}

type Authentication struct {
	CSRF   string
	Cookie []*http.Cookie
}

type Request struct {
	URL         string
	Method      string
	Header      map[string]string
	Data        map[string]string
	NotRedirect bool

	Authentication *Authentication
}

type Response struct {
	ResponseBody   []byte
	RedirectStatus bool
}

type Contest struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	DurationSeconds  int64  `json:"durationSeconds"`
	StartTimeSeconds int64  `json:"startTimeSeconds"`
}

type Contests struct {
	Status string    `json:"status"`
	Result []Contest `json:"result"`
}

type Problem struct {
	ContestID int    `json:"contestId"`
	Index     string `json:"index"`
	Name      string `json:"name"`
}

type ProblemList struct {
	Problems []Problem `json:"problems"`
}

type Problems struct {
	Status string      `json:"status"`
	Result ProblemList `json:"result"`
}

type Submission struct {
	SubmissionID int
	Language     string
	Code         string
}
