package contest

import (
    "github.com/hytzongxuan/Codeforces-Hacker/module/http"
    "encoding/json"
    "errors"
    "net/http"
)

type Contest struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Types string `json:"type"`
    Phase string `json:"phase"`
    Frozen bool `json:"frozen"`
    DurationSeconds int64 `json:"durationSeconds"`
    StartTimeSeconds int64 `json:"startTimeSeconds"`
    RelativeTimeSeconds int64 `json:"relativeTimeSeconds"`
}

type Contests struct {
    Status string `json:"status"`
    Result []Contest `json:"result"`
}

func queryContests (cookie *[]*http.Cookie) (string, error) {
    res, err := con.HttpGet("https://codeforces.com/api/contest.list?gym=false", cookie, map[string]string{"HOST":"codeforces.com"})

    if err != nil {
        return "", err
    }
    
    return res, nil
}

func GetContests (cookie *[]*http.Cookie) ([]Contest, error) {
    data, e := queryContests(cookie)

    if(e != nil){
        return nil, e
    }

    res := Contests{}
    e = json.Unmarshal([]byte(data), &res)

    if(e != nil){
        return nil, e
    }

    if(res.Status != "OK"){
        return nil, errors.New("Codeforces Return Error Response")
    }

    return res.Result, nil
}