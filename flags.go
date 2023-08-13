package cliff

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/pflag"
)

type Flags map[string]tFlag

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
	pfs := fs.PFlagSet(stderr, args[0])
	return pfs.Parse(args[1:])
}

// PFlagSet returns a pflag.FlagSet populated with defined flags.
func (fs Flags) PFlagSet(stderr io.Writer, name string) *pflag.FlagSet {
	pfs := pflag.NewFlagSet(name, pflag.ContinueOnError)
	pfs.SetOutput(stderr)
	for name, f := range fs {
		f.pflagAdd(name, pfs)
	}
	return pfs
}

// HandleError interrupts the program if an error occured when parsing arguments.
func HandleError(err error) {
	if err == nil {
		return
	}
	if err == pflag.ErrHelp {
		os.Exit(0)
	}
	fmt.Println(err)
	os.Exit(2)
}
