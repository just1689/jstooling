package model

import "github.com/robertkrimen/otto"

var VM *otto.Otto

type StandardFunction func(in []string) (msg string, err error)
