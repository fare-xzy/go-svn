package internal

import (
	"changeme/internal/util"
	"fmt"
	"os"
)

func DoList(path string) string {
	comm := fmt.Sprintf(SvnList, path)
	return util.Exec(comm)
}

func DoCheckOut(path, outPath string) string {
	param := fmt.Sprintf(Depth, immediates)
	comm := fmt.Sprintf(Checkout, param, path)
	fmt.Println(comm)
	os.Chdir(outPath)
	return util.Exec(comm)
}
