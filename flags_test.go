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
		cliff.F[string]{
			Name:    "host",
			T:       &config.host,
			Default: "127.0.0.1",
		},
		cliff.F[int]{
			Name:    "port",
			T:       &config.port,
			Default: 8080,
		},
	)
	err := flags.Parse([]string{"--host", "localhost"})
	is.NoErr(err)
	is.Equal(config.host, "localhost")
	is.Equal(config.port, 8080)
}
