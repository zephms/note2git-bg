package operate

import (
	"context"
	"fmt"
	ex "note2git/e"
	"time"
)

// todo: context link
func action(ctx context.Context, dir string, branch string) error {
	_, e, err := ex.Exec("dir", "git", "fetch", "--all")
	if err != nil || e != "" {
		return fmt.Errorf("fetch error, stdErr: %v, execErr: %v", e, err)
	}
	// git diff HEAD; 比较当前工作区和上次提交
	o, e, err := ex.Exec(dir, "git", "diff", "HEAD")
	if err != nil || e != "" {
		return fmt.Errorf("diff error, stdErr: %v, execErr: %v", e, err)
	}
	if o == "" {
		fmt.Println("当前无提交")
		return nil
	}
	ex.Exec(dir, "git", "add", ".")
	if err != nil || e != "" {
		return fmt.Errorf("git add error, stdErr: %v, execErr: %v", e, err)
	}
	ex.Exec(dir, "git", "commit", "-m", fmt.Sprintf("autosaved when %s", time.Now().Format("2006-01-02 15-04-05")))
	if err != nil || e != "" {
		return fmt.Errorf("git commit error, stdErr: %v, execErr: %v", e, err)
	}
	ex.Exec(dir, "git", "push", "origin")
	if err != nil || e != "" {
		return fmt.Errorf("push error, stdErr: %v, execErr: %v", e, err)
	}
	return nil
}
