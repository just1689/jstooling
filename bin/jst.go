package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/just1689/jstooling/features"
	"github.com/just1689/jstooling/model"
	"github.com/just1689/jstooling/util"
	"io/ioutil"
	"net/http"
)

var f = flag.String("file", "", "file to load")
var u = flag.String("url", "", "url to download an run to load")

func main() {
	flag.Parse()
	util.CheckForNoWork([]*string{f, u})

	features.StartVM()
	features.LoadFunctions()

	handleFile()
	handleURL()

}

func handleFile() {
	if *f == "" {
		return
	}
	b, err := ioutil.ReadFile(*f)
	if err != nil {
		color.Red("%s", fmt.Sprint("could not read file", *f))
		return
	}
	model.VM.Run(string(b))
}

func handleURL() {
	if *u == "" {
		return
	}

	// Get the data
	resp, err := http.Get(*u)
	if err != nil {
		color.Red("%s", fmt.Sprint("could not read url", *u))
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		color.Red("%s", fmt.Sprint("could not read body of url", *u))
		return
	}
	model.VM.Run(string(b))
}
