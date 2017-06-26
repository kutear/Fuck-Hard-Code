package utils

import "testing"

func TestContains(t *testing.T) {
	strArr := []string{"hello", "world", "hi"}
	for _, v := range strArr {
		if !Contains(strArr, v) {
			t.Errorf("%s is in strArr", v)
		}
	}
}
