package util

import (
	"fmt"
	"testing"
)

const COMMOND = "ping www.baidu.com"

func TestExec(t *testing.T) {
	fmt.Println(Exec(COMMOND))
}

func TestRealTimeExec(t *testing.T) {
	RealTimeExec(COMMOND)
}
