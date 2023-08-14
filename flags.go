package cliff

import (
	"fmt"
	"io"

	"github.com/spf13/pflag"
)

func MustParse[T any](
	stderr io.Writer,
	exit func(int),
	args []string,
	init func(c *T) Flags,
) T {
	config, err := Parse[T](stderr, args, init)
	HandleError(stderr, exit, err)
	return config
}

func Parse[T any](stderr io.Writer, args []string, init func(c *T) Flags) (T, error) {
	var config T
	flags := init(&config)
	err := flags.Parse(stderr, args)
	return config, err
}

// Flags is a mapping of CLI flag names to the flags.
type Flags map[string]Flag

// Parse the given arguments.
//
// The first argument is the program name.
//
// Help and warnings will be written into the given stderr stream.
//
// Typical usage:
//
//	flags.Parse(os.Stderr, os.Args)
func (fs Flags) Parse(stderr io.Writer, args []string) error {
	pfs, err := fs.PFlagSet(stderr, args[0])
	if err != nil {
		return err
	}
	return pfs.Parse(args[1:])
}

// PFlagSet returns a pflag.FlagSet populated with defined flags.
func (fs Flags) PFlagSet(stderr io.Writer, name string) (*pflag.FlagSet, error) {
	pfs := pflag.NewFlagSet(name, pflag.ContinueOnError)
	pfs.SetOutput(stderr)
	for name, f := range fs {
		err := f.pflagAdd(name, pfs)
		if err != nil {
			return nil, fmt.Errorf("add flag %s: %v", name, err)
		}
	}
	return pfs, nil
}

// HandleError interrupts the program if an error occured when parsing arguments.
func HandleError(stderr io.Writer, exit func(int), err error) {
	if err == nil {
		return
	}
	if err == pflag.ErrHelp {
		exit(0)
	}
	fmt.Fprintln(stderr, err)
	exit(2)
}
