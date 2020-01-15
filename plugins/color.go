package plugins

import (
	"github.com/fatih/color"
	"github.com/just1689/jstooling/model"
)

var Cyan model.StandardFunction = func(in []string) (msg string, err error) {
	color.Cyan("%s", in[0])
	return
}
var Yellow model.StandardFunction = func(in []string) (msg string, err error) {
	color.Yellow("%s", in[0])
	return
}
var Red model.StandardFunction = func(in []string) (msg string, err error) {
	color.Red("%s", in[0])
	return
}
var Blue model.StandardFunction = func(in []string) (msg string, err error) {
	color.Blue("%s", in[0])
	return
}
var Magenta model.StandardFunction = func(in []string) (msg string, err error) {
	color.Magenta("%s", in[0])
	return
}
var Green model.StandardFunction = func(in []string) (msg string, err error) {
	color.Green("%s", in[0])
	return
}
