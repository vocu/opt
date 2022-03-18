package opt

import (
	"fmt"
	"os"
	"strings"
)

var GlossaryOffset int = 2

var ColorMenu = White
var ColorFlag = Cyan
var ColorOption = Cyan
var ColorMeta = BBlack
var ColorCommand = Yellow
var ColorArgs = Magenta

// Usage prints usage text to os.Stdout.
func (parser *command) Usage() {
	Println(ColorMenu+"Usage:"+End, 0)
	s := strings.Repeat(" ", GlossaryOffset)
	s += parser.Name + " "

	// TEST
	//s += "[" + ColorOption + "options" + End + "] [" + ColorCommand + "command" + End + "] [" + ColorArgs + "args" + End + "]"
	//Println(s, GlossaryOffset)
	//return
	// TEST END

	if parser.usgShortFlags != "" {
		s += "[" + ColorFlag + "-" + parser.usgShortFlags + End + "] "
	}

	s += parser.usgFlags

	if !NoHelp {
		s += "[" + ColorFlag + "--" + "help" + End + "] "
	}
	if parser.Version != "" {
		s += "[" + ColorFlag + "--" + "version" + End + "] "
	}

	s += parser.usg

	if parser.Meta != "" {
		s += parser.Meta + " "
	} else {
		if len(parser.commands) > 0 {
			s += "<" + ColorCommand + "command" + End + "> ... "
		} else {
			if parser.MinArgs > 0 {
				s += ColorArgs + "<args>" + End + "... "
			} else {
				s += "[" + ColorArgs + "<args>" + End + "]... "
			}
		}
	}

	//Println(s[:len(s)-1], 7)
	Println(s[:len(s)-1], GlossaryOffset)
}

// Glossary prints glossary to os.Stdout.
func (parser *command) Glossary() {
	if parser.Version != "" && parser.longest == 0 {
		parser.longest = len("--version")
	}
	if !NoHelp && parser.longest == 0 {
		parser.longest = len("--help")
	}

	Println(ColorMenu+"Options: "+End, 0)

	// dynamic length of 3 spaces to print if there exists at least one optin/flag with long name
	dyn := 0
	if parser.gotAbbr {
		parser.longest += 3
		dyn = 3
	}
	// End Find

	// Print Flags -a
	for i := range parser.flags {
		s := strings.Repeat(" ", GlossaryOffset)

		if parser.flags[i].abbr != "" && parser.flags[i].name == "" {
			s += ColorFlag + "-" + parser.flags[i].abbr + End
			s += strings.Repeat(" ", parser.longest-2)
			s += "   " + parser.flags[i].help
			Println(s, GlossaryOffset)
		}
	}

	// Print Options -a=XXX
	for i := range parser.options {
		s := strings.Repeat(" ", GlossaryOffset)
		if parser.options[i].abbr != "" && parser.options[i].name == "" {
			s += ColorOption + "-" + parser.options[i].abbr + End + "=" + ColorMeta + parser.options[i].meta + End
			s += strings.Repeat(" ", parser.longest-2-1-len(parser.options[i].meta))
			s += "   " + parser.options[i].help
			Println(s, GlossaryOffset)
		}
	}

	// Print Flags -a|--aa
	for i := range parser.flags {
		s := strings.Repeat(" ", GlossaryOffset)
		if parser.flags[i].abbr != "" && parser.flags[i].name != "" {
			s += ColorFlag + "-" + parser.flags[i].abbr + End + "|" + ColorFlag + "--" + parser.flags[i].name + End
			s += strings.Repeat(" ", parser.longest-dyn-1-len(parser.flags[i].abbr)-len(parser.flags[i].name))
			s += "   " + parser.flags[i].help
			Println(s, GlossaryOffset)
		}
	}

	// Print Options -a|--aa=XXX
	for i := range parser.options {
		s := strings.Repeat(" ", GlossaryOffset)
		if parser.options[i].abbr != "" && parser.options[i].name != "" {
			s += ColorOption + "-" + parser.options[i].abbr + End + "|" + ColorOption + "--" + parser.options[i].name + End + "=" + ColorMeta + parser.options[i].meta + End
			s += strings.Repeat(" ", parser.longest-dyn-1-len(parser.options[i].abbr)-len(parser.options[i].name)-1-len(parser.options[i].meta))
			s += "   " + parser.options[i].help
			Println(s, GlossaryOffset)
		}
	}

	// Print Options --aa=XXX
	for i := range parser.options {
		s := strings.Repeat(" ", GlossaryOffset)
		if parser.options[i].abbr == "" && parser.options[i].name != "" {
			s += strings.Repeat(" ", dyn)
			s += ColorOption + "--" + parser.options[i].name + End + "=" + ColorMeta + parser.options[i].meta + End
			s += strings.Repeat(" ", parser.longest-dyn-2-len(parser.options[i].name)-1-len(parser.options[i].meta))
			s += "   " + parser.options[i].help
			Println(s, GlossaryOffset)
		}
	}

	// Flags just --long without -a
	for i := range parser.flags {
		s := strings.Repeat(" ", GlossaryOffset)
		if parser.flags[i].abbr == "" && parser.flags[i].name != "" {
			s += strings.Repeat(" ", dyn)
			s += ColorFlag + "--" + parser.flags[i].name + End
			s += strings.Repeat(" ", parser.longest-dyn-2-len(parser.flags[i].name))
			s += "   " + parser.flags[i].help
			Println(s, GlossaryOffset)
		}
	}

	if !NoHelp {
		s := strings.Repeat(" ", GlossaryOffset)
		s += strings.Repeat(" ", dyn)
		s += ColorFlag + "--help" + End
		s += strings.Repeat(" ", parser.longest-dyn-len("--help"))
		s += "   " + "Show this help"
		Println(s, GlossaryOffset)
	}

	if parser.Version != "" {
		s := strings.Repeat(" ", GlossaryOffset)
		s += strings.Repeat(" ", dyn)
		s += ColorFlag + "--version" + End
		s += strings.Repeat(" ", parser.longest-dyn-len("--version"))
		s += "   " + "Show version information"
		Println(s, GlossaryOffset)
	}

}

// Commands prints command help text to os.Stdout.
func (parser *command) Commands() {
	if len(parser.commands) > 0 {
		Println(ColorMenu+"Commands: "+End, 0)
		for i := range parser.commands {
			s := strings.Repeat(" ", GlossaryOffset)
			//s += strings.Repeat(" ", dyn)
			s += ColorCommand + parser.commands[i].Name + End
			s += strings.Repeat(" ", parser.longestCommand-len(parser.commands[i].Name))
			s += "   " + parser.commands[i].Short
			Println(s, GlossaryOffset)
		}
		fmt.Println()
	}
}

// Help prints help text to os.Stdout and exits.
func (parser *command) Help() {

	// Print Description
	if parser.Long != "" {
		Println(parser.Long, 0)
	}
	if parser.Long != "" {
		fmt.Println()
	}

	parser.Usage()
	fmt.Println()
	parser.Commands()
	parser.Glossary()

	if len(parser.commands) > 0 && !NoHelp {
		fmt.Println()
		Println("Run \""+parser.Name+" <command> --help\" for more information on a command.", 0)
	}

	os.Exit(0)
}
