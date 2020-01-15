package plugins

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/just1689/jstooling/model"
	"github.com/robertkrimen/otto"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

// in[0] = filename
// in[1] = output var name
var FileToVar model.StandardFunction = func(in []string) (msg string, err error) {
	if len(in) != 2 {
		err = errors.New("not enough parameters FileToVar")
		return
	}
	var b []byte
	b, err = ioutil.ReadFile(in[0])
	if err != nil {
		return
	}
	err = model.VM.Set(in[1], string(b))
	return
}

// in[0] = ouput var name
var ReadUserInputToVar model.StandardFunction = func(in []string) (msg string, err error) {
	if len(in) != 1 {
		err = errors.New("not enough parameters ReadUserLineToVar")
		return
	}
	var text string
	reader := bufio.NewReader(os.Stdin)
	text, err = reader.ReadString('\n')
	if err != nil {
		return
	}
	err = model.VM.Set(in[0], text)
	return
}

// in[0] = remote to download
// in[1] = ouput var name
var DownloadFile model.StandardFunction = func(in []string) (msg string, err error) {
	if len(in) != 2 {
		err = errors.New("not enough parameters DownloadFile")
		return
	}
	resp, err := http.Get(in[0])
	if err != nil {
		return
	}
	defer resp.Body.Close()
	out, err := os.Create(in[1])
	if err != nil {
		return
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return
}

// in[0] = the file
// in[1] = use sudo true or false
var GrantExecute model.StandardFunction = func(in []string) (msg string, err error) {
	if len(in) != 2 {
		err = errors.New("not enough parameters GrantExecute")
		return
	}
	app := "chmod"
	args := []string{"+x", in[0]}
	cmd := exec.Command(app, args[0], args[1])
	var b []byte
	b, err = cmd.Output()
	if err != nil {
		return
	}
	msg = string(b)
	return
}

// in[0] the path of the directory to create
var CreateDir model.StandardFunction = func(in []string) (msg string, err error) {
	if len(in) != 1 {
		err = errors.New("not enough parameters CreateDir")
		return
	}
	err = os.MkdirAll(in[0], os.ModePerm)
	return
}

// in[0] the file path
// in[1] the field to use to overwrite the file with
var OverwriteFileWithString model.StandardFunction = func(in []string) (msg string, err error) {
	if len(in) != 2 {
		err = errors.New("not enough parameters OverwriteFileWithString")
		return
	}

	var v otto.Value
	v, err = model.VM.Get(in[1])
	if err != nil {
		return
	}
	err = ioutil.WriteFile(in[0], []byte(v.String()), 0644)
	return
}

// in[0] from file
// in[1] to file
var CopyFile model.StandardFunction = func(in []string) (msg string, err error) {
	if len(in) != 2 {
		err = errors.New("not enough parameters CopyFile")
		return
	}
	input, err := ioutil.ReadFile(in[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(in[1], input, 0644)
	if err != nil {
		fmt.Println("Error creating", in[1])
		fmt.Println(err)
		return
	}
	return
}

// in[0] the file
var RemoveFile model.StandardFunction = func(in []string) (msg string, err error) {
	if len(in) != 1 {
		err = errors.New("not enough parameters RemoveFile")
		return
	}
	err = os.Remove(in[0])
	return
}
