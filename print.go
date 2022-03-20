package opt

import (
	"fmt"
	"os"
	"strings"
)

const (
	Black   = "\x1b[30m"
	Red     = "\x1b[31m"
	Green   = "\x1b[32m"
	Yellow  = "\x1b[33m"
	Blue    = "\x1b[34m"
	Magenta = "\x1b[35m"
	Cyan    = "\x1b[36m"
	White   = "\x1b[37m"

	BBlack   = "\x1b[90m"
	BRed     = "\x1b[91m"
	BGreen   = "\x1b[92m"
	BYellow  = "\x1b[93m"
	BBlue    = "\x1b[94m"
	BMagenta = "\x1b[95m"
	BCyan    = "\x1b[96m"
	BWhite   = "\x1b[97m"

	/*BBlack   = "\x1b[1;30m"
	BRed     = "\x1b[1;31m"
	BGreen   = "\x1b[1;32m"
	BYellow  = "\x1b[1;33m"
	BBlue    = "\x1b[1;34m"
	BMagenta = "\x1b[1;35m"
	BCyan    = "\x1b[1;36m"
	BWhite   = "\x1b[1;37m"*/

	End = "\x1b[0m"
)

var NoColor bool
var MaxWidth uint64 = 50

// Error prints an error message and causes the current program to exit with a status code indicating error.
func Error(a ...interface{}) {
	EColor("31", "Error: ")
	EColor("31", a...)
	EColor("31", "\n") // ??????????
	os.Exit(1)
}

// Color prints to os.Stdout. The first string is a ASNI SGR (Select Graphic Rendition) parameter to  control the color. e.g. 31 for red foreground.
func Color(c string, a ...interface{}) error {
	if !NoColor {
		if _, err := fmt.Fprint(os.Stdout, "\x1b["+c+"m"); err != nil {
			return err
		}
		if _, err := fmt.Fprint(os.Stdout, a...); err != nil {
			return err
		}

		if _, err := fmt.Fprint(os.Stdout, "\x1b[0m"); err != nil {
			return err
		}
	} else {
		_, err := fmt.Fprint(os.Stdout, a...)
		return err
	}
	return nil
}

// Color prints to os.Stderr. The first string is a ASNI SGR (Select Graphic Rendition) parameter to  control the color. e.g. 31 for red foreground.
func EColor(c string, a ...interface{}) error {
	if !NoColor {
		if _, err := fmt.Fprint(os.Stderr, "\x1b["+c+"m"); err != nil {
			return err
		}
		if _, err := fmt.Fprint(os.Stderr, a...); err != nil {
			return err
		}

		if _, err := fmt.Fprint(os.Stderr, "\x1b[0m"); err != nil {
			return err
		}
	} else {
		_, err := fmt.Fprint(os.Stderr, a...)
		return err
	}
	return nil
}

func RemoveColor(s string) string {
	s = strings.ReplaceAll(s, Black, "")
	s = strings.ReplaceAll(s, Red, "")
	s = strings.ReplaceAll(s, Green, "")
	s = strings.ReplaceAll(s, Yellow, "")
	s = strings.ReplaceAll(s, Blue, "")
	s = strings.ReplaceAll(s, Magenta, "")
	s = strings.ReplaceAll(s, Cyan, "")
	s = strings.ReplaceAll(s, White, "")

	s = strings.ReplaceAll(s, BBlack, "")
	s = strings.ReplaceAll(s, BRed, "")
	s = strings.ReplaceAll(s, BGreen, "")
	s = strings.ReplaceAll(s, BYellow, "")
	s = strings.ReplaceAll(s, BBlue, "")
	s = strings.ReplaceAll(s, BMagenta, "")
	s = strings.ReplaceAll(s, BCyan, "")
	s = strings.ReplaceAll(s, BWhite, "")

	s = strings.ReplaceAll(s, End, "")
	return s
}

func Println(s string, offset int) {
	if NoColor {
		s = RemoveColor(s)
	}

	lines := strings.Split(s, "\n")
	for i, line := range lines {
		currentOffset := 0
		if i != 0 {
			fmt.Print(strings.Repeat(" ", offset))
			currentOffset = offset
		}

		words := strings.Split(line, " ")
		for j, word := range words {
			fmt.Print(word)

			// Replace ansi SGR characters so we dont count them
			word = RemoveColor(word)

			currentOffset += len(word)

			if currentOffset > int(MaxWidth) {
				// So we dont print one \n too much
				//if i != len(lines)-1 || j != len(words)-1 {
				if j != len(words)-1 {
					fmt.Print("\n")
					fmt.Print(strings.Repeat(" ", offset))
				}

				currentOffset = offset
				continue
			}

			if j != len(words)-1 {
				fmt.Print(" ")
				currentOffset += 1
			}
		}
		fmt.Print("\n")

	}
}
