package judge

import "testing"

func Test_splitVersion(t *testing.T) {
	status := true

	first, second, third := splitVersion("3.4.5")

	if first != 3 || second != 4 || third != 5 {
		t.Error("Test Failed: (Case 01) splitVersion return wrong version info")
		t.Error(first, second, third)
		status = false
	}

	first, second, third = splitVersion("12.134.156")

	if first != 12 || second != 134 || third != 156 {
		t.Error("Test Failed: (Case 02) splitVersion return wrong version info")
		t.Error(first, second, third)
		status = false
	}

	if status == true {
		t.Log("Package judge - splitVersion test passed")
	} else {
		t.Log("Package judge - splitVersion test failed")
	}
}
