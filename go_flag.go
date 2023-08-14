package cliff

import (
	"errors"
	"flag"
	"fmt"

	"github.com/spf13/pflag"
)

// tGoFlag represents all info about a CLI flag except its name.
type tGoFlag struct {
	flag  flag.Flag
	short string // short alias for the flag

	depr      string // deprecation message
	shortDepr string // deprecation message for the shorthand
	hidden    bool   // don't show the flag in help
	internal  bool   // a check that the flag is constructed using the constructor
}

// F creates a new flag.
func GoFlag[T Constraint](short Short, flag flag.Flag) Flag {
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
	if !f.internal {
		return errors.New("cliff.tPFlag must be instantiated using cliff.F constructor")
	}
	if f.short != "" && !isAlNum(f.short) {
		return errors.New("flag short name must be an alpha-numeric ASCII character")
	}
	pf := pflag.PFlagFromGoFlag(&f.flag)
	pf.Name = name
	fs.AddFlag(pf)

	var err error
	if f.depr != "" {
		err = fs.MarkDeprecated(name, f.depr)
		if err != nil {
			return fmt.Errorf("mark deprecated: %v", err)
		}
	}
	if f.shortDepr != "" {
		err = fs.MarkShorthandDeprecated(name, f.shortDepr)
		if err != nil {
			return fmt.Errorf("mark short deprecated: %v", err)
		}
	}
	if f.hidden {
		err = fs.MarkHidden(name)
		if err != nil {
			return fmt.Errorf("mark hidden: %v", err)
		}
	}
	return nil
}
