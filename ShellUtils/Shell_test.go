package ShellUtils

import (
	"fmt"
	"testing"
)

func TestGetOutFromStdout(t *testing.T) {
	t.Run("测试windeos的shell交互:", test)
}

func test(t *testing.T) {
	out := GetOutFromStdout("dir")
	fmt.Println(out)
}
