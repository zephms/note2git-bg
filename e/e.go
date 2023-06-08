package e

import "os/exec"

func Exec(dir string, name string, args ...string) (string, string, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	o, e := cmd.Output()
	if e == nil {
		return string(o), "", cmd.Err
	}
	return string(o), e.Error(), cmd.Err
}
