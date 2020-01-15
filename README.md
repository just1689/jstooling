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
- result = FileExists(filename)

Color terminal logging
- Magenta(text)
- Yellow(text)
- Cyan(text)
- Red(text)
- Blue(text)
- Green(text)


## Planned features

Core features
- panic(msg)
- Touch(filename)
- FileContains(filename, needle)
- Rename the methods to something more conventional
- RunForOS(windows, linux)

Customizations
- Ytt(tempalte, overlay, outputDir)

Logging
- announceln(text)
- promptln(text)
- infoln(text)
- errorln(text)
- successln(text)


## Building

Windows: Clone the code and `go build -o jst.exe bin/jst.go`
Linux: Clone the code and `go build -o jst bin/jst.go`

## Download and run
- Visit the Releases page and download the binary for Linux or Windows.
- Ensure it has execute permissions
- Run `jst --file something.js` 