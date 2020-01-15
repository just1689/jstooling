package features

import (
	"github.com/just1689/jstooling/model"
	"github.com/just1689/jstooling/plugins"
	"github.com/just1689/jstooling/util"
	"github.com/robertkrimen/otto"
)

func StartVM() {
	model.VM = otto.New()

}

func LoadFunctions() {

	//IO
	model.VM.Set("FileToVar", util.Wrapper(plugins.FileToVar))
	model.VM.Set("ReadUserInputToVar", util.Wrapper(plugins.ReadUserInputToVar))
	model.VM.Set("DownloadFile", util.Wrapper(plugins.DownloadFile))
	model.VM.Set("GrantExecute", util.Wrapper(plugins.GrantExecute))
	model.VM.Set("CreateDir", util.Wrapper(plugins.CreateDir))
	model.VM.Set("OverwriteFileWithString", util.Wrapper(plugins.OverwriteFileWithString))
	model.VM.Set("CopyFile", util.Wrapper(plugins.CopyFile))
	model.VM.Set("RemoveFile", util.Wrapper(plugins.RemoveFile))

	//Color
	model.VM.Set("Magenta", util.Wrapper(plugins.Magenta))
	model.VM.Set("Yellow", util.Wrapper(plugins.Yellow))
	model.VM.Set("Cyan", util.Wrapper(plugins.Cyan))
	model.VM.Set("Red", util.Wrapper(plugins.Red))
	model.VM.Set("Blue", util.Wrapper(plugins.Blue))
	model.VM.Set("Green", util.Wrapper(plugins.Green))

	//YTT
	model.VM.Set("Ytt", util.Wrapper(plugins.Ytt))

}
