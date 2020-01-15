# JS Tooling

The goal of this application is to allow you to execute .js files with special functions available in Javascript but implemented and executed in Go.


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
- v = ReadFile(file)
- v = ReadLn()
- v = Download(url, filename)
- v = CreateDir(path)
- v = OverwriteFileWithString(filename, varName)
- v = CopyFile(fileFrom, fileTo)
- v = RemoveFile(filename)

Customizations
- Ytt(template, overlay, outputDir)

Logging
- announceln(text)
- promptln(text)
- infoln(text)
- errorln(text)
- successln(text)

Other
- Consider putting the functions in an object like `jst`.
- Add logging help-functions

## Building

Windows: Clone the code and `go build -o jst.exe bin/jst.go`


Linux: Clone the code and `go build -o jst bin/jst.go`

## Download and run
- Visit the Releases page and download the binary for Linux or Windows.
- Ensure it has execute permissions
- Run `jst --file something.js` 