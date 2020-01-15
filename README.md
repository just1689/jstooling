# JS Tooling

The goal of this application is to allow you to execute .js files with special functions in Javascript because executed in Go.


## Support built-in functions

IO functions
- FileToVar(filename, varName)
- ReadUserInputToVar(varName)
- DownloadFile(url, filename)
- GrantExecute(filename)
- CreateDir(path)
- OverwriteFileWithString(filename, varName)
- CopyFile(fileFrom, fileTo)
- RemoveFile(filename)

Color terminal logging
- Magenta(text)
- Yellow(text)
- Cyan(text)
- Red(text)
- Blue(text)
- Green(text)


## Planned features
- Ytt(tempalte, overlay, outputDir)
- Touch(filename)
- FileContains(filename, needle)
- FileExists

## Building

Clone the code and `go build app.go`

## Download and run
- Visit the Releases page and download the binary for Linux or Windows.
- Ensure it has execute permissions
- Run `jst --file something.js` 