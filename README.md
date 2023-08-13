# cliff

The simplest and safest golang library for making CLI tools.

## Simplicity

...

## Safety

The following is checked at compilation time:

* You don't use values before parsing arguments.
* You don't initialize the same flag twice.
* You use only one letter for flag shorthands.
* You provide default values of the correct types.
* You don't forget to provide help messages.
* You pass pointers, not values, to flag targets.
* You use only types supported by the library.

## Installation

```bash
go get github.com/orsinium-labs/cliff
```

## Usage

```go
type Config struct {
  host  string
  port  int
  https bool
}

flags := func(c *Config) cliff.Flags {
  return cliff.Flags{
    "host":  cliff.F(&c.host, 0, "127.0.0.1", "host to serve on"),
    "port":  cliff.F(&c.port, 'p', 8080, "port to listen to"),
    "https": cliff.F(&c.https, 0, true, "force https"),
  }
}

config := cliff.MustParse(os.Stderr, os.Args, flags)
fmt.Printf("%#v\n", config)
```
