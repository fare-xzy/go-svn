package util

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

type Charset string

const (
	LinuxOs       = "linux"
	LinuxPrefix   = "bash"
	LinuxC        = "-c"
	WindowsOs     = "windows"
	WindowsPrefix = "cmd"
	WindowsC      = "/c"
	UTF8          = Charset("UTF-8")
	GB18030       = Charset("GB18030")
)

var (
	envOs = "windows"
)

func init() {
	envOs = runtime.GOOS
}

// Exec 一次性输出
func Exec(cmd string) string {
	var c *exec.Cmd

	if strings.EqualFold(envOs, LinuxOs) {
		c = exec.Command(LinuxPrefix, LinuxC, cmd)
	} else {
		c = exec.Command(WindowsPrefix, WindowsC, cmd)
	}
	output, _ := c.CombinedOutput()
	return ConvertByte2String(output, GB18030)
}

// RealTimeExec 实时输出
func RealTimeExec(cmd string) error {
	var c *exec.Cmd

	if strings.EqualFold(envOs, LinuxOs) {
		c = exec.Command(LinuxPrefix, LinuxC, cmd)
	} else {
		c = exec.Command(WindowsPrefix, WindowsC, cmd)
	}

	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			readBty, err := reader.ReadBytes('\n')
			if err != nil || err == io.EOF {
				return
			}
			fmt.Print(ConvertByte2String(readBty, GB18030))
		}
	}()
	err = c.Start()
	wg.Wait()
	return err
}

// 转换byte为String
func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}
