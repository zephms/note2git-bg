package tool

import "testing"

func TestNowBranch(t *testing.T) {
	b := NowBranch()
	t.Log(b)
}
