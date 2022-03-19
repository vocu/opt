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
func (c *command) Flag(name string, abbr string, help string) *bool {
	if abbr != "" {
		c.gotAbbr = true
		c.usgShortFlags += abbr
	} else if name != "" {
		c.usgFlags += "[" + ColorOption + "--" + name + End + "] "
	}

	checkLength := 0
	if abbr != "" && name != "" {
		checkLength += 1
	}
	if name != "" {
		checkLength += 2 + len(name)
	}
	if checkLength > c.longest {
		c.longest = checkLength
	}

	c.flags = append(c.flags, &flag{false, name, abbr, help})
	return &c.flags[len(c.flags)-1].val
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
func (c *command) Option(val interface{}, name string, abbr string, help string, meta string, mandatory bool) *interface{} {
	if abbr != "" {
		c.gotAbbr = true
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
	if checkLength > c.longest {
		c.longest = checkLength
	}
	// ----

	// generate usage string
	if mandatory {
		c.usg += ColorOption + "-" + End
	} else {
		c.usg += "[" + ColorOption + "-" + End
	}
	if abbr != "" {
		c.usg += ColorOption + abbr + End + "=" + ColorMeta + meta + End
	} else if name != "" {
		c.usg += ColorOption + "-" + name + End + "=" + ColorMeta + meta + End
	}
	if mandatory {
		c.usg += " "
	} else {
		c.usg += "] "
	}
	// -------

	c.options = append(c.options, &option{val, name, abbr, help, meta, mandatory})
	return &c.options[len(c.options)-1].val
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
