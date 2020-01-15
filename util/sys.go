package util

import (
	"github.com/fatih/color"
	"os"
)

func CheckForNoWork(f *string) {
	if *f == "" {
		color.Red("%s", "nothing to do")
		os.Exit(0)
	}
}
