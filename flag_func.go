package cliff

import (
	"errors"
	"flag"

	"github.com/spf13/pflag"
)

// tGoFlag represents all info about a CLI flag except its name.
type tFuncFlag[T any] struct {
	tar    *T
	parser func(string) (T, error)
	short  string // short alias for the flag
	help   string // usage message
}

// FuncFlag creates a new flag that is parsed by the given function.
func FuncFlag[T any](
	tar *T,
	short Short,
	def T,
	parser func(string) (T, error),
	help Help,
) Flag {
	shortStr := ""
	if short != 0 {
		shortStr = string(short)
	}
	setter := tFuncFlag[T]{
		tar:    tar,
		parser: parser,
		short:  shortStr,
		help:   string(help),
	}
	return Flag{setter: setter}
}

func (f tFuncFlag[T]) AddTo(fs *pflag.FlagSet, name string) error {
	if f.short != "" && !isAlNum(f.short) {
		return errors.New("flag short name must be an alpha-numeric ASCII character")
	}

	patched := func(raw string) error {
		val, err := f.parser(raw)
		if err != nil {
			return err
		}
		*f.tar = val
		return nil
	}

	gfs := flag.NewFlagSet("", flag.ContinueOnError)
	gfs.Func(name, f.help, patched)
	goflag := gfs.Lookup(name)
	pf := pflag.PFlagFromGoFlag(goflag)
	pf.Name = name
	pf.Shorthand = f.short
	fs.AddFlag(pf)
	return nil
}
