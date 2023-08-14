# cliff

The simplest and safest golang library for making CLI tools.

😎 Features:

* 📔 Follows [POSIX argument syntax convention](https://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html) (flags look `--like-this`).
* 🛡 Safe, uses the full power of type safety and generics to detect errors at compilation time (see below)
* 🔨 Makes simple simple and hard possible.
* 💪 Reliable, just a thin wrapper around old, popular, and battle-tested [pflag](https://github.com/spf13/pflag/).
* 🍸 Can be mixed together with [flag](https://pkg.go.dev/flag), [pflag](https://github.com/spf13/pflag/), and [cobra](https://github.com/spf13/cobra).
* 🔋 Supports long and short names for flags, hidden flags, flag deprecation.
* 📑 Well-documented, with examples for every function.

## 🛡 Safety

The following is checked at compilation time:

* You don't use values before parsing arguments.
* You don't initialize the same flag twice.
* You use only one letter for flag shorthands.
* You provide default values of the correct types.
* You don't forget to provide help messages.
* You pass pointers, not values, to flag targets.
* You use only types supported by the library.

## 📦 Installation

```bash
go get github.com/orsinium-labs/cliff
```

## 🛠️ Usage

```go
type Config struct {
  host  string
  port  int
  debug bool
}

flags := func(c *Config) cliff.Flags {
  return cliff.Flags{
    "host":  cliff.F(&c.host, 0, "127.0.0.1", "host to serve on"),
    "port":  cliff.F(&c.port, 'p', 8080, "port to listen to"),
    "debug": cliff.F(&c.debug, 0, false, "run in debug mode"),
  }
}

config := cliff.MustParse(os.Stderr, os.Exit, os.Args, flags)
fmt.Printf("%#v\n", config)
```

Passing `os` parameters makes side-effects explicit and easy to override in tests, and defining flags inside a function makes it impossible to use the config before initialization.

Usage examples:

```bash
# show usage and exit with 0
example --help

# specify some flags
example --host localhost --debug

# use the short version of a flag
example -p 80

# try passing invalid flag, will print error and exit with 2
example -p hi
```

## 🔌 Integrating with other packages

Use cliff to specify flags for a [pflag](https://github.com/spf13/pflag/) flag set:

```go
flags := cliff.Flags{
  "host":  cliff.F(&c.host, 0, "127.0.0.1", ""),
}
subFlagSet, err := flags.PFlagSet()
if err != nil {
  return err
}

flagSet := pflag.NewFlagSet("example", pflag.ContinueOnError)
flagSet.AddFlagSet(subFlagSet)
```

Similarly, use cliff to specify flags for a [cobra](https://github.com/spf13/cobra) command:

```go
flags := cliff.Flags{
  "host":  cliff.F(&c.host, 0, "127.0.0.1", ""),
}
flagSet, err := flags.PFlagSet()
if err != nil {
  return err
}
someCmd.PersistentFlags().AddFlagSet(flagSet)
```

Use stdlib [flag](https://pkg.go.dev/flag) flags with cliff:

```go
var debug bool
flag.BoolVar(&debug, "debug", false, "run in debug mode")
cliff.Flags{
   "debug": cliff.GoFlag('d', flag.Lookup("debug")),
}
```
