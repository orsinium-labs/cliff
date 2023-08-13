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
	config := Config{}
	flags := cliff.Flags{
		"host":  cliff.F(&config.host, 0, "127.0.0.1", "host to serve on"),
		"port":  cliff.F(&config.port, 'p', 8080, "port to listen to"),
		"https": cliff.F(&config.https, 0, true, "force https"),
	}
	err := flags.Parse(os.Stderr, os.Args)
	cliff.HandleError(err)
	fmt.Printf("%#v\n", config)
}
