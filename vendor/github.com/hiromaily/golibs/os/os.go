package os

import (
	reg "github.com/hiromaily/golibs/regexp"
	"os"
)

// GetArgs is to get args
func GetArgs(i int) string {
	return os.Args[i]
}

// AddParam is to add value as arguments
func AddParam(val string) {
	os.Args = append(os.Args, val)
}

// FindParam is to find specific args
func FindParam(key string) (bRet bool) {
	//fmt.Println(os.Args)
	bRet = false
	for _, v := range os.Args {
		if reg.CheckRegexp(`^`+key, v) {
			bRet = true
			break
		}
	}
	return
}
