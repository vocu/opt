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
func (c *command) Usage() {
	Println(ColorMenu+"Usage:"+End, 0)
	s := strings.Repeat(" ", GlossaryOffset)
	s += c.Name + " "

	// TEST
	//s += "[" + ColorOption + "options" + End + "] [" + ColorCommand + "command" + End + "] [" + ColorArgs + "args" + End + "]"
	// maybe go mod <command> [arguments]
	// maybe prog [options] <command>
	//Println(s, GlossaryOffset)
	//return
	// TEST END

	if c.usgShortFlags != "" {
		s += "[" + ColorFlag + "-" + c.usgShortFlags + End + "] "
	}

	s += c.usgFlags

	if !NoHelp {
		s += "[" + ColorFlag + "--" + "help" + End + "] "
	}
	if c.Version != "" {
		s += "[" + ColorFlag + "--" + "version" + End + "] "
	}

	s += c.usg

	if c.Meta != "" {
		s += c.Meta + " "
	} else {
		if len(c.commands) > 0 {
			s += "<" + ColorCommand + "command" + End + "> ... "
		} else {
			if c.MinArgs > 0 {
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
func (c *command) Glossary() {
	if c.Version != "" && c.longest == 0 {
		c.longest = len("--version")
	}
	if !NoHelp && c.longest == 0 {
		c.longest = len("--help")
	}

	Println(ColorMenu+"Options: "+End, 0)

	// dynamic length of 3 spaces to print if there exists at least one optin/flag with long name
	dyn := 0
	if c.gotAbbr {
		c.longest += 3
		dyn = 3
	}
	// End Find

	// Print Flags -a
	for i := range c.flags {
		s := strings.Repeat(" ", GlossaryOffset)

		if c.flags[i].abbr != "" && c.flags[i].name == "" {
			s += ColorFlag + "-" + c.flags[i].abbr + End
			s += strings.Repeat(" ", c.longest-2)
			s += "   " + c.flags[i].help
			Println(s, GlossaryOffset)
		}
	}

	// Print Options -a=XXX
	for i := range c.options {
		s := strings.Repeat(" ", GlossaryOffset)
		if c.options[i].abbr != "" && c.options[i].name == "" {
			s += ColorOption + "-" + c.options[i].abbr + End + "=" + ColorMeta + c.options[i].meta + End
			s += strings.Repeat(" ", c.longest-2-1-len(c.options[i].meta))
			s += "   " + c.options[i].help
			Println(s, GlossaryOffset)
		}
	}

	// Print Flags -a|--aa
	for i := range c.flags {
		s := strings.Repeat(" ", GlossaryOffset)
		if c.flags[i].abbr != "" && c.flags[i].name != "" {
			s += ColorFlag + "-" + c.flags[i].abbr + End + "|" + ColorFlag + "--" + c.flags[i].name + End
			s += strings.Repeat(" ", c.longest-dyn-1-len(c.flags[i].abbr)-len(c.flags[i].name))
			s += "   " + c.flags[i].help
			Println(s, GlossaryOffset)
		}
	}

	// Print Options -a|--aa=XXX
	for i := range c.options {
		s := strings.Repeat(" ", GlossaryOffset)
		if c.options[i].abbr != "" && c.options[i].name != "" {
			s += ColorOption + "-" + c.options[i].abbr + End + "|" + ColorOption + "--" + c.options[i].name + End + "=" + ColorMeta + c.options[i].meta + End
			s += strings.Repeat(" ", c.longest-dyn-1-len(c.options[i].abbr)-len(c.options[i].name)-1-len(c.options[i].meta))
			s += "   " + c.options[i].help
			Println(s, GlossaryOffset)
		}
	}

	// Print Options --aa=XXX
	for i := range c.options {
		s := strings.Repeat(" ", GlossaryOffset)
		if c.options[i].abbr == "" && c.options[i].name != "" {
			s += strings.Repeat(" ", dyn)
			s += ColorOption + "--" + c.options[i].name + End + "=" + ColorMeta + c.options[i].meta + End
			s += strings.Repeat(" ", c.longest-dyn-2-len(c.options[i].name)-1-len(c.options[i].meta))
			s += "   " + c.options[i].help
			Println(s, GlossaryOffset)
		}
	}

	// Flags just --long without -a
	for i := range c.flags {
		s := strings.Repeat(" ", GlossaryOffset)
		if c.flags[i].abbr == "" && c.flags[i].name != "" {
			s += strings.Repeat(" ", dyn)
			s += ColorFlag + "--" + c.flags[i].name + End
			s += strings.Repeat(" ", c.longest-dyn-2-len(c.flags[i].name))
			s += "   " + c.flags[i].help
			Println(s, GlossaryOffset)
		}
	}

	if !NoHelp {
		s := strings.Repeat(" ", GlossaryOffset)
		s += strings.Repeat(" ", dyn)
		s += ColorFlag + "--help" + End
		s += strings.Repeat(" ", c.longest-dyn-len("--help"))
		s += "   " + "Show this help"
		Println(s, GlossaryOffset)
	}

	if c.Version != "" {
		s := strings.Repeat(" ", GlossaryOffset)
		s += strings.Repeat(" ", dyn)
		s += ColorFlag + "--version" + End
		s += strings.Repeat(" ", c.longest-dyn-len("--version"))
		s += "   " + "Show version information"
		Println(s, GlossaryOffset)
	}

}

// Commands prints command help text to os.Stdout.
func (c *command) Commands() {
	if len(c.commands) > 0 {
		Println(ColorMenu+"Commands: "+End, 0)
		for i := range c.commands {
			s := strings.Repeat(" ", GlossaryOffset)
			//s += strings.Repeat(" ", dyn)
			s += ColorCommand + c.commands[i].Name + End
			s += strings.Repeat(" ", c.longestCommand-len(c.commands[i].Name))
			s += "   " + c.commands[i].Short
			Println(s, GlossaryOffset)
		}
		fmt.Println()
	}
}

// Help prints help text to os.Stdout and exits.
func (c *command) Help() {

	// Print Description
	if c.Long != "" {
		Println(c.Long, 0)
	}
	if c.Long != "" {
		fmt.Println()
	}

	c.Usage()
	fmt.Println()
	c.Commands()
	c.Glossary()

	if len(c.commands) > 0 && !NoHelp {
		fmt.Println()
		Println("Run \""+c.Name+" <command> --help\" for more information on a command.", 0)
	}

	os.Exit(0)
}
