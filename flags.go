package cliff

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/spf13/pflag"
)

// MustParse parses CLI flags, writes errors and warnings into stderr and exits on error or help.
//
// Typical usage:
//
//	cliff.MustParse(os.Stderr, os.Exit, os.Args, flags)
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

// MustParse parses CLI flags, writes warnings into stderr and returns error on error or help.
//
// Typical usage:
//
//	cliff.Parse(os.Stderr, os.Args, flags)
func Parse[T any](stderr io.Writer, args []string, init func(c *T) Flags) (T, error) {
	var config T
	flags := init(&config)
	err := flags.Parse(stderr, args)
	return config, err
}

type Flag interface {
	// Mark the flag as deprecated.
	//
	// It won't be shown in help or usage messages
	// and when the user tries to use it,
	// the deprecation message will be shown.
	Deprecated(message string) Flag

	// ShortDeprecated marks the short alias of the flag as deprecated.
	//
	// The short flag won't be shown in help or usage messages
	// and when the user tries to use it,
	// the deprecation message will be shown.
	ShortDeprecated(message string) Flag

	// Hidden makes the flag to not be shown in help or usage messages.
	Hidden() Flag

	// AddTo adds the flag into the give pflag.FlagSet.
	AddTo(*pflag.FlagSet, string) error
}

// Flags is a mapping of CLI flag names to the flags.
type Flags map[string]Flag

// Parse the given arguments.
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

// PFlagSet returns a [pflag.FlagSet] populated with defined flags.
func (fs Flags) PFlagSet(stderr io.Writer, name string) (*pflag.FlagSet, error) {
	pfs := pflag.NewFlagSet(name, pflag.ContinueOnError)
	pfs.SetOutput(stderr)
	for name, flag := range fs {
		err := validateName(name)
		if err != nil {
			return nil, fmt.Errorf("validate flag name (%s): %v", name, err)
		}
		err = flag.AddTo(pfs, name)
		if err != nil {
			return nil, fmt.Errorf("add flag %s: %v", name, err)
		}
	}
	return pfs, nil
}

var hasUpper = regexp.MustCompile(`[A-Z]`).FindString
var isAlNum = regexp.MustCompile(`^[a-zA-Z0-9]$`).MatchString
var isValidFlag = regexp.MustCompile(`^[a-zA-Z0-9-]+$`).MatchString

func validateName(name string) error {
	if name == "" {
		return errors.New("must not be empty")
	}
	if hasUpper(name) != "" {
		return errors.New("must be lowercase")
	}
	if !isAlNum(name[:1]) {
		return errors.New("must start with alpha-numeric ASCII character")
	}
	if strings.Contains(name, "--") {
		return errors.New("must not contain --")
	}
	if !isValidFlag(name) {
		return errors.New("can contain only alpha-numeric ASCII characters and dashes")
	}
	return nil
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
