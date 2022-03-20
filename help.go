package opt

import (
	"fmt"
	"os"
	"strings"
)

var GlossaryOffset int = 2

var ColorMenu = White
var ColorOption = Cyan
var ColorMeta = BBlack
var ColorCommand = Yellow
var ColorArgs = Magenta

// Usage prints usage text to os.Stdout.
func (c *Command) Usage() {
	Println(ColorMenu+"Usage:"+End, 0)
	s := strings.Repeat(" ", GlossaryOffset)
	s += c.Name + " "

	s += "[" + ColorOption + "options" + End + "] "
	if len(c.commands) != 0 {
		s += "[" + ColorCommand + "command" + End + "]"
	} else {
		if c.Meta != "" {
			s += c.Meta + " "
		} else {
			if c.MinArgs > 0 {
				s += ColorArgs + "arguments" + End
			} else {
				s += "[" + ColorArgs + "arguments" + End + "]"
			}
		}
	}

	Println(s, GlossaryOffset)
}

// Glossary prints glossary to os.Stdout.
func (c *Command) Glossary() {
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
			s += ColorOption + "-" + c.flags[i].abbr + End
			s += strings.Repeat(" ", c.longest-2)
			s += "   "
			Println(s+c.flags[i].help, len(RemoveColor(s)))
		}
	}

	// Print Options -a=XXX
	for i := range c.options {
		s := strings.Repeat(" ", GlossaryOffset)
		if c.options[i].abbr != "" && c.options[i].name == "" {
			s += ColorOption + "-" + c.options[i].abbr + End + "=" + ColorMeta + c.options[i].meta + End
			s += strings.Repeat(" ", c.longest-2-1-len(c.options[i].meta))
			s += "   "
			Println(s+c.flags[i].help, len(RemoveColor(s)))
		}
	}

	// Print Flags -a|--aa
	for i := range c.flags {
		s := strings.Repeat(" ", GlossaryOffset)
		if c.flags[i].abbr != "" && c.flags[i].name != "" {
			s += ColorOption + "-" + c.flags[i].abbr + End + "|" + ColorOption + "--" + c.flags[i].name + End
			s += strings.Repeat(" ", c.longest-dyn-1-len(c.flags[i].abbr)-len(c.flags[i].name))
			s += "   "
			Println(s+c.flags[i].help, len(RemoveColor(s)))
		}
	}

	// Print Options -a|--aa=XXX
	for i := range c.options {
		s := strings.Repeat(" ", GlossaryOffset)
		if c.options[i].abbr != "" && c.options[i].name != "" {
			s += ColorOption + "-" + c.options[i].abbr + End + "|" + ColorOption + "--" + c.options[i].name + End + "=" + ColorMeta + c.options[i].meta + End
			s += strings.Repeat(" ", c.longest-dyn-1-len(c.options[i].abbr)-len(c.options[i].name)-1-len(c.options[i].meta))
			s += "   "
			Println(s+c.flags[i].help, len(RemoveColor(s)))
		}
	}

	// Print Options --aa=XXX
	for i := range c.options {
		s := strings.Repeat(" ", GlossaryOffset)
		if c.options[i].abbr == "" && c.options[i].name != "" {
			s += strings.Repeat(" ", dyn)
			s += ColorOption + "--" + c.options[i].name + End + "=" + ColorMeta + c.options[i].meta + End
			s += strings.Repeat(" ", c.longest-dyn-2-len(c.options[i].name)-1-len(c.options[i].meta))
			s += "   "
			Println(s+c.flags[i].help, len(RemoveColor(s)))
		}
	}

	// Flags just --long without -a
	for i := range c.flags {
		s := strings.Repeat(" ", GlossaryOffset)
		if c.flags[i].abbr == "" && c.flags[i].name != "" {
			s += strings.Repeat(" ", dyn)
			s += ColorOption + "--" + c.flags[i].name + End
			s += strings.Repeat(" ", c.longest-dyn-2-len(c.flags[i].name))
			s += "   "
			Println(s+c.flags[i].help, len(RemoveColor(s)))
		}
	}

	if !NoHelp {
		s := strings.Repeat(" ", GlossaryOffset)
		s += strings.Repeat(" ", dyn)
		s += ColorOption + "--help" + End
		s += strings.Repeat(" ", c.longest-dyn-len("--help"))
		s += "   "
		Println(s+"Show this help", len(RemoveColor(s)))
	}

	if c.Version != "" {
		s := strings.Repeat(" ", GlossaryOffset)
		s += strings.Repeat(" ", dyn)
		s += ColorOption + "--version" + End
		s += strings.Repeat(" ", c.longest-dyn-len("--version"))
		s += "   "
		Println(s+"Show version information", len(RemoveColor(s)))
	}

}

// Commands prints command help text to os.Stdout.
func (c *Command) Commands() {
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
func (c *Command) Help() {

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
		fmt.Println("Run \"" + c.Name + " [command] --help\" for more information on a command.")
	}

	os.Exit(0)
}
