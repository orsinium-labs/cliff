package cliff

import (
	"fmt"

	"github.com/spf13/pflag"
)

type setter interface {
	AddTo(*pflag.FlagSet, string) error
}

// Flag represents all info about a CLI flag except its name.
type Flag struct {
	setter    setter
	depr      string // deprecation message
	shortDepr string // deprecation message for the shorthand
	hidden    bool   // don't show the flag in help
}

// Mark the flag as deprecated.
//
// It won't be shown in help or usage messages
// and when the user tries to use it,
// the deprecation message will be shown.
func (f Flag) Deprecated(message string) Flag {
	f.depr = message
	return f
}

// ShortDeprecated marks the short alias of the flag as deprecated.
//
// The short flag won't be shown in help or usage messages
// and when the user tries to use it,
// the deprecation message will be shown.
func (f Flag) ShortDeprecated(message string) Flag {
	f.shortDepr = message
	return f
}

// Hidden makes the flag to not be shown in help or usage messages.
func (f Flag) Hidden() Flag {
	f.hidden = true
	return f
}

// AddTo adds the flag into the given [pflag.FlagSet] under the given name.
func (f Flag) AddTo(fs *pflag.FlagSet, name string) error {
	err := f.setter.AddTo(fs, name)
	if err != nil {
		return err
	}

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
