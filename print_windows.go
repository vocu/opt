package opt

// https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences
// https://github.com/hashicorp/vault/blob/c44f1c9817955d4c7cd5822a19fb492e1c2d0c54/vendor/github.com/bgentry/speakeasy/speakeasy_windows.go
// https://github.com/konsorten/go-windows-terminal-sequences/blob/master/sequences.go
// https://github.com/nine-lives-later/go-windows-terminal-sequences/blob/master/sequences.go

import "golang.org/x/sys/windows"

func init() {
	const ENABLE_VIRTUAL_TERMINAL_PROCESSING uint32 = 0x4

	var modeStdout uint32
	var modeStderr uint32

	windows.GetConsoleMode(windows.Stdout, &modeStdout)
	windows.GetConsoleMode(windows.Stderr, &modeStderr)

	modeStdout |= ENABLE_VIRTUAL_TERMINAL_PROCESSING
	modeStderr |= ENABLE_VIRTUAL_TERMINAL_PROCESSING

	windows.SetConsoleMode(windows.Stdout, uint32(modeStdout))
	windows.SetConsoleMode(windows.Stderr, uint32(modeStderr))
}
