package cliff

import (
	"flag"

	"github.com/spf13/pflag"
)

// tGoFlag represents all info about a CLI flag except its name.
type tGoFlag struct {
	baseFlag
	flag  *flag.Flag
	short string // short alias for the flag
}

// GoFlag creates a new flag from an stdlib [flag.Flag].
func GoFlag(short Short, flag *flag.Flag) Flag {
	shortStr := ""
	if short != 0 {
		shortStr = string(short)
	}
	return tGoFlag{
		flag:  flag,
		short: shortStr,
	}
}

func (f tGoFlag) Deprecated(message string) Flag {
	f.depr = message
	return f
}

func (f tGoFlag) ShortDeprecated(message string) Flag {
	f.shortDepr = message
	return f
}

func (f tGoFlag) Hidden() Flag {
	f.hidden = true
	return f
}

func (f tGoFlag) AddTo(fs *pflag.FlagSet, name string) error {
	pf := pflag.PFlagFromGoFlag(f.flag)
	pf.Name = name
	pf.Shorthand = f.short
	fs.AddFlag(pf)
	return f.setProperties(fs, name)
}
