package internal

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var (
	testPath = "http://192.168.2.1/develop/XSS_V2/3Workspace"
	outPath  = "D://1"
)

func TestDoList(t *testing.T) {
	list := DoList(testPath)
	split := strings.Split(list, "\r\n")
	for _, child := range split {
		if len(child) > 0 && strings.Contains(child, "/") {
			fmt.Println(testPath + "/" + child)
		}
	}
}
func TestDoCd(t *testing.T) {
	getwd, _ := os.Getwd()
	fmt.Println(getwd)
	os.Chdir("D://")
	getwd, _ = os.Getwd()
	fmt.Println(getwd)
}

func TestDoCheckOut(t *testing.T) {
	DoCheckOut(testPath, outPath)
}
