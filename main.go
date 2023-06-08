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

	// fmt.Println("ok")
	// ob, err := cmd.Output()
	// fmt.Println(ob)
	// fmt.Println(err.Error())

	// r, err := git.PlainOpen("D:/Acache/slinttest")
	// panic1(err)
	// headRef, err := r.Head()
	// ref := plumbing.NewHashReference("refs/heads/my-branch", headRef.Hash())
	// err = r.Storer.SetReference(ref)
	// panic1(err)
	// fmt.Println(branch.Name)

}
