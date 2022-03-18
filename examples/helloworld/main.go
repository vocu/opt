package main

import (
	"fmt"

	"github.com/vocu/opt"
)

func main() {
	// Create a new parser
	o := opt.New("hello")

	// Add a flag. The first parameter is the default value.
	flag := o.Flag(false, "flag", "f", "A simple flag")

	// You can also have a flag with no abbreviation.
	anotherFlag := o.Flag(false, "another-flag", "", "A flag that has no abbreviation")

	// An option always takes an argument. The first parameter is the default value and its type determines the type of the option.
	number := o.Option(0, "number", "n", "A number option", "<int>", false)

	// Use Run() to automatically parse arguments from os.Args
	o.Run()

	// Print values
	fmt.Println("flag:", *flag)
	fmt.Println("another-flag:", *anotherFlag)
	fmt.Println("number:", *number)

	// Print all positional arguments
	fmt.Println("arguments", o.Args)
}
