package app

import (
	"github.com/hytzongxuan/Codeforces-Hacker/module/contest"
	"encoding/json"
	"testing"
	"strconv"
	"time"
)

func Test_FindContest(t *testing.T){
	_, e := FindContest(nil)

	if e != nil && e.Error() != "[Info] Contests is an empty field" {
		t.Error("App contest failed in checking contests size")
	}

	data1_contest := `{"status":"OK","result":[{"id":1148,"name":"Codeforces Global Round 3","type":"CF","phase":"BEFORE","frozen":false,"durationSeconds":7200,"startTimeSeconds":1559399700,"relativeTimeSeconds":-3276603},{"id":1158,"name":"Codeforces Round #TBA (Div. 1)","type":"CF","phase":"BEFORE","frozen":false,"durationSeconds":7200,"startTimeSeconds":1557671700,"relativeTimeSeconds":-1548603},{"id":1159,"name":"Codeforces Round #TBA (Div. 2)","type":"CF","phase":"BEFORE","frozen":false,"durationSeconds":7200,"startTimeSeconds":1557671700,"relativeTimeSeconds":-1548603},{"id":1147,"name":"Forethought Future Cup - Final Round","type":"CF","phase":"BEFORE","frozen":false,"durationSeconds":7200,"startTimeSeconds":1556989500,"relativeTimeSeconds":-866403},{"id":1156,"name":"Educational Codeforces Round 64 (Rated for Div. 2)","type":"ICPC","phase":"BEFORE","frozen":false,"durationSeconds":7200,"startTimeSeconds":1556721300,"relativeTimeSeconds":-598203},{"id":1149,"name":"Codeforces Round #556 (Div. 1)","type":"CF","phase":"BEFORE","frozen":false,"durationSeconds":7200,"startTimeSeconds":1556548500,"relativeTimeSeconds":-425403},{"id":1150,"name":"Codeforces Round #556 (Div. 2)","type":"CF","phase":"BEFORE","frozen":false,"durationSeconds":7200,"startTimeSeconds":1556548500,"relativeTimeSeconds":-425403},{"id":1157,"name":"Codeforces Round #555 (Div. 3)","type":"ICPC","phase":"BEFORE","frozen":false,"durationSeconds":7200,"startTimeSeconds":1556289300,"relativeTimeSeconds":-166203},{"id":1152,"name":"Codeforces Round #554 (Div. 2)","type":"CF","phase":"CODING","frozen":false,"durationSeconds":7200,"startTimeSeconds":1556116500,"relativeTimeSeconds":6597},{"id":1155,"name":"Educational Codeforces Round 63 (Rated for Div. 2)","type":"ICPC","phase":"FINISHED","frozen":false,"durationSeconds":7200,"startTimeSeconds":1555943700,"relativeTimeSeconds":179397},{"id":1146,"name":"Forethought Future Cup - Elimination Round","type":"CF","phase":"FINISHED","frozen":false,"durationSeconds":9000,"startTimeSeconds":1555783500,"relativeTimeSeconds":339596},{"id":1151,"name":"Codeforces Round #553 (Div. 2)","type":"CF","phase":"FINISHED","frozen":false,"durationSeconds":7200,"startTimeSeconds":1555601700,"relativeTimeSeconds":521397},{"id":1154,"name":"Codeforces Round #552 (Div. 3)","type":"ICPC","phase":"FINISHED","frozen":false,"durationSeconds":7200,"startTimeSeconds":1555425300,"relativeTimeSeconds":697797},{"id":1153,"name":"Codeforces Round #551 (Div. 2)","type":"CF","phase":"FINISHED","frozen":false,"durationSeconds":7200,"startTimeSeconds":1555164300,"relativeTimeSeconds":958797}]}`

	res := contest.Contests{}
	e = json.Unmarshal([]byte(data1_contest), &res)

	_, e = FindContest(res.Result)

	if e == nil || (e != nil && e.Error() != "[Info] Open hacking phase finished") {
		t.Error("App contest failed in checking ontests open hacking phase")
	}

	data2_contest := "{\"status\":\"OK\",\"result\":[{\"id\":1148,\"name\":\"Educational Codeforces Round 63 (Rated for Div. 2)\",\"type\":\"B\", \"phase\": \"C\",\"frozen\":false,\"durationSeconds\":7200,\"startTimeSeconds\":" + strconv.FormatInt(time.Now().Unix() - 9000, 10) + ", \"relativeTimeSeconds\":-3276603}]}"

	res = contest.Contests{}
	e = json.Unmarshal([]byte(data2_contest), &res)

	id, e := FindContest(res.Result)

	if e != nil{
		t.Error(e)
	}else if id != 1148 {
		t.Error("App Contest failed in finding contest")
	}

	data3_contest := "{\"status\":\"OK\",\"result\":[{\"id\":1148,\"name\":\"A\",\"type\":\"B\", \"phase\": \"C\",\"frozen\":false,\"durationSeconds\":7200,\"startTimeSeconds\":" + strconv.FormatInt(time.Now().Unix() - 9000, 10) + ", \"relativeTimeSeconds\":-3276603}]}"

	res = contest.Contests{}
	e = json.Unmarshal([]byte(data3_contest), &res)

	_, e = FindContest(res.Result)

	if e == nil || (e != nil && e.Error() != "[Info] No vaild contest") {
		t.Error("App contest failed in checking vaild contest")
	}
}