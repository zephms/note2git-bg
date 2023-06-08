package main

import (
	"fmt"
	"os/exec"
)

func run(cmd *exec.Cmd) (string, string, error) {
	o, e := cmd.Output()
	if e == nil {
		return string(o), "", cmd.Err
	}
	return string(o), e.Error(), cmd.Err
}

func main() {
	fmt.Println("asdf")

	cmd := exec.Command("git", "branch")
	cmd.Dir = "D:/Acache/slinttest"
	stdOut, stdErr, err := run(cmd)

	fmt.Println(stdOut, stdErr, err)

}
