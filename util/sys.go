package util

import (
	"github.com/fatih/color"
	"os"
)

func CheckForNoWork(all []*string) {
	for _, f := range all {
		if *f != "" {
			return
		}
	}
	color.Red("%s", "nothing to do")
	os.Exit(0)

}
