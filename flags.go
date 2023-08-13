package cliff

import "github.com/spf13/pflag"

type tFlagSet struct {
	name  string
	flags []tFlag
}

func Flags(name string, flags ...tFlag) tFlagSet {
	return tFlagSet{name, flags}
}

func (fs tFlagSet) Parse(args []string) error {
	pfs := fs.PFlagSet()
	return pfs.Parse(args)
}

func (fs tFlagSet) PFlagSet() pflag.FlagSet {
	pfs := pflag.NewFlagSet(fs.name, pflag.ContinueOnError)
	for _, f := range fs.flags {
		pfs.AddFlag(f.PFlag())
	}
	return *pfs
}
