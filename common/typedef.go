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

type Config struct {
	Path   string
	Server string

	ContestID   int    `json:"contest"`
	ProblemID   string `json:"problem"`
	IsAutoLogin bool   `json:"autoLogin"`
	IsAutoHack  bool   `json:"autohack"`

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
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Types               string `json:"type"`
	Phase               string `json:"phase"`
	Frozen              bool   `json:"frozen"`
	DurationSeconds     int64  `json:"durationSeconds"`
	StartTimeSeconds    int64  `json:"startTimeSeconds"`
	RelativeTimeSeconds int64  `json:"relativeTimeSeconds"`
}

type Contests struct {
	Status string    `json:"status"`
	Result []Contest `json:"result"`
}

type Problem struct {
	ContestID int      `json:"contestId"`
	Index     string   `json:"index"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Points    float32  `json:"points"`
	Rating    int      `json:"rating"`
	Tags      []string `json:"tags"`
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
