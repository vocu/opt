package opt

import (
	"os"
	"strings"
)

var NoHelp bool

type Command struct {
	parsedDashDash bool
	parsedArgs     uint64

	gotAbbr        bool
	longest        int
	longestCommand int
	usgShortFlags  string
	usgFlags       string
	usg            string

	commands []*Command
	flags    []*flag
	options  []*option

	Args    []string
	Func    func(args []string)
	MaxArgs uint64
	MinArgs uint64

	Name    string
	Short   string
	Long    string
	Version string
	Meta    string
}

func New(name string) *Command {
	return &Command{Name: name, MaxArgs: uint64(18446744073709551615)}
}

func (c *Command) Add(name string) *Command {
	c.commands = append(c.commands, &Command{Name: name, MaxArgs: uint64(18446744073709551615)})
	if len(name) > c.longestCommand {
		c.longestCommand = len(name)
	}
	return c.commands[len(c.commands)-1]
}

func tooManyArguments() {
	Error("too many arguments")
}

func invalidOption(s string) {
	Error("invalid option \"", s, "\"")
}

func invalidCommand(s string) {
	Error("invalid command \"", s, "\"")
}

func missingArgument() {
	Error("missing argument(s)")
}

func (c *Command) Parse(args []string) {
OUTER:
	for i := 0; i < len(args); i++ {
		if c.parsedDashDash {
			// Parse as argument
			if c.parsedArgs < c.MaxArgs {
				c.parsedArgs += 1
				c.Args = append(c.Args, args[i])
			} else {
				tooManyArguments()
			}
		} else {
			if args[i] == "--" {
				c.parsedDashDash = true
			} else {
				// special case: '--help' takes precedence over other options
				if !NoHelp && args[i] == "--help" {
					c.Help()

					// EXIT!!!
				}

				if c.Version != "" && args[i] == "--version" {
					Println(c.Version, 0)
					os.Exit(0)
				}

				// Flags
				for j := range c.flags {
					if (c.flags[j].abbr != "" && args[i] == "-"+c.flags[j].abbr) || (c.flags[j].name != "" && args[i] == "--"+c.flags[j].name) {
						c.flags[j].val = true
						continue OUTER
					}
				}

				// Options
				for j := range c.options {
					if idx := strings.Index(args[i], "="); idx != -1 {
						if (c.options[j].abbr != "" && strings.HasPrefix(args[i], "-"+c.options[j].abbr+"=")) || (c.options[j].name != "" && strings.HasPrefix(args[i], "--"+c.options[j].name+"=")) {
							c.options[j].set(args[i][idx+1:])
							continue OUTER
						}
					} else {
						if (c.options[j].abbr != "" && args[i] == "-"+c.options[j].abbr) || (c.options[j].name != "" && args[i] == "--"+c.options[j].name) {
							if len(args) >= i+2 {
								if strings.HasPrefix(args[i+1], "-") {
									missingOptionArgument(c.options[j].abbr, c.options[j].name)
								}
								c.options[j].set(args[i+1])
								i += 1
								continue OUTER
							} else {
								missingOptionArgument(c.options[j].abbr, c.options[j].name)
							}
						}
					}
				}

				// We already parsed all long options
				if strings.HasPrefix(args[i], "--") {
					invalidOption(args[i])
				}

				// -abc without =
				if idx := strings.Index(args[i], "="); idx == -1 && len(args[i]) >= 3 && strings.HasPrefix(args[i], "-") {
				INNER_1:
					for j := 1; j < len(args[i])-1; j++ {
						for k := range c.flags {
							if c.flags[k].abbr != "" && string(args[i][j]) == c.flags[k].abbr {
								c.flags[k].val = true
								continue INNER_1
							}
						}
						invalidOption("-" + string(args[i][j]))
					}
					args[i] = "-" + string(args[i][len(args[i])-1])
					i -= 1
					continue OUTER
				}

				// -abc=XXX
				if idx := strings.Index(args[i], "="); idx != -1 && len(args[i]) >= 5 && strings.HasPrefix(args[i], "-") {
				INNER_2:
					for j := 1; j < idx-1; j++ {
						for k := range c.flags {
							if c.flags[k].abbr != "" && string(args[i][j]) == c.flags[k].abbr {
								c.flags[k].val = true
								continue INNER_2
							}
						}
						invalidOption("-" + string(args[i][j]))
					}
					args[i] = "-" + string(args[i][idx-1:])
					i -= 1
					continue OUTER
				}

				// We alrady parsed all options
				if strings.HasPrefix(args[i], "-") {
					invalidOption(args[i])
				}

				// Parse commands if any
				if len(c.commands) > 0 {
					for j := range c.commands {
						if args[i] == c.commands[j].Name {
							c.commands[j].Parse(args[i+1:])
							return
						}
					}
					invalidCommand(args[i])
				}

				// Parse as argument
				if c.parsedArgs < c.MaxArgs {
					c.parsedArgs += 1
					c.Args = append(c.Args, args[i])
				} else {
					tooManyArguments()
				}

			}
		}

	} // Outer for loop

	if c.parsedArgs < c.MinArgs {
		missingArgument()
	}

	if c.Func != nil {
		c.Func(c.Args)
	}
}

func (c *Command) Run() {
	c.Parse(os.Args[1:])
}
