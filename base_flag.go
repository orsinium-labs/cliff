package cliff

import (
	"fmt"

	"github.com/spf13/pflag"
)

type baseFlag struct {
	depr      string // deprecation message
	shortDepr string // deprecation message for the shorthand
	hidden    bool   // don't show the flag in help
}

func (f baseFlag) setProperties(fs *pflag.FlagSet, name string) error {
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
