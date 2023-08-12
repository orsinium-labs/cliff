package cliff

import "github.com/spf13/pflag"

type anyFlag interface {
	PFlag() *pflag.Flag
}

type flagSet struct {
	name  string
	flags []anyFlag
}

func Flags(name string, flags ...anyFlag) flagSet {
	return flagSet{name, flags}
}

func (fs flagSet) Parse(args []string) error {
	pfs := fs.PFlagSet()
	return pfs.Parse(args)
}

func (fs flagSet) PFlagSet() pflag.FlagSet {
	pfs := pflag.NewFlagSet(fs.name, pflag.ContinueOnError)
	for _, f := range fs.flags {
		pfs.AddFlag(f.PFlag())
	}
	return *pfs
}
