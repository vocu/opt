package opt

import (
	"os"
	"strings"
)

var NoHelp bool

type command struct {
	parsedDashDash bool
	parsedArgs     uint64

	gotAbbr        bool
	longest        int
	longestCommand int
	usgShortFlags  string
	usgFlags       string
	usg            string

	commands []*command
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

func New(name string) *command {
	return &command{Name: name, MaxArgs: uint64(18446744073709551615)}
}

func (parser *command) Command(name string) *command {
	parser.commands = append(parser.commands, &command{Name: name, MaxArgs: uint64(18446744073709551615)})
	if len(name) > parser.longestCommand {
		parser.longestCommand = len(name)
	}
	return parser.commands[len(parser.commands)-1]
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

func (parser *command) Parse(args []string) {
OUTER:
	for i := 0; i < len(args); i++ {
		if parser.parsedDashDash {
			// Parse as argument
			if parser.parsedArgs < parser.MaxArgs {
				parser.parsedArgs += 1
				parser.Args = append(parser.Args, args[i])
			} else {
				tooManyArguments()
			}
		} else {
			if args[i] == "--" {
				parser.parsedDashDash = true
			} else {
				// special case: '--help' takes precedence over other options
				if !NoHelp && args[i] == "--help" {
					parser.Help()

					// EXIT!!!
				}

				if parser.Version != "" && args[i] == "--version" {
					Println(parser.Version, 0)
					os.Exit(0)
				}

				// Flags
				for j := range parser.flags {
					if (parser.flags[j].abbr != "" && args[i] == "-"+parser.flags[j].abbr) || (parser.flags[j].name != "" && args[i] == "--"+parser.flags[j].name) {
						parser.flags[j].val = true
						continue OUTER
					}
				}

				// Options
				for j := range parser.options {
					if idx := strings.Index(args[i], "="); idx != -1 {
						if (parser.options[j].abbr != "" && strings.HasPrefix(args[i], "-"+parser.options[j].abbr+"=")) || (parser.options[j].name != "" && strings.HasPrefix(args[i], "--"+parser.options[j].name+"=")) {
							parser.options[j].set(args[i][idx+1:])
							continue OUTER
						}
					} else {
						if (parser.options[j].abbr != "" && args[i] == "-"+parser.options[j].abbr) || (parser.options[j].name != "" && args[i] == "--"+parser.options[j].name) {
							if len(args) >= i+2 {
								if strings.HasPrefix(args[i+1], "-") {
									missingOptionArgument(parser.options[j].abbr, parser.options[j].name)
								}
								parser.options[j].set(args[i+1])
								i += 1
								continue OUTER
							} else {
								missingOptionArgument(parser.options[j].abbr, parser.options[j].name)
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
						for k := range parser.flags {
							if parser.flags[k].abbr != "" && string(args[i][j]) == parser.flags[k].abbr {
								parser.flags[k].val = true
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
						for k := range parser.flags {
							if parser.flags[k].abbr != "" && string(args[i][j]) == parser.flags[k].abbr {
								parser.flags[k].val = true
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
				if len(parser.commands) > 0 {
					for j := range parser.commands {
						if args[i] == parser.commands[j].Name {
							parser.commands[j].Parse(args[i+1:])
							return
						}
					}
					invalidCommand(args[i])
				}

				// Parse as argument
				if parser.parsedArgs < parser.MaxArgs {
					parser.parsedArgs += 1
					parser.Args = append(parser.Args, args[i])
				} else {
					tooManyArguments()
				}

			}
		}

	} // Outer for loop

	if parser.parsedArgs < parser.MinArgs {
		missingArgument()
	}

	if parser.Func != nil {
		parser.Func(parser.Args)
	}
}

func (parser *command) Run() {
	parser.Parse(os.Args[1:])
}
