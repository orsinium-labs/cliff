package cliff_test

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/orsinium-labs/cliff"
	"github.com/spf13/pflag"
)

func ExampleF() {
	type Config struct{ host string }
	flags := func(c *Config) cliff.Flags {
		return cliff.Flags{
			"host": cliff.F(&c.host, 0, "127.0.0.1", "host to serve on"),
		}
	}
	args := []string{"example", "--host", "localhost"}
	config := cliff.MustParse(os.Stderr, os.Exit, args, flags)
	fmt.Println(config.host)
	// Output: localhost
}

func ExampleFuncFlag() {
	// The example show how to use FuncFlag to parse JSON input
	type Addr struct {
		Host string `json:"host"`
	}
	type Config struct{ addr Addr }
	flags := func(c *Config) cliff.Flags {
		addrParser := func(raw string) (Addr, error) {
			var addr Addr
			err := json.Unmarshal([]byte(raw), &addr)
			return addr, err
		}
		return cliff.Flags{
			"addr": cliff.FuncFlag(&c.addr, 0, Addr{}, addrParser, "address info as JSON"),
		}
	}
	args := []string{"example", "--addr", `{"host": "localhost"}`}
	config := cliff.MustParse(os.Stderr, os.Exit, args, flags)
	fmt.Println(config.addr.Host)
	// Output: localhost
}

func ExampleMustParse() {
	type Config struct{ host string }
	flags := func(c *Config) cliff.Flags {
		return cliff.Flags{
			"host": cliff.F(&c.host, 0, "127.0.0.1", "host to serve on"),
		}
	}
	args := []string{"example", "--host", "localhost"}
	config := cliff.MustParse(os.Stderr, os.Exit, args, flags)
	fmt.Println(config.host)
	// Output: localhost
}

func ExampleParse() {
	type Config struct{ host string }
	flags := func(c *Config) cliff.Flags {
		return cliff.Flags{
			"host": cliff.F(&c.host, 0, "127.0.0.1", "host to serve on"),
		}
	}
	args := []string{"example", "--host", "localhost"}
	config, err := cliff.Parse(os.Stderr, args, flags)
	cliff.HandleError(os.Stderr, os.Exit, err)
	fmt.Println(config.host)
	// Output: localhost
}

func ExampleHandleError() {
	type Config struct{ host string }
	flags := func(c *Config) cliff.Flags {
		return cliff.Flags{
			"host": cliff.F(&c.host, 0, "127.0.0.1", "host to serve on"),
		}
	}
	args := []string{"example", "--host", "localhost"}
	config, err := cliff.Parse(os.Stderr, args, flags)
	cliff.HandleError(os.Stderr, os.Exit, err)
	fmt.Println(config.host)
	// Output: localhost
}

func ExampleFlag_Deprecated() {
	type Config struct {
		host string
		port int
		addr string
	}
	flags := func(c *Config) cliff.Flags {
		return cliff.Flags{
			"host": cliff.F(&c.host, 0, "127.0.0.1", "host to serve on"),
			"port": cliff.F(&c.port, 0, 8080, "port to serve on"),
			"addr": cliff.F(
				&c.addr, 0, "127.0.0.1:8080", "",
			).Deprecated("use --host and --port instead"),
		}
	}
	args := []string{"example", "--addr", "localhost:80"}
	config := cliff.MustParse(os.Stdout, os.Exit, args, flags)
	fmt.Println(config.addr)

	// Output:
	// Flag --addr has been deprecated, use --host and --port instead
	// localhost:80
}

func ExampleGoFlag() {
	type Config struct{ host string }
	var debug bool
	flags := func(c *Config) cliff.Flags {
		flag.BoolVar(&debug, "debug", false, "run in debug mode")
		return cliff.Flags{
			"host":  cliff.F(&c.host, 0, "127.0.0.1", "host to serve on"),
			"debug": cliff.GoFlag('d', flag.Lookup("debug")),
		}
	}
	args := []string{"example", "--debug"}
	config := cliff.MustParse(os.Stderr, os.Exit, args, flags)
	fmt.Println(config.host)
	fmt.Println(debug)
	// Output:
	// 127.0.0.1
	// true
}

func ExampleFlags() {
	type Config struct{ host string }
	flags := func(c *Config) cliff.Flags {
		return cliff.Flags{
			"host": cliff.F(&c.host, 0, "127.0.0.1", "host to serve on"),
		}
	}
	args := []string{"example", "--host", "localhost"}
	config := cliff.MustParse(os.Stderr, os.Exit, args, flags)
	fmt.Println(config.host)
	// Output: localhost
}

func ExampleFlags_PFlagSet() {
	type Config struct{ host string }
	var config Config
	flags := cliff.Flags{
		"host": cliff.F(&config.host, 0, "127.0.0.1", "host to serve on"),
	}
	subFlagSet, err := flags.PFlagSet(os.Stderr, "example")
	cliff.HandleError(os.Stderr, os.Exit, err)

	flagSet := pflag.NewFlagSet("example", pflag.ContinueOnError)
	flagSet.AddFlagSet(subFlagSet)
	args := []string{"--host", "localhost"}
	err = flagSet.Parse(args)
	cliff.HandleError(os.Stderr, os.Exit, err)
	fmt.Println(config.host)
	// Output: localhost
}

func ExampleFlags_Parse() {
	type Config struct{ host string }
	var config Config
	flags := cliff.Flags{
		"host": cliff.F(&config.host, 0, "127.0.0.1", "host to serve on"),
	}
	args := []string{"example", "--host=localhost"}
	err := flags.Parse(os.Stderr, args)
	cliff.HandleError(os.Stderr, os.Exit, err)
	fmt.Println(config.host)
	// Output: localhost
}

func ExampleCount() {
	type Config struct {
		verbosity cliff.Count
	}
	flags := func(c *Config) cliff.Flags {
		return cliff.Flags{
			"v": cliff.F(&c.verbosity, 'v', 0, "host to serve on"),
		}
	}
	args := []string{"example", "-vvv"}
	config := cliff.MustParse(os.Stderr, os.Exit, args, flags)
	fmt.Println(config.verbosity)
	// Output: 3
}

func ExampleBytesHex() {
	type Config struct {
		data cliff.BytesHex
	}
	flags := func(c *Config) cliff.Flags {
		return cliff.Flags{
			"data": cliff.F(&c.data, 'd', nil, "some binary data"),
		}
	}
	// 0x4F is 79 in decimal
	args := []string{"example", "-d", "4F"}
	config := cliff.MustParse(os.Stderr, os.Exit, args, flags)
	fmt.Println(config.data)
	// Output: [79]
}

func ExampleBytesBase64() {
	type Config struct {
		data cliff.BytesBase64
	}
	flags := func(c *Config) cliff.Flags {
		return cliff.Flags{
			"data": cliff.F(&c.data, 'd', nil, "some binary data"),
		}
	}
	args := []string{"example", "-d", "YQ=="}
	config := cliff.MustParse(os.Stderr, os.Exit, args, flags)
	fmt.Println(config.data)
	// Output: [97]
}
