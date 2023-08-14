package cliff

import (
	"errors"
	"flag"

	"github.com/spf13/pflag"
)

// tGoFlag represents all info about a CLI flag except its name.
type tGoFlag struct {
	baseFlag
	flag  flag.Flag
	short string // short alias for the flag
}

// GoFlag creates a new flag from an stdlib [flag.Flag].
func GoFlag[T Constraint](short Short, flag flag.Flag) Flag {
	shortStr := ""
	if short != 0 {
		shortStr = string(short)
	}
	return tGoFlag{
		flag:     flag,
		short:    shortStr,
		baseFlag: baseFlag{internal: true},
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
	if !f.internal {
		return errors.New("cliff.tPFlag must be instantiated using cliff.F constructor")
	}
	if f.short != "" && !isAlNum(f.short) {
		return errors.New("flag short name must be an alpha-numeric ASCII character")
	}
	pf := pflag.PFlagFromGoFlag(&f.flag)
	pf.Name = name
	fs.AddFlag(pf)
	return f.setProperties(fs, name)
}
