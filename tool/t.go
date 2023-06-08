package tool

import (
	"fmt"
	"time"
)

func NowBranch() string {
	t := time.Now().Format("0602")
	return fmt.Sprintf("autosave%s", t)
}
