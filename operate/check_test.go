package operate

import "testing"

var testDir = "D:/Awork/model_management_service"

// var testDir = "D:/Acache/note2git/tmp"

func TestCheckLocol(t *testing.T) {
	err := LocalAuto(testDir, "test0")
	t.Log(err)
}

func TestCheckRemote(t *testing.T) {
	err := RemoteHasOrigin(testDir)
	t.Log(err)
}

func TestAutoToAuto(t *testing.T) {
	err := UpstreamToAuto(testDir, "auto0")
	t.Log(err)
}
