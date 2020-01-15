package util

import (
	"github.com/fatih/color"
	"github.com/just1689/jstooling/model"
	"github.com/robertkrimen/otto"
)

func ArgsToStrings(list []otto.Value) []string {
	result := make([]string, len(list))
	for i, item := range list {
		result[i] = item.String()
	}
	return result
}

func Wrapper(f model.StandardFunction) func(call otto.FunctionCall) otto.Value {
	return func(call otto.FunctionCall) otto.Value {
		s, err := f(ArgsToStrings(call.ArgumentList))
		if err != nil {
			color.Red("%s", err.Error())
			r, _ := otto.ToValue("")
			return r
		}
		result, err := otto.ToValue(s)
		if err != nil {
			color.Red("%s", err.Error())
		}
		return result
	}
}
