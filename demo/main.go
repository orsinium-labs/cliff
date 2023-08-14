package main

import (
	"fmt"
	"os"

	"github.com/orsinium-labs/cliff"
)

type Config struct {
	host  string
	port  int
	https bool
}

func main() {
	flags := func(c *Config) cliff.Flags {
		return cliff.Flags{
			"host":  cliff.F(&c.host, 0, "127.0.0.1", "host to serve on"),
			"port":  cliff.F(&c.port, 'p', 8080, "port to listen to"),
			"https": cliff.F(&c.https, 0, true, "force https"),
		}
	}
	config := cliff.MustParse(os.Stderr, os.Exit, os.Args, flags)
	fmt.Printf("%#v\n", config)
}
