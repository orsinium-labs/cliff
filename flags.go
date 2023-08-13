package cliff

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

type Flags map[string]tFlag

func (fs Flags) Parse(args []string) error {
	pfs := fs.PFlagSet(args[0])
	return pfs.Parse(args[1:])
}

func (fs Flags) PFlagSet(name string) pflag.FlagSet {
	pfs := pflag.NewFlagSet(name, pflag.ContinueOnError)
	for name, f := range fs {
		f.pflagAdd(name, pfs)
	}
	return *pfs
}

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
