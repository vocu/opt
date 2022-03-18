package opt

import (
	"strconv"
)

func wrongFormat(abbr string, name string) {
	Error("wrong format for option " + str(abbr, name))
}

func missingOptionArgument(abbr string, name string) {
	Error("missing argument for option " + str(abbr, name))
}

func str(abbr string, name string) string {
	res := "\""
	if abbr != "" {
		res += "-" + abbr
	}
	if abbr != "" && name != "" {
		res += "|"
	}
	if name != "" {
		res += "--" + name
	}
	res += "\""
	return res
}

type flag struct {
	val  bool
	name string
	abbr string
	help string
}

// Flag adds a flag.
func (parser *command) Flag(val bool, name string, abbr string, help string) *bool {
	if abbr != "" {
		parser.gotAbbr = true
		parser.usgShortFlags += abbr
	} else if name != "" {
		parser.usgFlags += "[" + ColorFlag + "--" + name + End + "] "
	}

	checkLength := 0
	if abbr != "" && name != "" {
		checkLength += 1
	}
	if name != "" {
		checkLength += 2 + len(name)
	}
	if checkLength > parser.longest {
		parser.longest = checkLength
	}

	parser.flags = append(parser.flags, &flag{val, name, abbr, help})
	return &parser.flags[len(parser.flags)-1].val
}

type option struct {
	val       interface{}
	name      string
	abbr      string
	help      string
	meta      string
	mandatory bool
}

// Option adds an option.
func (parser *command) Option(val interface{}, name string, abbr string, help string, meta string, mandatory bool) *interface{} {
	if abbr != "" {
		parser.gotAbbr = true
	}

	// Information for Glossary
	checkLength := 0
	if abbr != "" && name != "" {
		checkLength += 1
	}
	if name != "" {
		checkLength += 2 + len(name)
	}
	checkLength += 1 + len(meta)
	if checkLength > parser.longest {
		parser.longest = checkLength
	}
	// ----

	// generate usage string
	if mandatory {
		parser.usg += ColorOption + "-" + End
	} else {
		parser.usg += "[" + ColorOption + "-" + End
	}
	if abbr != "" {
		parser.usg += ColorOption + abbr + End + "=" + ColorMeta + meta + End
	} else if name != "" {
		parser.usg += ColorOption + "-" + name + End + "=" + ColorMeta + meta + End
	}
	if mandatory {
		parser.usg += " "
	} else {
		parser.usg += "] "
	}
	// -------

	parser.options = append(parser.options, &option{val, name, abbr, help, meta, mandatory})
	return &parser.options[len(parser.options)-1].val
}

func (o *option) set(s string) {
	switch o.val.(type) {
	case int:
		con, err := strconv.Atoi(s)
		if err != nil {
			wrongFormat(o.abbr, o.name)
		} else {
			o.val = int(con)
		}
	case string:
		o.val = s
	default:
		panic("the type of option " + str(o.abbr, o.name) + " cannot be parsed. This happens when you pass an unsupported type to the Add() method.")
	}
}
