package cliff_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/cliff"
)

func Test(t *testing.T) {
	is := is.New(t)

	type Config struct {
		host string
		port int
	}
	config := Config{}
	flags := cliff.Flags("example",
		cliff.Fn(&config.host, "host", 0, "127.0.0.1", "host to serve on"),
		cliff.Fn(&config.port, "port", 0, 8080, "port to listen to"),
	)
	err := flags.Parse([]string{"--host", "localhost"})
	is.NoErr(err)
	is.Equal(config.host, "localhost")
	is.Equal(config.port, 8080)
}
