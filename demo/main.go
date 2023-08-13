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
	config, err := cliff.Parse(os.Stderr, os.Args, func(c *Config) cliff.Flags {
		return cliff.Flags{
			"host":  cliff.F(&c.host, 0, "127.0.0.1", "host to serve on"),
			"port":  cliff.F(&c.port, 'p', 8080, "port to listen to"),
			"https": cliff.F(&c.https, 0, true, "force https"),
		}
	})
	cliff.HandleError(err)
	fmt.Printf("%#v\n", config)
}
