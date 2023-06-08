package operate

import (
	"context"
	"fmt"
	"note2git/e"
	"note2git/tool"
	"strings"
)

func Check(ctx context.Context, dir string) error {
	branch := tool.NowBranch()
	if err := LocalAuto(dir, branch); err != nil {
		return err
	}
	if err := RemoteHasOrigin(dir); err != nil {
		return err
	}
	if err := UpstreamToAuto(dir, branch); err != nil {
		return err
	}
	return nil
}

func LocalAuto(dir, branch string) error {
	// git checkout -b auto0
	_, stdErr, err := e.Exec(dir, "git", "checkout", "-b", branch)
	if err != nil {
		return err
	}
	if stdErr != "" {
		if strings.Contains(stdErr, "exist") {
			_, stdErr, err := e.Exec(dir, "git", "checkout", branch)
			if err != nil {
				return err
			}
			if stdErr != "" {
				return fmt.Errorf("err %v", stdErr)
			}
		}
	}
	// git branch
	o, e, err := e.Exec(dir, "git", "status")
	if err != nil {
		return err
	}
	if e != "" {
		return fmt.Errorf("err %v", e)
	}
	if !strings.Contains(o, "On branch "+branch) {
		return fmt.Errorf("创建branch失败")
	}
	return nil
}

func RemoteHasOrigin(dir string) error {
	o, e, err := e.Exec(dir, "git", "remote")
	if err != nil {
		return err
	}
	if e != "" {
		return fmt.Errorf(e)
	}
	if strings.HasPrefix(o, "origin\n") || strings.Contains(o, "\norigin\n") {
		return nil
	} else {
		return fmt.Errorf("建议使用git remote add origin xxxxx")
	}
}

func UpstreamToAuto(dir, branch string) error {
	// git push --set-upstream origin auto
	_, e, err := e.Exec(dir, "git", "push", "--set-upstream", "origin", branch)
	if err != nil {
		return err
	}
	if e == "" {
		return nil
	}
	return fmt.Errorf(e)
}
