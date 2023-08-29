# 🏔 cliff

[ [📄 docs](https://pkg.go.dev/github.com/orsinium-labs/cliff) ] [ [🐙 github](https://github.com/orsinium-labs/cliff) ] [ [❤️ sponsor](https://github.com/sponsors/orsinium) ]

The simplest and safest golang library for making CLI tools.

😎 Features:

* 📔 Follows [POSIX argument syntax convention](https://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html) (flags look `--like-this`).
* 🛡 Safe, uses the full power of type safety and generics to detect errors at compilation time (see below)
* 🔨 Makes simple simple and hard possible.
* 💪 Reliable, just a thin wrapper around old, popular, and battle-tested [pflag].
* 🍸 Can be mixed together with [flag], [pflag], [ff], and [cobra].
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

Also in runtime:

* Never panics.
* Makes sure all flag names and short names follow POSIX recommendations.

Read the blog post to learn more: [Writing safe-to-use Go libraries](https://blog.orsinium.dev/posts/go/safe-api/).

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
    "debug": cliff.F(&c.debug, 'd', false, "run in debug mode"),
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

Use cliff to specify flags for a [pflag] flag set:

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

Similarly, use cliff to specify flags for a [cobra] command:

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

Use cliff to specify flags for [ff] commands:

```go
flags := cliff.Flags{
  "host":  cliff.F(&c.host, 0, "127.0.0.1", ""),
}
flagSet, err := flags.FlagSet()
if err != nil {
  return err
}
err := ff.Parse(flagSet, os.Args[1:])
```

Use stdlib [flag] flags with cliff:

```go
var debug bool
flag.BoolVar(&debug, "debug", false, "run in debug mode")
cliff.Flags{
   "debug": cliff.GoFlag('d', flag.Lookup("debug")),
}
```

## 🤔 QnA

1. 🤷 **Q: Why to make yet another library?** A: All the big CLI libraries in Go (like [flag] and [pflag]) were born long before generics, and so their API is full of messy functions for each possible variable type like `Float64SliceVarP`. The main goal of the project is to make the API nice, small, and clean. And along the way I had opportunity to improve quite a few things in terms of safety and best practices by stripping away global state and side-effects and using maps and closures.
1. 😡 **Q: Why it doesn't support subcommands, autocomplete for all shells, aliases, env vars, config files, and all other features I can't live without?** A: The project is designed to be simple and reliable for small projects and simple CLIs, a better version of [pflag]. If you need more, take a look at [ff], [kong](https://github.com/alecthomas/kong), [cobra], and [urfave/cli](https://github.com/urfave/cli).
1. 🤝 **Q: How can I contribute?** If you found a bug or want to improve something a bit, please, send a PR, and I'll merge it. I'm easy to agree with and I usually merge everything within a day.
1. 🕵 **Q: Why there are so many ways to do things?** A: The only function you need to use is `cliff.MustParse`, and for that you'll natuarally need `cliff.Flags` and `cliff.F`. That's it. Everything elsle is here for the situations when you need to mix cliff with another library, emit results into multiple variables, parse some tricky custom values, and so on. Exposing all these things is the cost of flexibility.
1. 🦀 **Q: Rust is better.** I think [clap](https://github.com/clap-rs/clap) is pretty neat and I like the idea that you can define a single struct with some fields and their attributes and the CLI is magically generated for it. However, while Rust has a standard syntax for such attributes and powerful compile-time macros, in Go we have to use struct field tags like in [encoding/json](https://pkg.go.dev/encoding/json) and that is easy to mess up and doesn't provide any compile-time guarantees.

[pflag]: https://github.com/spf13/pflag/
[flag]: https://pkg.go.dev/flag
[ff]: https://github.com/peterbourgon/ff
[cobra]: https://github.com/spf13/cobra
